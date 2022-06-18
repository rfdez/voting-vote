package main

import (
	"context"
	"log"
	"time"

	"github.com/kelseyhightower/envconfig"
	"github.com/rfdez/voting-vote/internal/platform/bus/inmemory"
	"github.com/rfdez/voting-vote/internal/platform/server/http"
)

func main() {
	var cfg config
	err := envconfig.Process("vote", &cfg)
	if err != nil {
		log.Fatal(err)
	}

	var (
		commandBus = inmemory.NewCommandBus()
	)

	ctx, httpSrv := http.NewServer(context.Background(), cfg.HttpHost, cfg.HttpPort, cfg.ShutdownTimeout, commandBus)
	if err := httpSrv.Run(ctx); err != nil {
		log.Fatal(err)
	}
}

type config struct {
	// HTTP Server Configuration
	HttpHost        string        `default:""`
	HttpPort        uint          `default:"8080"`
	ShutdownTimeout time.Duration `default:"10s"`
}
