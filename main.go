package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/digineo/3cx_exporter/exporter"
	"github.com/digineo/3cx_exporter/handlers"
	"github.com/digineo/3cx_exporter/models"
	"github.com/digineo/3cx_exporter/services"
	"github.com/prometheus/client_golang/prometheus"
	"go.uber.org/zap"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*60)

	//Parse app configuration flags
	config := flag.String("config", "config.json", "Path to config file")
	listen := flag.String("listen", ":9523", "Listening on")
	logLevel := flag.String("log_level", "INFO", "Log level")

	flag.Parse()
	InitLogger(*logLevel)

	//Inital citrix config and api
	citrixConf := models.Config{ConfigPath: *config}
	_, err := citrixConf.Get()

	api := exporter.API{}

	if err != nil {
		panic(err)
	}

	if err := api.SetCreds(citrixConf.Host, citrixConf.Login, citrixConf.Password); err != nil {
		Logger.Error("Citrix login error", zap.Error(err))
	}

	// Start citrix connection checker process
	statusService := services.AppState{Status: &api}

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				statusService.CheckConnection()
				time.Sleep(5 * time.Second)
			}

		}
	}()

	//Ragister prometeus metrix
	prometheus.MustRegister(&exporter.Exporter{API: api, Logger: Logger})

	//Create and start http server
	router := handlers.NewRouter(&statusService, &api, *config, Logger)
	srv := &http.Server{
		Addr:    *listen,
		Handler: router,
	}

	go func() {
		Logger.Info(fmt.Sprintf("Listen started on port %s", *listen))
		if err := srv.ListenAndServe(); err != nil {
			Logger.Panic("Handle server error", zap.Error(err))
		}
	}()

	// Listen for os sygnals
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	Logger.Info("App Interrputtes. Waiting for graseful shutdown")
	defer cancel()
	srv.Shutdown(ctx)
	Logger.Info("Http server stopped")
	os.Exit(0)
}
