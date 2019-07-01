package response

import (
	"encoding/json"
	"net/http"
)

func SendOK(w http.ResponseWriter, response interface{}) {
	encodedResponse, err := json.Marshal(response)
	if err != nil {
		panic(err)
	}
	w.Header().Set("content-type", "application/json")
	w.Write([]byte(string(encodedResponse)))
}

func SendBadRequestError(w http.ResponseWriter, err error) {
	SendBadRequest(w, err.Error())
}

func SendBadRequest(w http.ResponseWriter, message string) {
	http.Error(w, "{ \"message\": \"" + message + "\"}", http.StatusBadRequest)
}
