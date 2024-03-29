package request

import (
	"encoding/json"
	"io"
	"net/http"
	"quoridor/game"
	"github.com/gorilla/mux"
)

func GetGameID(r *http.Request) string {
	vars := mux.Vars(r)
	return vars["gameId"]
}

func GetGameConfiguration(r *http.Request) (game.Configuration, error) {
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()
	var conf game.Configuration
	err := decoder.Decode(&conf)
	if err == io.EOF {
		conf = game.Configuration{9, 10}
	} else if err != nil {
		return game.Configuration{}, err
	}
	return conf, nil
}

func GetFence(r *http.Request) (game.Fence, error) {
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()
	var fence game.Fence
	err := decoder.Decode(&fence)
	if err != nil {
		return game.Fence{}, err
	}
	return fence, nil
}

func GetPosition(r *http.Request) (game.Position, error) {
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()
	var position game.Position
	err := decoder.Decode(&position)
	if err != nil {
		return game.Position{}, err
	}
	return position, nil
}
