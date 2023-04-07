package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/jlopezvaldez/project-greenlight/internal/data"
)

func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "create a new movie")
}

func (app *application) showMovieHandler(w http.ResponseWriter, r *http.Request) {

	//httprouter stores interpolated URL params in the request context
	//use ParamsFromContext() to retrieve a slice containing param names and values
	id, err := app.readIDParam(r)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	movie := data.Movie{
		ID:        id,
		CreatedAt: time.Now(),
		Title:     "Blade Runner: 2049",
		Runtime:   242,
		Genres:    []string{"sci-fi", "action"},
		Version:   1,
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"movie": movie}, nil)
	if err != nil {
		app.logger.Print(err)
		http.Error(w, "Server encountered a problem.", http.StatusInternalServerError)
	}
}
