package main

import (
	"context"
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	port := flag.Int("port", 8080, "port to listen on")
	flag.Parse()

	// context
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("I'm running!\n"))
	})
	mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK\n"))
	})
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%v", *port),
		Handler: mux,
	}
	go func() {
		logger.Info("Starting HTTP server", "port", *port)
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			// Error starting or closing listener:
			logger.Error("HTTP Server ListenAndServe", "error", err)
			os.Exit(1)
		}
		logger.Info("HTTP Server stopped serving new connections")
	}()

	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-sigint

	ctxTimeout, ctxCancel := context.WithTimeout(ctx, 10*time.Second)
	defer ctxCancel()

	if err := srv.Shutdown(ctxTimeout); err != nil {
		// Error from closing listeners, or context timeout:
		logger.Error("HTTP server Shutdown", "error", err)
	}

	logger.Info("graceful shutdown complete.")
}
