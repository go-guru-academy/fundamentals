package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	SUCCESS = "success"
)

type SimpleResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (h *Handler) writeSuccess() {
	h.writeJsonResponse(200, &SimpleResponse{
		Code:    http.StatusOK,
		Message: SUCCESS,
	})
}

func (h *Handler) writeJsonResponse(status int, data interface{}) {
	j, err := json.Marshal(data)
	if err != nil {
		h.writeServerError(err)
		return
	}
	h.ResponseWriter.WriteHeader(status)
	h.ResponseWriter.Header().Set("Content-Type", "application/json")
	h.ResponseWriter.Write(j)
}

func (h *Handler) writeServerError(err error) {
	fmt.Println(err)
	j, _ := json.Marshal(&SimpleResponse{
		Code:    http.StatusInternalServerError,
		Message: "Internal Server Error",
	})
	h.ResponseWriter.WriteHeader(http.StatusInternalServerError)
	h.ResponseWriter.Header().Set("Content-Type", "application/json")
	h.ResponseWriter.Write(j)
}
