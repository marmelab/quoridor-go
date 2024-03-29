package server

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"quoridor/controller"
	"quoridor/game"
	"quoridor/server/request"
	"quoridor/server/response"

	"github.com/gorilla/mux"
	"github.com/lithammer/shortuuid"
)

// Port is the default server port
const (
	Port = 8383
	AuthorizationHeaderName = "Authorization"
)

type Message struct {
	Message string
}

type AuthorizationToken struct {
	AuthToken string
}

// Start launch the server
func Start() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", welcomeHandler).Methods("GET")
	router.HandleFunc("/games", CreateGameHandler).Methods("POST")
	router.HandleFunc("/games/{gameId}", getGameHandler).Methods("GET")
	router.HandleFunc("/games/{gameId}/join", joinGameHandler).Methods("PUT")
	router.HandleFunc("/games/{gameId}/add-fence", addFenceHandler).Methods("PUT")
	router.HandleFunc("/games/{gameId}/add-fence/possibilities", getFencePossibilitiesHandler).Methods("GET")
	router.HandleFunc("/games/{gameId}/move-pawn", movePawnHandler).Methods("PUT")
	router.HandleFunc("/games/{gameId}/move-pawn/possibilities", getMovePossibilitiesHandler).Methods("GET")
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

func CreateGameHandler(w http.ResponseWriter, r *http.Request) {
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
	sendGameRepresentation(w, r, *game)
}

func getGameHandler(w http.ResponseWriter, r *http.Request) {
	id := request.GetGameID(r)
	game, err := gamecontroller.GetGame(id)
	if err != nil {
		response.SendBadRequestError(w, err)
		return
	}
	sendGameRepresentation(w, r, game)
}

func joinGameHandler(w http.ResponseWriter, r *http.Request) {
	id := request.GetGameID(r)
	authToken:= shortuuid.New()
	err := gamecontroller.JoinGame(id, authToken)
	if err != nil {
		response.SendBadRequestError(w, err)
		return
	}
	response.SendOK(w, AuthorizationToken{authToken})
}

func addFenceHandler(w http.ResponseWriter, r *http.Request) {
	id := request.GetGameID(r)
	fence, err := request.GetFence(r)
	if err != nil {
		response.SendBadRequestError(w, err)
		return
	}
	authToken := r.Header.Get(AuthorizationHeaderName)
	game, err := gamecontroller.AddFence(id, fence, authToken)
	if err != nil {
		response.SendBadRequestError(w, err)
		return
	}
	sendGameRepresentation(w, r, game)
}

func getFencePossibilitiesHandler(w http.ResponseWriter, r *http.Request) {
	id := request.GetGameID(r)
	possibilities, err := gamecontroller.GetFencePossibilities(id)
	if err != nil {
		response.SendBadRequestError(w, err)
		return
	}
	response.SendOK(w, possibilities)
}

func movePawnHandler(w http.ResponseWriter, r *http.Request) {
	id := request.GetGameID(r)
	to, err := request.GetPosition(r)
	if err != nil {
		response.SendBadRequestError(w, err)
		return
	}
	authToken := r.Header.Get(AuthorizationHeaderName)
	game, err := gamecontroller.MovePawn(id, to, authToken)
	if err != nil {
		response.SendBadRequestError(w, err)
		return
	}
	sendGameRepresentation(w, r, game)
}

func getMovePossibilitiesHandler(w http.ResponseWriter, r *http.Request) {
	id := request.GetGameID(r)
	possibilities, err := gamecontroller.GetMovePossibilities(id)
	if err != nil {
		response.SendBadRequestError(w, err)
		return
	}
	response.SendOK(w, possibilities)
}

func sendGameRepresentation(w http.ResponseWriter, r *http.Request, game game.Game) {
	accept := r.Header.Get("Accept")
	if accept == "text/plain" {
		response.SendPlainOK(w, game.GetTextBoard())
		return
	}
	response.SendOK(w, game)
}
