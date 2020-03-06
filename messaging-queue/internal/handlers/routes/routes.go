package routes

import (
	"github.com/go-guru-academy/fundamentals/messaging-queue/internal/handlers"
	"github.com/gorilla/mux"
)

func New() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/message", handlers.CreateMessage).Methods("POST")
	r.HandleFunc("/message", handlers.GetMessage).Methods("GET")
	return r
}
