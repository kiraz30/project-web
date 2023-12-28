package main

import (
	// "encoding/json"
	"context"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"project-web/models"
	"time"

	_ "github.com/lib/pq"
)

const version = "1.0.0"

type config struct {
	port int
	env  string
	db   struct {
		dsn string
	}
}

type AppStatus struct {
	Status      string `json:"status"`
	Environment string `json:"environment"`
	Version     string `json:"version"`
}

type application struct {
	config config
	logger *log.Logger
	models models.Models
}

func main() {
	var cfg config

	flag.IntVar(&cfg.port, "port", 4000, "server port to listen")
	flag.StringVar(&cfg.env, "env", "development", "Aplication environment (development || production)")
	flag.StringVar(&cfg.db.dsn, "dsn", "postgres://postgres@localhost/id_gomovie?sslmode=disable", "Postgres connection config")
	flag.Parse()

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	db, err := openDB(cfg)
	if err != nil {
		logger.Fatal(err)
	}
	app := &application{
		config: cfg,
		logger: logger,
		models: models.NewModels(db),
	}

	//blocking
	defer db.Close()

	fmt.Println("Server is running...")
	//	setting limit timeout
	serve := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,      // 1 minunte
		ReadTimeout:  10 * time.Second, // 10 second
		WriteTimeout: 30 * time.Second, // 30 second
	}

	logger.Printf("Strarting serve on port ", cfg.port)
	err = serve.ListenAndServe()
	if err != nil {
		log.Println(err)
	}

}

func openDB(cfg config) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.db.dsn)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	return db, nil
}
