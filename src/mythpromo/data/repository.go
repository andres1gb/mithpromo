package data

import (
	"context"

	"github.com/andres1gb/mythpromo/data/models"
	"go.mongodb.org/mongo-driver/bson"
)

type DataRepository interface {
	FindProducts(category string, maxPrice uint32) (products []models.Product, err error)
	FindPromosAvailable(product models.Product) (promos []models.Promo, err error)
}

type DefaultRepository struct {
	db MongoDb
}

func (r DefaultRepository) FindProducts(category string, maxPrice uint32) (products []models.Product, err error) {
	var filter bson.M

	switch true {
	case (category != "" && maxPrice > 0):
		filter = bson.M{
			"category": category,
			"price":    bson.M{"$lte": maxPrice},
		}
	case (category != "" && maxPrice == 0):
		filter = bson.M{"category": category}
	case (category == "" && maxPrice > 0):
		filter = bson.M{"price": bson.M{"$lte": maxPrice}}
	default:
		filter = bson.M{}
	}

	result := make([]bson.M, 0)
	cursor, err := r.db.Client.Database(config.DbName).Collection("products").Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	err = cursor.All(context.TODO(), &result)
	if err != nil {
		return nil, err
	}
	for _, item := range result {
		var product models.Product
		bytes, err := bson.Marshal(item)
		if err != nil {
			return nil, err
		}
		err = bson.Unmarshal(bytes, &product)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return
}

func (r DefaultRepository) FindPromosAvailable(product models.Product) (promos []models.Promo, err error) {
	filter := bson.M{
		"$or": []bson.M{
			{"categories": product.Category},
			{"skus": product.Sku},
		},
	}
	result := make([]bson.M, 0)
	cursor, err := r.db.Client.Database(config.DbName).Collection("promos").Find(context.TODO(), filter)

	if err != nil {
		return nil, err
	}
	err = cursor.All(context.TODO(), &result)
	if err != nil {
		return nil, err
	}
	for _, item := range result {
		var promo models.Promo
		bytes, err := bson.Marshal(item)
		if err != nil {
			return nil, err
		}
		err = bson.Unmarshal(bytes, &promo)
		if err != nil {
			return nil, err
		}
		promos = append(promos, promo)
	}
	return
}
