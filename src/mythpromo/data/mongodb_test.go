package data

import (
	"testing"
)

func TestConnect(t *testing.T) {
	db := MongoDb{}
	err := db.Connect("mongodb://localhost/test")
	if err != nil {
		t.Fatalf("failed to connect to database: %s", err.Error())
	}
	if db.Client == nil {
		t.Fatalf("failed to create mongodb client")
	}
}
