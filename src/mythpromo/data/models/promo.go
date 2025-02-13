package models

type Promo struct {
	DiscountPercentage int      `bson:"discountpercentage"`
	Categories         []string `bson:"categories"`
	Skus               []string `bson:"skus"`
}
