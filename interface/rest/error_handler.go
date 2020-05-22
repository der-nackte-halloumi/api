package rest

import (
	"net/http"
)

type errorHandler struct{}

func NewErrorHandler() *errorHandler {
	return &errorHandler{}
}

func (n *errorHandler) NotFound(w http.ResponseWriter, req *http.Request) {
	respondWithError(w, http.StatusNotFound, "the requested path '"+req.RequestURI+"' was not found")
}
