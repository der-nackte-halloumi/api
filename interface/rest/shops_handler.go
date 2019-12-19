package rest

import (
	"net/http"

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
}
