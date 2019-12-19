package rest

import (
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
	// validation
	// logging
	query := req.URL.Query()
	search := query.Get("query")
	lat, _ := strconv.ParseFloat(query.Get("lat"), 32)
	long, _ := strconv.ParseFloat(query.Get("long"), 32)

	result, err := s.searchShopService.FindShopsByQuery(req.Context(), search, float32(lat), float32(long))

	if err != nil {
		//
	}

	respondSuccess(w, http.StatusOK, result)
}
