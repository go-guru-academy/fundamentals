package main

import (
	"fmt"

	"github.com/go-guru-academy/fundamentals/messaging-queue/internal/exitmanager"
	"github.com/go-guru-academy/fundamentals/messaging-queue/internal/handlers/hub"
	"github.com/go-guru-academy/fundamentals/messaging-queue/internal/http"
)

func main() {
	fmt.Println("app started")

	// Init the exit manager
	exitManager := exitmanager.Init()

	// Setup the service
	go setup(exitManager)

	// Block until the service is gracefully exited
	exitManager.Wait()
}

func setup(exitManager *exitmanager.ExitManager) {

	// Init hub
	hub.Init()

	// Init HTTP server
	http.Init(exitManager)

}
