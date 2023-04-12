package main

import (
	"net/http"
)

// health handler that returns plain text response with info about app status, os, and version
// this is implemented as a method on our application struct - makes dependencies available to handlers w/o gloval vars or closures
func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {

	env := envelope{
		"status": "available", "system_info": map[string]string{
			"environment": app.config.env,
			"version":     version},
	}

	err := app.writeJSON(w, http.StatusOK, env, nil)
	if err != nil {
		app.logger.Print(err)
		app.serverErrorResponse(w, r, err)
	}
}
