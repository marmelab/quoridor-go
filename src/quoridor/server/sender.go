package server

import (
	"encoding/json"
	"net/http"
)

func sendResponse(w http.ResponseWriter, response interface{}) {
	encodedResponse, err := json.Marshal(response)
	if err != nil {
		panic(err)
	}
	w.Header().Set("content-type", "application/json")
	w.Write([]byte(string(encodedResponse)))
}

func sendBadRequestError(w http.ResponseWriter, err error) {
	sendBadRequestResponse(w, err.Error())
}

func sendBadRequestResponse(w http.ResponseWriter, message string) {
	http.Error(w, "{ \"message\": \"" + message + "\"}", http.StatusBadRequest)
}
