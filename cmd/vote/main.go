package main

import (
	"log"
	"time"

	"github.com/kelseyhightower/envconfig"
)

func main() {
	var cfg config
	err := envconfig.Process("vote", &cfg)
	if err != nil {
		log.Fatal(err)
	}
}

type config struct {
	// HTTP Server Configuration
	HttpHost        string        `default:""`
	HttpPort        uint          `default:"8080"`
	ShutdownTimeout time.Duration `default:"10s"`
}
