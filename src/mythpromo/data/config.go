package data

import (
	"log"
)

const DefaultDriver = "mongodb"

type Config struct {
	Driver string
	Conn   string
	DbName string
}

var db MongoDb
var config *Config

func New(cfg *Config) (DataRepository, error) {
	log.Printf("initializing database %s", cfg.Driver)
	loadConfig(cfg)
	db = MongoDb{}
	err := db.Connect(config.Conn)
	if err != nil {
		return nil, err
	}
	repo := DefaultRepository{
		db: db,
	}
	return repo, nil
}

func loadConfig(cfg *Config) error {
	if cfg.Driver == "" {
		cfg.Driver = DefaultDriver
	}
	config = cfg
	return nil
}
