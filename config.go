package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

const (
	CTX_HOST  string = "CTX_HOST"
	USERNAME  string = "USERNAME"
	PASSWORD  string = "PASSWORD"
	PORT      string = "PORT"
	LOG_LEVEL string = "LOG_LEVEL"
)

type AppConfig struct {
	Host     string
	Username string
	Password string
	AppPort  string
	LogLevel string
}

func parseConfig() (conf AppConfig, err error) {
	if env := os.Getenv("ENV"); env == "DEV" {
		err := godotenv.Load()
		if err != nil {
			fmt.Println("No .env files found. Using real environment")
		}

	}
	conf.Host = os.Getenv(CTX_HOST)
	conf.Username = os.Getenv(USERNAME)
	conf.Password = os.Getenv(PASSWORD)

	if conf.Host == "" || conf.Username == "" || conf.Password == "" {
		return conf, fmt.Errorf("%s, %s or %s is nil", CTX_HOST, USERNAME, PASSWORD)
	}

	conf.AppPort = os.Getenv(PORT)
	if conf.AppPort == "" {
		conf.AppPort = ":8080"
	}
	conf.LogLevel = os.Getenv(LOG_LEVEL)
	if conf.LogLevel == "" {
		conf.LogLevel = "INFO"
	}
	return conf, nil
}
