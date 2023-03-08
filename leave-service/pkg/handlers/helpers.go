package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

// readJSON tries to read the body of a request and converts it into JSON
func (rep *Repository) readJSON(w http.ResponseWriter, r *http.Request, data any) error {
	maxBytes := 1048576 // one megabyte

	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))

	dec := json.NewDecoder(r.Body)
	err := dec.Decode(data)
	if err != nil {
		return err
	}

	log.Println("Data: ", data)

	err = dec.Decode(&struct{}{})
	if err != io.EOF {
		return errors.New("body must have only a single JSON value")
	}

	return nil
}

// writeJSON takes a response status code and arbitrary data and writes a json response to the client
func (rep *Repository) writeJSON(w http.ResponseWriter, status int, data any) error {
	out, err := json.Marshal(data)
	if err != nil {
		return err
	}

	// Set header
	w.Header().Set("Content-Type", "application/json")
	// Write header
	w.WriteHeader(status)
	_, err = w.Write(out)
	if err != nil {
		return err
	}

	return nil
}

// errorJSON takes an error, and optionally a response status code, and generates and sends
// a json error response
func (rep *Repository) errorJSON(w http.ResponseWriter, err error, status ...int) error {
	statusCode := http.StatusBadRequest

	if len(status) > 0 {
		statusCode = status[0]
	}

	var payload jsonResponse
	payload.Error = true
	payload.Message = err.Error()

	return rep.writeJSON(w, statusCode, payload)
}

// function to check if directory exists or not and return false; if exist it will return true
func directoryExist(dir string) (bool, error) {

	//
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		log.Printf("Directory %s does not exist\n", dir)
		return false, nil
	} else if err != nil {
		log.Printf("Error checking if directory %s exists, Please check : %v\n", dir, err)
		return false, err
	} else {
		fmt.Printf("Directory %s exists\n", dir)
		return true, nil
	}
}

// This function will attemps to check directory exist, if not create the directory
func (rep *Repository) Mkdir(dir string) (string, error) {

	// Check directory exist

	exist, err := directoryExist(dir)
	if err != nil {
		log.Println("Error checking directory exist, Please check for permission : ", err)
		return dir, err
	}

	if !exist {
		// create directory
		err := os.Mkdir(dir, 0755)
		if err != nil {
			log.Println("Error creating directory, Please check : ", err)
			return dir, err
		}
	}

	return dir, nil
}
