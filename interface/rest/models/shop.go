package rest

import "github.com/der-nackte-halloumi/api/domain"

type Shop struct {
	ID      string  `json:"id"`
	Name    string  `json:"name"`
	Address string  `json:"address"`
	Lat     float32 `json:"lat"`
	Lng     float32 `json:"lng"`
}

func ShopFromDomain(shop domain.Shop) *Shop {
	return &Shop{
		ID:      shop.ID.String(),
		Name:    shop.Name,
		Address: shop.Address,
		Lat:     shop.Lat,
		Lng:     shop.Lng,
	}
}
