package main

import (
	"net/http"
	"project-web/models"
	"strconv"
	"time"

	"github.com/julienschmidt/httprouter"
)

func (app *application) getOneMovie(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())

	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		// app.logger.Print(errors.New("invalid id parameter"))
		app.errorJson(w, err)
		return
	}

	app.logger.Println("ID", id)

	movie := models.Movie{
		ID:          id,
		Title:       "Warkop Dki",
		Description: "Comedy movie",
		Year:        2010,
		ReleaseDate: time.Date(1990, 04, 30, 0, 0, 0, 0, time.Local),
		Runtime:     120,
		Rating:      5,
		MPAARating:  "PG-12",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	err = app.writeJson(w, http.StatusOK, movie, "movie")
}
