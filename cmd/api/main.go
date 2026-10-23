package main

import (
	"bsmart/config"
	"bsmart/logs"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		log.Printf("Unable to load port Number %v", port)
	}

	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{},
		ExposedHeaders:   []string{},
		AllowCredentials: false,
		MaxAge:           300,
	}))
	logger := logs.New()

	conn, err := config.DatabaseConnection()
	if err != nil {
		log.Fatalf("Failed to Load Database connection %v", err)
	}
	cfg := config.Apiconfig{
		DB:     conn,
		Logger: logger,
	}
	defer cfg.DB.Close()

	v1router := chi.NewRouter()
	router.Mount("/v1", v1router)

	server := &http.Server{
		Addr:         ":3030",
		Handler:      router,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Fatal(err)
		}
	}()

	signChannel := make(chan os.Signal, 1)
	signal.Notify(signChannel, os.Interrupt, syscall.SIGTERM)
	defer signal.Stop(signChannel)
	sig := <-signChannel

	// ("Received terminated, gracefully shutting down %v", sig)
	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)

	server.Shutdown(tc)
}