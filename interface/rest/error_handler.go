package rest

import (
	"net/http"

	rest "github.com/der-nackte-halloumi/api/interface/rest/models"
)

type errorHandler struct{}

func NewErrorHandler() *errorHandler {
	return &errorHandler{}
}

func (n *errorHandler) NotFound(w http.ResponseWriter, req *http.Request) {
	respond(w, http.StatusNotFound, rest.Error{
		Code:    http.StatusNotFound,
		Message: "the requested path '" + req.RequestURI + "' was not found",
	})
}
