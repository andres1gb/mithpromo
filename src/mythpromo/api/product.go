package api

type ProductsRequest struct {
	Category      string `form:"category"`
	PriceLessThan uint32 `form:"priceLessThan"`
}

type ProductResponse struct {
	Sku      string `json:"sku"`
	Name     string `json:"name"`
	Category string `json:"category"`
	Price    Price  `json:"price"`
}

type Price struct {
	Original           uint32 `json:"original"`
	Final              uint32 `json:"final"`
	DiscountPercentage string `json:"discount_percentage"`
	Currency           string `json:"currency"`
}
