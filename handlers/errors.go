package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

var (
	ErrBadJsonError = BadRequestError{Message: "Bad json format"}
)

func HandleError(w http.ResponseWriter, err error) {
	log.Print(err)
	e := createErrorInternal(err)
	e.render(w)
}

func createErrorInternal(err error) *apiError {
	e := &apiError{}
	switch v := err.(type) {
	case BadRequestError:
		e.Message = v.Message
		e.StatusCode = http.StatusBadRequest
	default:
		e.Message = "internal error"
		e.StatusCode = http.StatusInternalServerError
	}
	return e
}

func (v *apiError) render(w http.ResponseWriter) {
	b, _ := json.Marshal(v)
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(v.StatusCode)
	_, _ = w.Write(b)
}

type BadRequestError struct {
	Message string
}

func (e BadRequestError) Error() string {
	return e.Message
}
