package main

import (
	"context"

	"github.com/rs/zerolog/log"
	"github.com/siddg97/project-rook/pkg"
)

func main() {

	cfg, err := InitConfig()
	if err != nil {
		log.Fatal().Msgf("Failed to initialize server config due to fata error %v", err)
		panic(err)
	}

	// Propagate cancellation signals
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Start goroutines for the server
	done := make(chan error)
	go func() {
		done <- pkg.NewApiServer(ctx, cfg.LogLevel, cfg.Port).Start()
	}()

	// Shutdown server if worker(s) die
	if err = <-done; err != nil {
		log.Fatal().Msgf("Worker encountered fatal error %v", err)
	}
}
