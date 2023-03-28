// Filename: ./cmd/api/helpers.go
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

type envelope map[string]interface{} //envelope is a type that will contain school infos //container inside a container

func (app *application) readIDParams(r *http.Request) (int64, error) {
	params := httprouter.ParamsFromContext(r.Context())
	id, err := strconv.ParseInt(params.ByName("id"), 10, 64)
	if err != nil || id < 1 {
		return 0, errors.New("invalid Id Parameter")
	}
	return id, nil
}

func (app *application) WriteJSON(w http.ResponseWriter, status int, data envelope, headers http.Header) error {
	//convert the data to json format
	js, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return err
	}
	js = append(js, '\n') //if no error
	//add any headers that were sent
	for key, value := range headers {
		w.Header()[key] = value
	}

	//add header info
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write([]byte(js))
	return nil

}


