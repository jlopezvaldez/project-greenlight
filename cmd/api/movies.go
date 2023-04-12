package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/jlopezvaldez/project-greenlight/internal/data"
)

func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {
	//anon struct to hold information expected to be in http request bod
	//struct is target decode destination
	var input struct {
		Title   string   `json:"title"`
		Year    int32    `json:"year"`
		Runtime int32    `json:"runtime"`
		Genres  []string `json:"genres"`
	}

	// Initialize a new json.Decoder instance which reads from the request body, and
	// then use the Decode() method to decode the body contents into the input struct.
	//When calling Decode() you must pass a non-nil pointer as the target decode destination.
	//If you don’t use a pointer, it will return a json.InvalidUnmarshalError error at runtime.
	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	fmt.Fprintf(w, "%+v\n", input)

}

func (app *application) showMovieHandler(w http.ResponseWriter, r *http.Request) {

	//httprouter stores interpolated URL params in the request context
	//use ParamsFromContext() to retrieve a slice containing param names and values
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
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
		app.serverErrorResponse(w, r, err)
	}
}
