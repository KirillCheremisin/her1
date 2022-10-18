package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi"

	"workshop/internal/api/jokes"
	"workshop/internal/handler"
)

func main() {
	apiClient := jokes.NewJokeClient("https://geek-jokes.sameerkumar.website/")

	h := handler.NewHandler(apiClient)

	r := chi.NewRouter()

	r.Get("/hello", h.Hello)

	srv := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: r,
	}

	quit := make(chan os.Signal, 1)
	done := make(chan error, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-quit
		ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
		defer cancel()
		err := srv.Shutdown(ctx)
		// ...
		done <- err
	}()

	log.Print("starting server")
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

	err = <-done

	log.Printf("shutting server down with %v", err)
}
