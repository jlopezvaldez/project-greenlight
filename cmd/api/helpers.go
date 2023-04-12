package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type envelope map[string]any

//retrieve id url param from request context
//convert to int and return it, return 0 and an error when unsuccessful
//helper method to parse URL param

func (app *application) readIDParam(r *http.Request) (int64, error) {
	params := httprouter.ParamsFromContext(r.Context())

	id, err := strconv.ParseInt(params.ByName("id"), 10, 64)
	if err != nil || id < 1 {
		return 0, errors.New("Invalid ID parameter.")
	}
	return id, nil
}

func (app *application) writeJSON(w http.ResponseWriter, status int, data envelope, headers http.Header) error {
	//encode data to json
	js, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return err
	}

	js = append(js, '\n')

	for key, value := range headers {
		w.Header()[key] = value
	}

	//add json header
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)

	return nil
}

func (app *application) readJSON(w http.ResponseWriter, r *http.Request, dst any) error {
	//decode request body into target
	err := json.NewDecoder(r.Body).Decode(dst)
	if err != nil {
		//if there is error begin triage
		var syntaxError *json.SyntaxError
		var unmarshallTypeError *json.UnmarshalTypeError
		var invalidUnmarshallError *json.InvalidUnmarshalError

		switch {
		//check if error type is syntax
		case errors.As(err, &syntaxError):
			return fmt.Errorf("body contains badly-formed JSON (at character %d)", syntaxError.Offset)

		//in some cases syntax errors return io.ErrUnexpectedEOF
		//check issue -> https://github.com/golang/go/issues/25956
		case errors.Is(err, io.ErrUnexpectedEOF):
			return errors.New("body contains badly-formed JSON")

		//check for wrong type in target destination
		//if error relates to specific field include that in message
		case errors.As(err, &unmarshallTypeError):
			if unmarshallTypeError.Field != "" {
				return fmt.Errorf("body contains incorrect JSON type for field %q", unmarshallTypeError.Field)
			}
			return fmt.Errorf("body contains incorrect JSON type (at character %d)", unmarshallTypeError.Offset)

		//EOF error if body is empty return message for that
		case errors.Is(err, io.EOF):
			return errors.New("body must not be empty")

		//panic if a non-nil pointer is passed to Decode()
		case errors.As(err, &invalidUnmarshallError):
			panic(err)

			//anything else return vanilla message
		default:
			return err
		}
	}
	return nil
}
