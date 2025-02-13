package shop

import (
	"github.com/andres1gb/mythpromo/data"
)

const DefaultMaxResults = 5

type ProductSearchCriteria struct {
	Category      string
	PriceLessThan uint32
	MaxResults    uint32
}

type Shop interface {
	FindProducts(category string, priceLessThan uint32) ([]Product, error)
}

type DefaultShop struct {
	promoStrategy PromoStrategy
	mgr           data.DataRepository
}

func (s DefaultShop) FindProducts(category string, priceLessThan uint32) ([]Product, error) {
	result := make([]Product, 0)

	products, err := s.mgr.FindProducts(category, priceLessThan)
	if err != nil {
		return nil, err
	}
	// limit the number of results, it could be done at data layer
	if len(products) > DefaultMaxResults {
		products = products[0:DefaultMaxResults]
	}

	for _, product := range products {
		promos, err := s.mgr.FindPromosAvailable(product)
		if err != nil {
			return nil, err
		}
		productFound := Product{
			Sku:      product.Sku,
			Category: product.Category,
			Name:     product.Name,
			Price: Price{
				Original:           product.Price,
				Final:              product.Price,
				Currency:           config.DefaultCurrency,
				DiscountPercentage: 0,
			},
		}
		s.promoStrategy.apply(&productFound.Price, promos)
		result = append(result, productFound)
	}
	return result, nil
}
