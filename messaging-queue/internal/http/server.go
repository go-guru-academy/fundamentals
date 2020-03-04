package http

import (
	"net/http"

	"github.com/go-guru-academy/fundamentals/messaging-queue/internal/handlers"
)

func Init() {
	run()
}

func run() {
	http.HandleFunc("/message", handlers.CreateMessage)
	http.ListenAndServe(":8888", nil)
}
