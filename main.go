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
	configFile := flag.String("config", "config.json", "Path to config file")
	listen := flag.String("listen", ":9523", "Listening on")
	flag.Parse()

	config, err := parseConfig(*configFile)
	checkErr(err)
	api := exporter.NewAPI(config.Hostname, config.Username, config.Password)
	checkErr(api.Login())

	prometheus.MustRegister(&exporter.Exporter{API: api})
	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(*listen, nil))
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
