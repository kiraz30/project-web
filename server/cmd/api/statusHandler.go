package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (app *application) statusHandler(w http.ResponseWriter, r *http.Request) {
	currentStatus := AppStatus{
		Status:      "Online",
		Environment: app.config.env,
		Version:     version,
	}

	res, err := json.MarshalIndent(currentStatus, "", "\t")
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}
