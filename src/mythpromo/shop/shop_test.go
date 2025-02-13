package shop

import (
	"testing"

	"github.com/andres1gb/mythpromo/data"
)

func TestFindProducts(t *testing.T) {
	config := &Config{
		DefaultCurrency: "EUR",
		PromoStrategy:   "BestPromo",
	}
	repo := data.MockRepository{}
	sh, err := New(config, repo)
	if err != nil {
		t.Fatalf("can't create shop: %s", err)
	}
	products, err := sh.FindProducts("cool_stuff", 90000)
	if err != nil {
		t.Fatalf("can't search products: %s", err)
	}
	if len(products) < 1 {
		t.Fatal("product not found")
	}
}
