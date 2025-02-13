package main

import (
	"log"
	"os"

	"github.com/andres1gb/mythpromo/api"
	"github.com/andres1gb/mythpromo/data"
	"github.com/andres1gb/mythpromo/shop"
	"github.com/pelletier/go-toml/v2"
)

const ( // exit codes
	ok int = iota
	cfgFileError
	cfgFormatError
	initDbError
	initShopError
	initApiError
)

type Config struct { // config file struct
	Api  api.Config
	Shop shop.Config
	Data data.Config
}

var config Config

func main() {
	cfg, err := os.ReadFile("config.toml")
	if err != nil {
		log.Fatal(err)
		os.Exit(cfgFileError)
	}

	err = toml.Unmarshal([]byte(cfg), &config)
	if err != nil {
		log.Fatal(err)
		os.Exit(cfgFormatError)
	}

	repo, err := data.New(&config.Data)
	if err != nil {
		log.Fatal(err)
		os.Exit(initDbError)
	}
	s, err := shop.New(&config.Shop, repo)
	if err != nil {
		log.Fatal(err)
		os.Exit(initShopError)
	}
	err = api.Init(&config.Api, s)
	if err != nil {
		log.Fatal(err)
		os.Exit(initApiError)
	}

}
