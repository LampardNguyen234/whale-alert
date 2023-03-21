package cmd

import (
	"context"
	"fmt"
	"github.com/LampardNguyen234/whale-alert/config"
	"github.com/LampardNguyen234/whale-alert/internal"
	"github.com/urfave/cli/v2"
	"log"
	"path/filepath"
)

const (
	flagConfig = "config"
)

func Run(args []string) error {
	cliApp := &cli.App{
		Name:                 filepath.Base(args[0]),
		Usage:                "Astra Whale Alert",
		Version:              "v0.0.1",
		Copyright:            "(c) 2023 stellalab.com",
		EnableBashCompletion: true,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  flagConfig,
				Value: "./config.json",
				Usage: "The config file to load from",
			},
		},
		Action: func(ctx *cli.Context) error {
			if args := ctx.Args(); args.Len() > 0 {
				return fmt.Errorf("unexpected arguments: %q", args.Get(0))
			}

			// Prepare FileConfig
			configPath := ctx.String(flagConfig)
			cfg, err := config.LoadConfigFromFile(configPath)
			if err != nil {
				log.Printf("failed to load config from file %v: %v", configPath, err)
				tmpCfg := config.DefaultConfig()
				cfg = &tmpCfg
			}

			mainApp, err := internal.NewApp(cfg)
			if err != nil {
				return err
			}

			mainApp.Start(context.Background())

			return nil
		},
	}

	err := cliApp.Run(args)
	if err != nil {
		return err
	}

	return nil
}
