package main

import (
	"context"
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

	conf, err := parseConfig()
	if err != nil {
		panic(err)
	}
	InitLogger(conf.LogLevel)
	api := &exporter.API{Hostname: conf.Host, Username: conf.Username, Password: conf.Password, Client: &http.Client{}}
	if err := api.Login(); err != nil {
		Logger.Error("Citrix login error", zap.Error(err))
	}
	//Ragister metrics
	prometheus.MustRegister(&exporter.Exporter{API: *api, Logger: Logger})
	srv := &http.Server{
		Addr: conf.AppPort,
	}
	http.Handle("/metrics", promhttp.Handler())

	go func() {
		Logger.Info(fmt.Sprintf("Listen started on port %s", conf.AppPort))
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
