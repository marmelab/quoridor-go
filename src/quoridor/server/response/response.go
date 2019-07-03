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

func SendPlainOK(w http.ResponseWriter, message string) {
	w.Header().Set("content-type", "text/plain")
	w.Write([]byte(message))
}

func SendBadRequestError(w http.ResponseWriter, err error) {
	SendBadRequest(w, err.Error())
}

func SendBadRequest(w http.ResponseWriter, message string) {
	http.Error(w, "{ \"message\": \""+message+"\"}", http.StatusBadRequest)
}

func SendNotFound(w http.ResponseWriter) {
	http.Error(w, "", http.StatusNotFound)
}
