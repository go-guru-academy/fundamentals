package routes

import (
	"github.com/go-guru-academy/fundamentals/messaging-queue/internal/handlers"
	"github.com/go-guru-academy/fundamentals/messaging-queue/internal/handlers/hub"
	"github.com/go-guru-academy/fundamentals/messaging-queue/internal/http/middleware"
	"github.com/gorilla/mux"
)

func New(h *hub.Hub) *mux.Router {
	m := &middleware.Middleware{
		Hub: h,
	}
	r := mux.NewRouter()
	r.HandleFunc(m.First("/message", handlers.CreateMessage)).Methods("POST")
	r.HandleFunc(m.First("/message", handlers.GetMessage)).Methods("GET")
	return r
}
