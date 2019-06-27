package server

import (
	"encoding/json"
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
	router.HandleFunc("/games/{gameId}", getGameHandler).Methods("GET")
	router.HandleFunc("/games/{gameId}/add-fence", addFenceHandler).Methods("PUT")
	port := getListeningPort()
	fmt.Printf("Server started on port: %v\n", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}

func getListeningPort() string {
	return strconv.Itoa(Port)
}

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

func getGameID(r *http.Request) string {
	vars := mux.Vars(r)
	return vars["gameId"]
}

func getConfiguration(r *http.Request) (*game.Configuration, error) {
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()
	var conf game.Configuration
	err := decoder.Decode(&conf)
	if err == io.EOF {
		conf = game.Configuration{9}
	} else if err != nil {
		return nil, err
	}
	return &conf, nil
}

func getFence(r *http.Request) (game.Fence, error) {
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()
	var fence game.Fence
	err := decoder.Decode(&fence)
	if err != nil {
		return game.Fence{}, err
	}
	return fence, nil
}

func welcomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	w.Write([]byte(`{"message": "Welcome to the Quoridor API!"}`))
}

func createGameHandler(w http.ResponseWriter, r *http.Request) {
	configuration, err := getConfiguration(r)
	if err != nil {
		sendBadRequestError(w, err)
		return
	}
	newGame, err := game.CreateGame(configuration)
	if err != nil {
		sendBadRequestError(w, err)
		return
	}
	sendResponse(w, newGame)
}

func getGameHandler(w http.ResponseWriter, r *http.Request) {
	id := getGameID(r)
	game, err := game.GetGame(id)
	if err != nil {
		sendBadRequestError(w, err)
	} else {
		sendResponse(w, game)
	}
}

func addFenceHandler(w http.ResponseWriter, r *http.Request) {
	id := getGameID(r)
	fence, err := getFence(r)
	if err != nil {
		sendBadRequestError(w, err)
		return
	}
	game, err := game.AddFence(id, fence)
	if err != nil {
		sendBadRequestError(w, err)
		return
	}
	sendResponse(w, game)
}
