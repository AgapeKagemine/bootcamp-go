package server

import (
	"consume-api/internal/provider/routes"
	"consume-api/internal/provider/server/domain"
	"context"
	"fmt"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/rs/zerolog/log"
)

func Run() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	config := domain.NewServerConfig()

	server := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", config.Address, config.Port),
		Handler: routes.NewRoutes().Server,
	}

	go func() {
		log.Info().Msg("Starting server on port 8080...")

		err := server.ListenAndServe()

		if err != nil && err != http.ErrServerClosed {
			log.Error().Err(err).Msg("Error starting server")
		}
	}()

	<-ctx.Done()
	stop()
	log.Info().Msg("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := server.Shutdown(ctx)
	if err != nil {
		log.Error().Err(err).Msg("Error shutting down server")
	}

	log.Error().Msg("HTTP server stopped")
}
