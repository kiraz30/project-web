package main

import (
	"encoding/json"
	"net/http"
)

func (app *application) writeJson(w http.ResponseWriter, status int, data interface{}, wrap string) error {
	wrapper := make(map[string]interface{})
	wrapper[wrap] = data

	res, err := json.Marshal(wrapper)
	if err != nil {
		app.logger.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
	return nil
}

// function errorJson dengan receiver app
func (app *application) errorJson(w http.ResponseWriter, err error) {
	type jsonError struct {
		Message string `json:"message`
	}

	errMessage := jsonError{
		Message: err.Error(),
	}

	app.writeJson(w, http.StatusBadRequest, errMessage, "error_message")
}
