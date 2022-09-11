package render

import (
	"github.com/go-chi/render"
	"net/http"
	"strings"
	"tasks-api/internal/apperror"
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

func NewAppError(w http.ResponseWriter, r *http.Request, err error, mapErrors map[apperror.ErrCode]int) {
	NewError(w, r, err.Error(), apperror.GetStatusCode(err, mapErrors))
}

func JSON(w http.ResponseWriter, r *http.Request, v interface{}) {
	render.JSON(w, r, v)
}
