package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/digineo/3cx_exporter/handlers"
	"go.uber.org/zap"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*60)

	//Parse app configuration flags
	listen := flag.String("listen", ":9523", "Listening on")
	logLevel := flag.String("log_level", "INFO", "Log level")

	flag.Parse()
	InitLogger(*logLevel)

	//Create and start http server
	router := handlers.NewRouter()
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
