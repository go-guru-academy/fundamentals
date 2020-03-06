package http

import (
	"net/http"

	"github.com/go-guru-academy/fundamentals/messaging-queue/internal/exitmanager"
	"github.com/go-guru-academy/fundamentals/messaging-queue/internal/handlers/routes"
)

type Server struct {
	_           struct{}
	ExitManager *exitmanager.ExitManager
}

func Init(e *exitmanager.ExitManager) {
	s := &Server{
		ExitManager: e,
	}
	go s.run()
}

func (s *Server) run() {

	// Create routes
	r := routes.New()

	// Listen and serve
	if err := http.ListenAndServe(":8888", r); err != nil {
		s.ExitManager.ServerError(err)
		return
	}

}
