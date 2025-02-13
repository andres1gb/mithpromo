package shop

import "github.com/andres1gb/mythpromo/data"

const DefaultCurrency = "EUR"
const DefaultPromoStrategy = "BestPromo"

type Config struct {
	DefaultCurrency string
	PromoStrategy   string
}

var config *Config

func New(cfg *Config, repo data.DataRepository) (Shop, error) {
	loadConfig(cfg)
	config = cfg
	st, err := getStrategy(cfg.PromoStrategy)
	if err != nil {
		return nil, err
	}
	return DefaultShop{
		promoStrategy: st,
		mgr:           repo,
	}, nil
}

func loadConfig(cfg *Config) {
	if cfg.DefaultCurrency == "" {
		cfg.DefaultCurrency = DefaultCurrency
	}
	if cfg.PromoStrategy == "" {
		cfg.PromoStrategy = DefaultPromoStrategy
	}
}
