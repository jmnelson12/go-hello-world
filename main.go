package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

type HelloWorldResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func main() {
	// handlers
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")

		e := json.NewEncoder(w)
		err := e.Encode(&HelloWorldResponse{Status: "ok", Message: "Hello World"})
		if err != nil {
			// we should never be here but log the error just incase
			fmt.Println("[error] unable to decode message", err)
		}
	})

	// create a new server
	port, exists := os.LookupEnv("PORT")
	if !exists {
		port = "8080"
	}

	s := &http.Server{
		Addr:         fmt.Sprintf(":%s", port),
		ReadTimeout:  5 * time.Second,   // max time to read request from the client
		WriteTimeout: 10 * time.Second,  // max time to write response to the client
		IdleTimeout:  120 * time.Second, // max time for connections using TCP Keep-Alive
	}

	// start the server
	go func() {
		fmt.Printf("[info] starting server on port %s\n", port)

		err := s.ListenAndServe()
		if err != nil {
			fmt.Println("[error] starting server", err)
			os.Exit(1)
		}
	}()

	// trap sigterm or interupt and gracefully shutdown the server
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	// Block until a signal is received.
	sig := <-c
	log.Println("Got signal:", sig)

	// gracefully shutdown the server, waiting max 30 seconds for current operations to complete
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(ctx)

}
