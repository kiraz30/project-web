package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
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

func main() {
	var cfg config

	flag.IntVar(&cfg.port, "port", 4000, "server port to listen")
	flag.StringVar(&cfg.env, "env", "development", "Aplication environment (development || production)")
	flag.Parse()

	fmt.Println("Server is running...")

	//routing
	http.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		currentStatus := AppStatus{
			Status:      "Online",
			Environment: cfg.env,
			Version:     version,
		}

		res, err := json.MarshalIndent(currentStatus, "", "\t")
		if err != nil {
			fmt.Println(err)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(res)
	})

	err := http.ListenAndServe(fmt.Sprintf(":%d", cfg.port), nil)
	if err != nil {
		log.Println(err)
	}
}
