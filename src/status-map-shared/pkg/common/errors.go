package common

import (
	"net/http"

	"github.com/go-chi/render"
)

const ERR_VALIDATION_EMPTY_VALUE = "value is empty"
const ERR_NOT_FOUND = "not found"
const ERR_WRONG_TYPE = "bad type or struct provided"
const ERR_DB_NOT_FOUND = "database not found"

type ErrorPayload struct {
	HTTPStatus int    `json:"httpStatus"`
	Error      string `json:"error,omitempty"`
}

// ErrorResponse returns an error response
func ErrorResponse(w http.ResponseWriter, r *http.Request, httpStatusCode int, err error) {
	w.WriteHeader(httpStatusCode)
	errorPayload := &ErrorPayload{
		HTTPStatus: httpStatusCode,
		Error:      err.Error(),
	}
	render.JSON(w, r, errorPayload)
}
