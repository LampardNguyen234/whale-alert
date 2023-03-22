package telegram

import (
	"fmt"
	"github.com/LampardNguyen234/whale-alert/common"
	"github.com/LampardNguyen234/whale-alert/logger"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"sync"
	"time"
)

// TeleHook implements a Telegram webhook for pushing messages.
type TeleHook struct {
	*tgbotapi.BotAPI

	cfg *TeleConfig

	updateChannel chan interface{}
	closeChannel  chan interface{}
	stopped       bool

	mtx *sync.Mutex
	log logger.Logger
}

// NewTeleHook creates a new TeleHook with the given TeleConfig.
func NewTeleHook(cfg TeleConfig, log logger.Logger) (*TeleHook, error) {
	if _, err := cfg.IsValid(); err != nil {
		return nil, fmt.Errorf("invalid TeleConfig %v: %v", cfg, err)
	}
	bot, err := tgbotapi.NewBotAPI(cfg.Token)
	if err != nil {
		return nil, err
	}

	tmpLogger := log.WithPrefix("Telegram")
	tmpLogger.Infof("Authorized on account %s (https://t.me/%s)", bot.Self.UserName, bot.Self.UserName)

	return &TeleHook{
		BotAPI:        bot,
		cfg:           &cfg,
		updateChannel: make(chan interface{}, cfg.MessageQueueSize),
		closeChannel:  make(chan interface{}, 1),
		mtx:           new(sync.Mutex),
		log:           tmpLogger,
	}, nil
}

// Ping checks if the remote Tele service is alive.
func (t *TeleHook) Ping() error {
	return nil
}

// Start starts the TeleHook.
func (t *TeleHook) Start() error {
	go t.handleUpdateChan()

	return nil
}

// Stop stops the TeleHook.
func (t *TeleHook) Stop() {
	t.closeChannel <- true

	for !t.stopped {
		time.Sleep(1 * time.Second)
	}

	// close channels
	close(t.updateChannel)
	close(t.closeChannel)

	t.log.Infof("TeleHook STOPPED!")
}

// Info pushes the given information message to the TeleHook.
func (t *TeleHook) Info(msg interface{}) error {
	if !t.cfg.EnabledMessageTypes[common.InfoType] {
		return nil
	}
	return t.pushMessage(map[string]interface{}{"Type": common.InfoType, "Message": msg})
}

// Error pushes the given error message to the TeleHook.
func (t *TeleHook) Error(msg interface{}) error {
	if !t.cfg.EnabledMessageTypes[common.ErrorType] {
		return nil
	}
	return t.pushMessage(map[string]interface{}{"Type": common.ErrorType, "Message": msg})
}

// Alert pushes the given alert message to the TeleHook.
func (t *TeleHook) Alert(msg interface{}) error {
	if !t.cfg.EnabledMessageTypes[common.AlertType] {
		return nil
	}

	return t.pushMessage(msg)
}

func (t *TeleHook) pushMessage(msg interface{}) error {
	if len(t.updateChannel) >= t.cfg.MessageQueueSize {
		t.log.Errorf("channel is full")
		return fmt.Errorf("channel is full")
	}

	t.updateChannel <- msg

	return nil
}

func (t *TeleHook) handleUpdateChan() {
	stop := false
	for {
		select {
		case _ = <-t.closeChannel:
			t.mtx.Lock()
			t.stopped = true
			stop = true
			t.mtx.Unlock()
			break
		case msg := <-t.updateChannel:
			t.log.Debugf("New message received: %v", common.MustFormatJson(msg))
			for _, subChannel := range t.cfg.SubChannels {
				var chattableMsg tgbotapi.Chattable
				msgStr, ok := msg.(string)
				if !ok {
					chattableMsg = tgbotapi.NewMessageToChannel(subChannel, common.MustFormatJson(msg))
				} else {
					chattableMsg = tgbotapi.NewMessageToChannel(subChannel, msgStr)
				}

				if _, err := t.Send(chattableMsg); err != nil {
					t.log.Errorf("Error sending message to telegram.\nMessage: %v\nError: %v", msg, err)
				}
			}
		default:
			time.Sleep(1 * time.Second)
		}

		if stop {
			break
		}
	}
}
