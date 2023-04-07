package main

import (
	"fmt"
	"net/http"
)

// health handler that returns plain text response with info about app status, os, and version
// this is implemented as a method on our application struct - makes dependencies available to handlers w/o gloval vars or closures
func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "status: available")
	fmt.Fprintf(w, "environment: %s\n", app.config.env) //environmaent name is retrieved from the app struct
	fmt.Fprintf(w, "version: %s\n", version)
}
