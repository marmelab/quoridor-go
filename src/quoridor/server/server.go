package server

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"quoridor/service"

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
	game, err := service.CreateGame(configuration)
	if err != nil {
		sendBadRequestError(w, err)
		return
	}
	sendResponse(w, game)
}

func getGameHandler(w http.ResponseWriter, r *http.Request) {
	id := getGameID(r)
	game, err := service.GetGame(id)
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
	game, err := service.AddFence(id, fence)
	if err != nil {
		sendBadRequestError(w, err)
		return
	}
	sendResponse(w, game)
}
