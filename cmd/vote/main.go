package main

import (
	"context"
	"log"
	"time"

	"github.com/kelseyhightower/envconfig"
	"github.com/rfdez/voting-vote/internal/creating"
	"github.com/rfdez/voting-vote/internal/platform/bus/inmemory"
	"github.com/rfdez/voting-vote/internal/platform/server/http"
	"github.com/rfdez/voting-vote/internal/platform/storage/mongodb"
)

func main() {
	var cfg config
	err := envconfig.Process("vote", &cfg)
	if err != nil {
		log.Fatal(err)
	}

	mongodbClient, err := mongodb.NewConnection(context.Background(), cfg.DbUser, cfg.DbPass, cfg.DbHost, cfg.DbPort, cfg.DbParams)
	if err != nil {
		log.Fatal(err)
	}
	defer mongodbClient.Disconnect(context.Background())

	db := mongodbClient.Database(cfg.DbName)

	var (
		commandBus = inmemory.NewCommandBus()
	)

	var (
		voteRepository = mongodb.NewVoteRepository(db, cfg.DbTimeout)
	)

	var (
		creatingService = creating.NewService(voteRepository)
	)

	var (
		createVoteCommandHandler = creating.NewVoteCommandHandler(creatingService)
	)

	commandBus.Register(creating.VoteCommandType, createVoteCommandHandler)

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

	// Database configuration
	DbUser    string        `default:"vote"`
	DbPass    string        `default:"vote"`
	DbHost    string        `default:"localhost"`
	DbPort    uint          `default:"27017"`
	DbName    string        `default:"vote"`
	DbParams  string        `default:""`
	DbTimeout time.Duration `default:"5s"`
}
