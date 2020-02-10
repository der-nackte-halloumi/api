package rest

import (
	"log"
	"net/http"
	"strconv"

	rest "github.com/der-nackte-halloumi/api/interface/rest/models"
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
		// TODO: make error object
		respond(w, http.StatusBadRequest, "must provide a search query")
		return
	}

	result, err := s.searchShopService.FindShopsByQuery(req.Context(), search, float32(lat), float32(long))
	// TODO: log/track searches and len(result)

	if err != nil {
		log.Printf("error when searching shops: %v", err)
		// TODO: return proper status code
		respond(w, http.StatusInternalServerError, err)
		return
	}

	var shops []*rest.Shop
	for _, shop := range result {
		shops = append(shops, rest.ShopFromDomain(shop))
	}

	respond(w, http.StatusOK, shops)
}
