package api

import (
	"fmt"

	"github.com/andres1gb/mythpromo/shop"
	"github.com/gin-gonic/gin"
)

const (
	DefaultAddress = "0.0.0.0"
	DefaultPort    = 8080
)

type Config struct {
	Ip   string
	Port uint32
}

var config *Config
var manager shop.Shop

func Init(cfg *Config, mgr shop.Shop) error {
	manager = mgr
	loadConfig(cfg)
	r := gin.Default()
	loadRoutes(r)
	r.Run(fmt.Sprintf("%s:%d", config.Ip, config.Port))
	return nil
}

func loadConfig(cfg *Config) error {
	if cfg.Ip == "" {
		cfg.Ip = DefaultAddress
	}
	if cfg.Port == 0 {
		cfg.Port = DefaultPort
	}
	config = cfg
	return nil
}
