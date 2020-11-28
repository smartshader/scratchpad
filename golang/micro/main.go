package main

import (
	"context"
	"log"
	"micro/handlers"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	logger := log.New(os.Stdout, "product-api", log.LstdFlags | log.Lshortfile)

	productsHandler := handlers.NewProducts(logger)

	sm := http.NewServeMux()
	sm.Handle("/", productsHandler)

	server := http.Server{
		Addr: ":9090",
		Handler: sm,
		IdleTimeout: 120 * time.Second,
		ReadTimeout: 1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		logger.Printf("starting server on port 9090")
		err := server.ListenAndServe()
		if err != nil {
			logger.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <- sigChan
	logger.Println("received terminate, graceful shutdown", sig)

	tc, _ := context.WithTimeout(context.Background(), 30 * time.Second)
	server.Shutdown(tc)
}
