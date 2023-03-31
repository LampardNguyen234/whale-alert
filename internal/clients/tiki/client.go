package tiki

import (
	"github.com/LampardNguyen234/whale-alert/logger"
	"net/http"
)

type TikiClient struct {
	*http.Client
	cfg TikiClientConfig
	log logger.Logger
}

func NewTikiClient(cfg TikiClientConfig, log logger.Logger) *TikiClient {
	return &TikiClient{
		Client: &http.Client{},
		cfg:    cfg,
		log:    log.WithPrefix("Tiki-Client"),
	}
}
