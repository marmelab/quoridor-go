package server

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Port is the default server port
const Port = 8383

// Start launch the server
func Start() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", welcome).Methods("GET")
	port := getListeningPort()
	fmt.Println("Server started on port : " + port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}

func getListeningPort() string {
	return strconv.Itoa(Port)
}

func welcome(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	w.Write([]byte(`{"message": "Welcome to the Quoridor API!"}`))
}
