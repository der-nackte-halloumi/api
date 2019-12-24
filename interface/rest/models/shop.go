package rest

type Shop struct {
	ID   string  `json:"id"`
	Name string  `json:"name"`
	lat  float32 `json:"lat"`
	lng  float32 `json:"long"`
}
