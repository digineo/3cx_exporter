package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/digineo/3cx_exporter/exporter"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	config := flag.String("config", "config.json", "Path to config file")
	listen := flag.String("listen", ":8080", "Listening on")
	flag.Parse()

	api, err := parseConfig(*config)
	checkErr(err)
	checkErr(api.Login())

	prometheus.MustRegister(&exporter.Exporter{API: *api})
	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(*listen, nil))
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
