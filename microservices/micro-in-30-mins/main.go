package main

import (
	"log"
	"micro-in-30-mins/homepage"
	"micro-in-30-mins/server"
	"net/http"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var (
	CertFile = os.Getenv("CERT_FILE")
	KeyFile = os.Getenv("KEY_FILE")
	ServiceAddr = os.Getenv("SERVICE_ADDR")
)

func main() {
	logger := log.New(os.Stdout, "gc ", log.LstdFlags | log.Lshortfile)

	db, err := sqlx.Open("postgres", "postgres://postgres:postgres@localhost:5432/postgres")
	if err != nil {
		logger.Fatalln(err)
	}

	err = db.Ping()
	if err != nil {
		logger.Fatalln(err)
	}

	h := homepage.NewHandlers(logger, db)

	mux := http.NewServeMux()
	h.SetupRoutes(mux)

	srv := server.New(mux, ServiceAddr)

	logger.Println("server starting")
	err = srv.ListenAndServeTLS(CertFile, KeyFile)
	if err != nil {
		logger.Fatalf("server failed to start: %v", err)
	}
}

