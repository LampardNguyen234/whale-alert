package webhook

import (
	"fmt"
	"github.com/LampardNguyen234/whale-alert/logger"
	"github.com/LampardNguyen234/whale-alert/webhook/discord"
	"github.com/LampardNguyen234/whale-alert/webhook/telegram"
)

// WebHookManager manages the list of WebHooks and is the main entrance to posting messages.
type WebHookManager struct {
	webHooks map[string]WebHook
}

// NewWebHookManager returns a new WebHookManager.
func NewWebHookManager(whs map[string]WebHook) *WebHookManager {
	if whs == nil {
		whs = make(map[string]WebHook)
	}
	return &WebHookManager{webHooks: whs}
}

// NewWebHookManagerFromConfig returns a new WebHookManager from the given config.
func NewWebHookManagerFromConfig(cfg WebHookConfig, log logger.Logger) (*WebHookManager, error) {
	whs := make(map[string]WebHook)
	var err error

	if cfg.Telegram != nil {
		whs["Telegram"], err = telegram.NewTeleHook(*cfg.Telegram, log)
		if err != nil {
			fmt.Println("tele", err)
			return nil, err
		}
	}
	if cfg.Discord != nil {
		whs["Discord"], err = discord.NewDiscordHook(*cfg.Discord, log)
		if err != nil {
			fmt.Println("discord", err)
			return nil, err
		}
	}

	return &WebHookManager{webHooks: whs}, nil
}

// AddWebHook adds a new WebHook to the WebHookManager.
func (m *WebHookManager) AddWebHook(name string, wh WebHook) {
	m.webHooks[name] = wh
}

// RemoveWebHook removes a WebHook given its name.
func (m *WebHookManager) RemoveWebHook(name string) {
	if m.webHooks[name] != nil {
		delete(m.webHooks, name)
	}
}

// Start starts all the WebHook of the WebHookManager.
func (m *WebHookManager) Start() error {
	var err error
	for _, wh := range m.webHooks {
		err = wh.Start()
		if err != nil {
			return err
		}
	}

	return nil
}

// Stop terminates all the processes of the WebHookManager.
func (m *WebHookManager) Stop() {
	for _, wh := range m.webHooks {
		go wh.Stop()
	}
}

// Info pushes the given information message to the WebHook's.
func (m *WebHookManager) Info(msg interface{}) error {
	for _, wh := range m.webHooks {
		_ = wh.Info(msg)
	}

	return nil
}

// Error pushes the given error message to the WebHook's.
func (m *WebHookManager) Error(msg interface{}) error {
	for _, wh := range m.webHooks {
		_ = wh.Error(msg)
	}

	return nil
}

// Alert pushes the given alert message to the WebHook's.
func (m *WebHookManager) Alert(msg interface{}) error {
	for _, wh := range m.webHooks {
		_ = wh.Alert(msg)
	}

	return nil
}
