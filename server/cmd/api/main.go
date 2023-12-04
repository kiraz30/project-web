package main

import (
	// "encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

const version = "1.0.0"

type config struct {
	port int
	env  string
}

type AppStatus struct {
	Status      string `json:"status"`
	Environment string `json:"environment"`
	Version     string `json:"version"`
}

type application struct {
	config config
	logger *log.Logger
}

func main() {
	var cfg config

	flag.IntVar(&cfg.port, "port", 4000, "server port to listen")
	flag.StringVar(&cfg.env, "env", "development", "Aplication environment (development || production)")
	flag.Parse()

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	app := &application{
		config: cfg,
		logger: logger,
	}

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
	err := serve.ListenAndServe()
	if err != nil {
		log.Println(err)
	}

}
