package rest

import (
	"encoding/json"
	"net/http"
	"time"

	rest "github.com/der-nackte-halloumi/api/interface/rest/models"
	"github.com/der-nackte-halloumi/api/usecase/search_shop"
	"github.com/gorilla/mux"
)

type RestAPI struct {
	server *http.Server
}

func NewRestAPI(
	searchShopService search_shop.Service,
	port string,
) (*RestAPI, error) {
	router := mux.NewRouter()

	router.HandleFunc("/_health", NewStatusHandler().Health).Methods(http.MethodGet)
	router.HandleFunc("/shops", NewShopsHandler(searchShopService).Search).Methods(http.MethodGet)
	router.NotFoundHandler = http.HandlerFunc(NewErrorHandler().NotFound)

	server := &http.Server{
		Addr:         "0.0.0.0:" + port,
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      router,
	}

	return &RestAPI{server: server}, nil
}

func (r *RestAPI) Start() error {
	return r.server.ListenAndServe()
}

func respond(w http.ResponseWriter, code int, body interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(body)
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(rest.Error{
		Code:    code,
		Message: message,
	})
}
