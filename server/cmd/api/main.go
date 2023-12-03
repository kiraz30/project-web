package main

import (
	"flag"
	"fmt"
	"net/http"
)

const version = "1.0.0"

type config struct {
	port int
	env  string
}

type AppStatus struct {
	Status      string `json:"status`
	Environment string `json:"environment"`
	Version     string `json:"version`
}

func main() {
	var cfg config

	flag.IntVar(&cfg.port, "port", 4000, "server port to listen")
	flag.StringVar(&cfg.env, "env", "development", "Aplication environment (development || production)")
	flag.Parse()

	fmt.Println("Server is running...")

	//routing
	http.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {})
}
