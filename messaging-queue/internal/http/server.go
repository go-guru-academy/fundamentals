package http

import (
	"net/http"

	"github.com/go-guru-academy/fundamentals/messaging-queue/internal/exitmanager"
	"github.com/go-guru-academy/fundamentals/messaging-queue/internal/handlers/hub"
	"github.com/go-guru-academy/fundamentals/messaging-queue/internal/handlers/routes"
)

type Server struct {
	_           struct{}
	ExitManager *exitmanager.ExitManager
	Hub         *hub.Hub
}

func Init(e *exitmanager.ExitManager, h *hub.Hub) {
	s := &Server{
		ExitManager: e,
		Hub:         h,
	}
	go s.run()
}

func (s *Server) run() {

	// Create routes
	r := routes.New(s.Hub)

	// Listen and serve
	if err := http.ListenAndServe(":8888", r); err != nil {
		s.ExitManager.ServerError(err)
		return
	}

}
