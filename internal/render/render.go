package render

import (
	"github.com/go-chi/render"
	"net/http"
	"strings"
)

type ErrorResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func NewError(w http.ResponseWriter, r *http.Request, message string, statusCode int) {
	err := ErrorResponse{
		Code:    strings.ReplaceAll(strings.ToLower(http.StatusText(statusCode)), " ", "_"),
		Message: message,
	}
	w.WriteHeader(statusCode)
	JSON(w, r, err)
}

func JSON(w http.ResponseWriter, r *http.Request, v interface{}) {
	render.JSON(w, r, v)
}
