// File: cmd/api/main.go

package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

// the application version number
const version = "1.0.0"

// the cofiguration settings
type config struct {
	port int
	env  string // development, staging, production. etc
	db   struct {
		dsn string
	}
}

// dependency Injection
type application struct {
	config config
	logger *log.Logger
}

func main() {
	var cfg config
	// read in the flags that are needed to populate the our config
	flag.IntVar(&cfg.port, "port", 4000, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (Development | staging| production)")
	flag.StringVar(&cfg.db.dsn, "db-dsn", os.Getenv("TODO_DB_DSN"), "PostgreSQL DSN")
	flag.Parse()

	// Create a logger
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)
	//create a connection pool
	db, err := openDB(cfg)
	if err != nil {
		logger.Fatal(err)
	}

	defer db.Close()
	//Create an instance of our application struct
	app := &application{
		config: cfg,
		logger: logger,
	}

	// Create our new servemux
	mux := http.NewServeMux()
	mux.HandleFunc("/v1/healthcheck", app.healthcheckHandler)
	// Create our http server
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
	//start our server
	logger.Printf("starting %s server on %s", cfg.env, srv.Addr)
	err = srv.ListenAndServe()
	logger.Fatal(err)

}

// The openDB() function returns a *sql.db connection pool
func openDB(cfg config) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.db.dsn)
	if err != nil {
		return nil, err
	}
	// Create a context with a 5-second timeout deadline
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}
	return db, nil
}
