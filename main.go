package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/digineo/3cx_exporter/exporter"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {

	conf, err := parseConfig()
	if err != nil {
		panic(err)
	}
	api := &exporter.API{Hostname: conf.Host, Username: conf.Username, Password: conf.Password}
	if err := api.Login(); err != nil {
		panic(err)
	}
	//Ragister metrics
	prometheus.MustRegister(&exporter.Exporter{API: *api})
	srv := &http.Server{
		Addr: conf.AppPort,
	}
	http.Handle("/metrics", promhttp.Handler())

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			panic(err)
		}
	}()

	// Listen for os sygnals
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*60)
	defer cancel()
	srv.Shutdown(ctx)
	os.Exit(0)
}
