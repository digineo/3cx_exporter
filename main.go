package main

import (
	"context"
	"flag"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/digineo/3cx_exporter/exporter"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	config := flag.String("config", "config.json", "Path to config file")
	listen := flag.String("listen", ":9523", "Listening on")
	flag.Parse()

	api, err := parseConfig(*config)
	if err != nil {
		panic(err)
	}
	if err := api.Login(); err != nil {
		panic(err)
	}
	//Ragister metrics
	prometheus.MustRegister(&exporter.Exporter{API: *api})
	srv := &http.Server{
		Addr: *listen,
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
