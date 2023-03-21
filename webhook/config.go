package webhook

import (
	"fmt"
	"github.com/LampardNguyen234/whale-alert/webhook/discord"
	"github.com/LampardNguyen234/whale-alert/webhook/telegram"
)

type WebHookConfig struct {
	Discord *discord.DiscordConfig `json:"Discord,omitempty"`

	Telegram *telegram.TeleConfig `json:"Telegram,omitempty"`
}

func DefaultConfig() WebHookConfig {
	tele := telegram.DefaultConfig()
	discrd := discord.DefaultConfig()
	return WebHookConfig{
		Telegram: &tele,
		Discord:  &discrd,
	}
}

func (cfg WebHookConfig) IsValid() (bool, error) {
	if cfg.Telegram != nil {
		if _, err := cfg.Telegram.IsValid(); err != nil {
			return false, fmt.Errorf("invalid Telegram config: %v", err)
		}
	}
	if cfg.Discord != nil {
		if _, err := cfg.Discord.IsValid(); err != nil {
			return false, fmt.Errorf("invalid Discord config: %v", err)
		}
	}
	if cfg.Discord == nil && cfg.Telegram == nil {
		return false, fmt.Errorf("empty config")
	}

	return true, nil
}
