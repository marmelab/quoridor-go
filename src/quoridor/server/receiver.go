package server

import (
	"encoding/json"
	"io"
	"net/http"

	"quoridor/game"

	"github.com/gorilla/mux"
)

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