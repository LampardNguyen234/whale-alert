package main

import (
	"fmt"
	"github.com/LampardNguyen234/whale-alert/cmd"
	_ "github.com/LampardNguyen234/whale-alert/docs"
	"os"
)

// @title           Whale Alert Admin APIs
// @version         1.0
// @description     This service is dedicated to administration of a Whale Alert service.

// @contact.name   StellaLab
// @contact.url    https://www.stellalab.com
// @contact.email  support@stellalab.com
// @securityDefinitions.apiKey	ApiKeyAuth
// @in							header
// @name						Authorization
// @description				    Use for admin authorization
func main() {
	if err := cmd.Run(os.Args); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
