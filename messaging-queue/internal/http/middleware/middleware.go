package middleware

import (
	"net/http"

	"github.com/go-guru-academy/fundamentals/messaging-queue/internal/handlers"
	"github.com/go-guru-academy/fundamentals/messaging-queue/internal/handlers/hub"
)

type Middleware struct {
	Hub *hub.Hub
}

func (m *Middleware) First(pattern string, handler func(*handlers.Handler)) (string, http.HandlerFunc) {
	return pattern, func(w http.ResponseWriter, r *http.Request) {
		h := &handlers.Handler{
			ResponseWriter: w,
			Request:        r,
			Hub:            m.Hub,
		}
		handler(h)
	}
}
