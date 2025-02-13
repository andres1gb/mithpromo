package data

import (
	"context"
	"fmt"
	"testing"

	"github.com/andres1gb/mythpromo/data/models"
)

func TestFindProductsAndPromos(t *testing.T) {
	// TODO: replace for a mock db driver
	repo, err := New(&Config{
		Driver: "mongodb",
		Conn:   "mongodb://localhost:27017/test",
		DbName: "test",
	})
	if err != nil {
		t.Fatalf("error creating repo: %s", err)
	}
	product := models.Product{
		Name:     "test product",
		Sku:      "99999",
		Category: "tests",
		Price:    80000,
	}
	db.Client.Database(config.DbName).Collection("products").Drop(context.TODO())
	_, err = db.Client.Database(config.DbName).Collection("products").InsertOne(context.TODO(), product)
	if err != nil {
		t.Fatalf("error inserting product: %s", err)
	}
	result, err := repo.FindProducts("tests", 0)
	if err != nil {
		t.Fatalf("error accesing product: %s", err)
	}
	if len(result) != 1 {
		t.Fatal("error reading product")
	}
	db.Client.Database(config.DbName).Collection("promos").Drop(context.TODO())
	promo := models.Promo{
		DiscountPercentage: 10,
		Categories:         []string{"tests"},
		Skus:               []string{},
	}
	_, err = db.Client.Database(config.DbName).Collection("promos").InsertOne(context.TODO(), promo)
	if err != nil {
		t.Fatalf("error inserting promo: %s", err)
	}
	promo = models.Promo{
		DiscountPercentage: 15,
		Categories:         []string{},
		Skus:               []string{"99999"},
	}
	_, err = db.Client.Database(config.DbName).Collection("promos").InsertOne(context.TODO(), promo)
	if err != nil {
		t.Fatalf("error inserting promo: %s", err)
	}
	promos, err := repo.FindPromosAvailable(product)
	if err != nil {
		t.Fatalf("error querying promo: %s", err)
	}
	if len(promos) != 2 {
		fmt.Println(promos)
		t.Fatal("promos not found")
	}
}
