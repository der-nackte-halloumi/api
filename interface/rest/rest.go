package rest

import "net/http"

import "github.com/der-nackte-halloumi/api/usecase/search_shop"

import "github.com/gorilla/mux"

type RestAPI struct {
	server *http.Server
}

func NewRestAPI(
	searchShopService search_shop.Service,
	port string,
) (*RestAPI, error) {
	router := mux.NewRouter()

	router.HandleFunc("/shops", NewShopsHandler(searchShopService).Search).Methods(http.MethodGet)

	return nil, nil
}
