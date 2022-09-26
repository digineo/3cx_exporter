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
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"
)

func main() {
	config := flag.String("config", "config.json", "Path to config file")
	listen := flag.String("listen", ":9523", "Listening on")
	logLevel := flag.String("log_level", "INFO", "Log level")
	flag.Parse()

	api, err := parseConfig(*config)
	if err != nil {
		panic(err)
	}
	InitLogger(*logLevel)
	if err := api.Login(); err != nil {
		Logger.Error("Citrix login error", zap.Error(err))
	}
	//Ragister metrics
	prometheus.MustRegister(&exporter.Exporter{API: *api, Logger: Logger})
	srv := &http.Server{
		Addr: *listen,
	}
	http.Handle("/metrics", promhttp.Handler())

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
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*60)
	defer cancel()
	srv.Shutdown(ctx)
	Logger.Info("Http server stopped")
	os.Exit(0)
}
