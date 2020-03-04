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

func writeSuccess(w http.ResponseWriter) {
	writeJsonResponse(w, 200, &SimpleResponse{
		Code:    http.StatusOK,
		Message: SUCCESS,
	})
}

func writeJsonResponse(w http.ResponseWriter, status int, data interface{}) {
	j, err := json.Marshal(data)
	if err != nil {
		writeServerError(w, err)
		return
	}
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	w.Write(j)
}

func writeServerError(w http.ResponseWriter, err error) {
	fmt.Println(err)
	j, _ := json.Marshal(&SimpleResponse{
		Code:    http.StatusInternalServerError,
		Message: "Internal Server Error",
	})
	w.WriteHeader(http.StatusInternalServerError)
	w.Header().Set("Content-Type", "application/json")
	w.Write(j)
}
