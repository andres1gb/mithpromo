package models

type Product struct {
	Sku      string `bson:"sku"`
	Name     string `bson:"name"`
	Category string `bson:"category"`
	Price    uint32 `bson:"price"`
}
