package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"quoridor/game"

	"github.com/gorilla/mux"
)

// Port is the default server port
const Port = 8383

// Start launch the server
func Start() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", welcome).Methods("GET")
	router.HandleFunc("/games", createGame).Methods("POST")
	port := getListeningPort()
	fmt.Printf("Server started on port: %v\n", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}

func getListeningPort() string {
	return strconv.Itoa(Port)
}

func welcome(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	w.Write([]byte(`{"message": "Welcome to the Quoridor API!"}`))
}

func createGame(w http.ResponseWriter, r *http.Request) {
	configuration := getConfiguration(r)
	newGame := game.CreateGame(configuration)
	sendResponse(w, newGame)
}

func sendResponse(w http.ResponseWriter, response interface{}) {
	encodedResponse, err := json.Marshal(response)
	if err != nil {
		panic(err)
	}
	w.Header().Set("content-type", "application/json")
	w.Write([]byte(string(encodedResponse)))
}

func getConfiguration(r *http.Request) game.Configuration {
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()
	var conf game.Configuration
	err := decoder.Decode(&conf)
	if err != nil {
		conf = game.Configuration{9}
	}
	return conf
}
