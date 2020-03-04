package main

import (
	"fmt"

	"github.com/go-guru-academy/fundamentals/messaging-queue/internal/http"
)

func main() {
	fmt.Println("app started")
	http.Init()
}
