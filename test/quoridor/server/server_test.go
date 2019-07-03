package server

import (
	"net/http"
    "net/http/httptest"
	"testing"

	"quoridor/server"
	"quoridor/storage"

	"github.com/gorilla/mux"
)

func setUp() {
	storage.Init()
}

func TestWelcomeAPI(t *testing.T) {
	//Given
    req, err := http.NewRequest("GET", "/", nil)
    if err != nil {
        t.Fatal(err)
    }
    rr := httptest.NewRecorder()
	//When
	http.HandlerFunc(server.WelcomeHandler).ServeHTTP(rr, req)
	//Then
    if status := rr.Code; status != http.StatusOK {
		t.Errorf("Status code differs. Expected %d .\n Got %d instead", http.StatusOK, status)
	}
	body := rr.Body.String()
	expected := "{\"Message\":\"Welcome to the Quoridor API!\"}"
	if body != expected {
		t.Errorf("Body code differs. Expected %v .\n Got %v instead", expected, body)
	}
}

func TestGetGameNotFound(t *testing.T) {
	//Given
	setUp()
    req, err := http.NewRequest("GET", "/games/AszerIOtDE", nil)
    if err != nil {
        t.Fatal(err)
    }
	rr := httptest.NewRecorder()
	r := mux.NewRouter()
	r.HandleFunc("/games/{gameId}", server.GetGameHandler).Methods("GET")
	//When
	r.ServeHTTP(rr, req)
	//Then
    if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("Status code differs. Expected %d .\n Got %d instead", http.StatusOK, status)
	}
	body := rr.Body.String()
	expected := "\n"
	if body != expected {
		t.Errorf("Body code differs. Expected %v .\n Got %v instead", expected, body)
	}
}
