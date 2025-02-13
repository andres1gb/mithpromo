package data

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type MongoDb struct {
	Context context.Context
	Client  *mongo.Client
}

func (m *MongoDb) Connect(uri string) (err error) {
	var cancel context.CancelFunc
	m.Context, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	m.Client, err = mongo.Connect(options.Client().ApplyURI(uri))
	return err
}
