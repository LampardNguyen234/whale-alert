package tiki

import (
	"context"
	"github.com/LampardNguyen234/whale-alert/logger"
)

var (
	c   *TikiClient
	ctx = context.Background()
)

func init() {
	c = NewTikiClient(DefaultConfig(), logger.NewZeroLogger(""))
}
