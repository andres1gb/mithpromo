package shop

type Price struct {
	Original           uint32
	Final              uint32
	DiscountPercentage int
	Currency           string
}

type Product struct {
	Sku      string
	Category string
	Name     string
	Price    Price
}
