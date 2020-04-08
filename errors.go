package yuri

import "net/http"

var ErrNotFound = &ErrResponse{HTTPStatusCode: 404, StatusText: "Item not found."}
var ErrInvalidRequest = &ErrResponse{HTTPStatusCode: http.StatusBadRequest, StatusText: "Invalid request."}
var ErrInternalServerError = &ErrResponse{HTTPStatusCode: http.StatusInternalServerError, StatusText: "Internal Server Error"}
var ErrUnauthorized = &ErrResponse{HTTPStatusCode: http.StatusUnauthorized, StatusText: "Authorization Error"}

type ErrResponse struct {
	HTTPStatusCode int    `json:"http_status_code"`
	Message        string `json:"message"`
	StatusText     string `json:"status"`
}
