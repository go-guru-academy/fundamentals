package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/go-guru-academy/fundamentals/messaging-queue/internal/handlers/hub"
	"github.com/go-guru-academy/fundamentals/messaging-queue/internal/models"
)

type Handler struct {
	ResponseWriter http.ResponseWriter
	Request        *http.Request
	Hub            *hub.Hub
}

func CreateMessage(h *Handler) {
	fmt.Println("post: CreateMessage")

	// Read the request body
	body, err := ioutil.ReadAll(h.Request.Body)
	if err != nil {
		h.writeServerError(err)
		return
	}

	// Validate that the body is valid JSON
	if err := json.Unmarshal(body, &models.Message{}); err != nil {
		h.writeJsonResponse(http.StatusBadRequest, &SimpleResponse{
			Code:    http.StatusBadRequest,
			Message: "Invalid Message Format. Must be valid JSON.",
		})
		return
	}

	if err := h.Hub.Enqueue(body); err != nil {
		h.writeJsonResponse(http.StatusTooManyRequests, &SimpleResponse{
			Code:    http.StatusTooManyRequests,
			Message: "Message Limit Exceeded.",
		})
		return
	}

	h.writeSuccess()

}

func GetMessage(h *Handler) {
	fmt.Println("get: GetMessage")
	message, err := h.Hub.Dequeue()
	if err != nil {
		h.writeJsonResponse(http.StatusOK, &SimpleResponse{
			Code:    http.StatusOK,
			Message: "No messages in queue.",
		})
		return
	}
	messageS := string(message)
	h.writeJsonResponse(http.StatusOK, &messageS)
}
