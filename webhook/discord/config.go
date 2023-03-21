package discord

import (
	"fmt"
	"github.com/LampardNguyen234/whale-alert/common"
)

// DiscordConfig consists of configurations of the DiscordHook.
type DiscordConfig struct {
	// Token is the OAuth token for pushing messages.
	Token string `json:"Token"`

	// SubChannels is the list of channels for posting messages.
	//
	// Each channel must start with an `@` symbol. For example: @astra_alert
	SubChannels []string `json:"SubChannels"`

	// MessageQueueSize specifies the message queue size.
	MessageQueueSize int `json:"MessageQueueSize"`

	// EnabledMessageTypes specifies which message types are allowed to be posted.
	EnabledMessageTypes map[string]bool `json:"EnabledMessageTypes"`
}

func DefaultConfig() DiscordConfig {
	return DiscordConfig{
		Token:               "DISCORD_TOKEN",
		SubChannels:         []string{"CHANNEL_ID_0", "CHANNEL_ID_1"},
		MessageQueueSize:    1024,
		EnabledMessageTypes: map[string]bool{common.InfoType: true, common.AlertType: true},
	}
}

// IsValid checks if the current TeleConfig is valid.
func (cfg DiscordConfig) IsValid() (bool, error) {
	if cfg.Token == "" {
		return false, fmt.Errorf("empty bot token")
	}

	if len(cfg.SubChannels) == 0 {
		return false, fmt.Errorf("empty subscribing channel")
	}

	if cfg.MessageQueueSize == 0 {
		return false, fmt.Errorf("MessageQueueSize must be greater than 0")
	}

	return true, nil
}
