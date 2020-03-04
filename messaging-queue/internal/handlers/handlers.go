package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/go-guru-academy/fundamentals/messaging-queue/internal/models"
)

func CreateMessage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("message received")

	// Read the request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		writeServerError(w, err)
		return
	}

	// Validate that the body is valid JSON
	if err := json.Unmarshal(body, &models.Message{}); err != nil {
		writeJsonResponse(w, http.StatusBadRequest, &SimpleResponse{
			Code:    http.StatusBadRequest,
			Message: "Invalid Message Format. Must be valid JSON.",
		})
		return
	}

	writeSuccess(w)

}
