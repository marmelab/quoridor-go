package server

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"quoridor/controller"
	"quoridor/server/request"
	"quoridor/server/response"

	"github.com/gorilla/mux"
)

// Port is the default server port
const Port = 8383

type Message struct {
	Message string
}

// Start launch the server
func Start() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", welcomeHandler).Methods("GET")
	router.HandleFunc("/games", createGameHandler).Methods("POST")
	router.HandleFunc("/games/{gameId}", getGameHandler).Methods("GET")
	router.HandleFunc("/games/{gameId}/add-fence", addFenceHandler).Methods("PUT")
	router.HandleFunc("/games/{gameId}/add-fence/possibilities", addFencePossibilitiesHandler).Methods("GET")
	port := getListeningPort()
	fmt.Printf("Server started on port: %v\n", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}

func getListeningPort() string {
	return strconv.Itoa(Port)
}

func welcomeHandler(w http.ResponseWriter, r *http.Request) {
	response.SendOK(w, Message{"Welcome to the Quoridor API!"})
}

func createGameHandler(w http.ResponseWriter, r *http.Request) {
	configuration, err := request.GetGameConfiguration(r)
	if err != nil {
		response.SendBadRequestError(w, err)
		return
	}
	game, err := gamecontroller.CreateGame(configuration)
	if err != nil {
		response.SendBadRequestError(w, err)
		return
	}
	response.SendOK(w, game)
}

func getGameHandler(w http.ResponseWriter, r *http.Request) {
	id := request.GetGameID(r)
	game, err := gamecontroller.GetGame(id)
	if err != nil {
		response.SendBadRequestError(w, err)
		return
	}
	response.SendOK(w, game)
}

func addFenceHandler(w http.ResponseWriter, r *http.Request) {
	id := request.GetGameID(r)
	fence, err := request.GetFence(r)
	if err != nil {
		response.SendBadRequestError(w, err)
		return
	}
	game, err := gamecontroller.AddFence(id, fence)
	if err != nil {
		response.SendBadRequestError(w, err)
		return
	}
	response.SendOK(w, game)
}

func addFencePossibilitiesHandler(w http.ResponseWriter, r *http.Request) {
	id := request.GetGameID(r)
	possibilities, err := gamecontroller.AddFencePossibilities(id)
	if err != nil {
		response.SendBadRequestError(w, err)
		return
	}
	response.SendOK(w, possibilities)
}
