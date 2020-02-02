package rest

import (
	"log"
	"net/http"
	"strconv"

	"github.com/der-nackte-halloumi/api/usecase/search_shop"
)

type shopsHandler struct {
	searchShopService search_shop.Service
}

func NewShopsHandler(searchShopService search_shop.Service) *shopsHandler {
	return &shopsHandler{searchShopService: searchShopService}
}

func (s *shopsHandler) Search(w http.ResponseWriter, req *http.Request) {
	// TODO: logging
	query := req.URL.Query()

	// TODO: validation all inputs
	search := query.Get("query")
	lat, _ := strconv.ParseFloat(query.Get("lat"), 32)
	long, _ := strconv.ParseFloat(query.Get("long"), 32)

	if len(search) == 0 {
		respond(w, http.StatusOK, nil)
		return
	}

	result, err := s.searchShopService.FindShopsByQuery(req.Context(), search, float32(lat), float32(long))
	// TODO: log/track searches and len(result)

	if err != nil {
		log.Printf("error when searching shops: %v", err)
		respond(w, http.StatusInternalServerError, err)
	}

	respond(w, http.StatusOK, result)
}
