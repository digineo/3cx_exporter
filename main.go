package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/digineo/3cx_exporter/db"
	"github.com/digineo/3cx_exporter/handlers"
	"github.com/digineo/3cx_exporter/services"
	"github.com/prometheus/common/log"
	"go.uber.org/zap"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*60)

	//Parse app configuration flags
	listen := flag.String("listen", ":9523", "Listening on")
	logLevel := flag.String("log_level", "INFO", "Log level")
	dbConnectionString := flag.String("db_connection", "", "Db connection string")
	//exportPeriod := flag.Int("export_period", 1, "Export period in minutes")

	flag.Parse()
	InitLogger(*logLevel)
	store := db.New(*dbConnectionString, Logger)

	//Create and start http server
	router := handlers.NewRouter(store)
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

	//Create monitoring job
	monitor := services.NewMonitor(store, services.NewStatusGetter, Logger)
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				instances, err := store.GetInstances()
				if err != nil {
					log.Error("Get list of instances error", zap.Error(err))
				} else {
					monitor.ProcessInstances(instances)
				}
				time.Sleep(10 * time.Second)
			}
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
