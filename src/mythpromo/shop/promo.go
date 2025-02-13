package shop

import (
	"errors"

	"github.com/andres1gb/mythpromo/data/models"
)

type PromoStrategy interface {
	apply(price *Price, promos []models.Promo)
}

func getStrategy(name string) (PromoStrategy, error) {
	switch name {
	case "BestPromo":
		return &BestPromoStrategy{}, nil
	default:
		return nil, errors.New("unknown promo strategy")
	}
}

type BestPromoStrategy struct{}

func (ps *BestPromoStrategy) apply(price *Price, promos []models.Promo) {
	for _, promo := range promos {
		if price.DiscountPercentage < promo.DiscountPercentage {
			price.DiscountPercentage = promo.DiscountPercentage
		}
	}
	discount := price.Original * uint32(price.DiscountPercentage) / 100
	price.Final = price.Final - discount
}
