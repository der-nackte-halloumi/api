package rest

import (
	"net/http"

	rest "github.com/der-nackte-halloumi/api/interface/rest/models"
)

type statusHandler struct{}

func NewStatusHandler() *statusHandler {
	return &statusHandler{}
}

func (s *statusHandler) Health(w http.ResponseWriter, req *http.Request) {
	respond(w, http.StatusOK, rest.HealthCheck{Alive: true})
}
