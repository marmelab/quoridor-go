package server

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
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
	router.HandleFunc("/", welcomeHandler).Methods("GET")
	router.HandleFunc("/games", createGameHandler).Methods("POST")
	port := getListeningPort()
	fmt.Printf("Server started on port: %v\n", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}

func getListeningPort() string {
	return strconv.Itoa(Port)
}

func welcomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	w.Write([]byte(`{"message": "Welcome to the Quoridor API!"}`))
}

func createGameHandler(w http.ResponseWriter, r *http.Request) {
	configuration, err := getConfiguration(r)
	if err != nil {
		send400Error(w, err)
	}
	newGame, err := game.CreateGame(configuration)
	if err != nil {
		send400Error(w, err)
	} else {
		sendResponse(w, newGame)
	}
}

func sendResponse(w http.ResponseWriter, response interface{}) {
	encodedResponse, err := json.Marshal(response)
	if err != nil {
		panic(err)
	}
	w.Header().Set("content-type", "application/json")
	w.Write([]byte(string(encodedResponse)))
}

func send400Error(w http.ResponseWriter, err error) {
	send400Response(w, err.Error())
}

func send400Response(w http.ResponseWriter, message string) {
	http.Error(w, "{ \"message\": \"" + message + "\"}", http.StatusBadRequest)
}

func getConfiguration(r *http.Request) (*game.Configuration, error) {
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()
	var conf game.Configuration
	err := decoder.Decode(&conf)
	if err == io.EOF {
		conf = game.Configuration{9}
	} else if err != nil {
		return nil, errors.New(err.Error())
	}
	return &conf, nil
}
