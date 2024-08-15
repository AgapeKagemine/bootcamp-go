package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"package/internal/config"
	"package/internal/provider/messaging/kafka"

	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func Run() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	defer stop()

	serverConfig := config.NewServerConfig()

	server := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", serverConfig.Address, serverConfig.Port),
		Handler: gin.New(),
	}

	// Start the Kafka consumer in the background
	go func() {
		// Add a short delay to allow the server to start
		time.Sleep(2 * time.Second)
		kafka.NewOrchestrator()
	}()

	go func() {
		log.Info().Msg(fmt.Sprintf("Starting server on port %d...", serverConfig.Port))

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
		log.Fatal().Err(err).Msg("Error shutting down server")
	}

	log.Info().Msg("HTTP server stopped")

}
