package data

import (
	"github.com/andres1gb/mythpromo/data/models"
)

type MockRepository struct {
}

func (r MockRepository) FindProducts(category string, maxPrice uint32) (products []models.Product, err error) {
	product := models.Product{
		Name:     "Test product",
		Sku:      "99999",
		Category: category,
		Price:    maxPrice,
	}
	products = append(products, product)
	return
}

func (r MockRepository) FindPromosAvailable(product models.Product) (promos []models.Promo, err error) {
	promo := models.Promo{
		DiscountPercentage: 10,
		Categories:         []string{product.Category},
		Skus:               []string{product.Sku},
	}
	promos = append(promos, promo)
	return
}
