package discord

import (
	"fmt"
	"github.com/LampardNguyen234/whale-alert/common"
	"github.com/LampardNguyen234/whale-alert/logger"
	"github.com/bwmarrin/discordgo"
	"sync"
	"time"
)

// DiscordHook implements a Discord webhook for pushing messages.
type DiscordHook struct {
	*discordgo.Session
	cfg *DiscordConfig

	updateChannel chan interface{}
	closeChannel  chan interface{}
	stopped       bool

	mtx *sync.Mutex
	log logger.Logger
}

// NewDiscordHook creates and returns a new DiscordHook with the given DiscordConfig.
func NewDiscordHook(cfg DiscordConfig, log logger.Logger) (*DiscordHook, error) {
	session, err := discordgo.New(fmt.Sprintf("Bot %v", cfg.Token))
	if err != nil {
		return nil, err
	}

	return &DiscordHook{
		Session:       session,
		cfg:           &cfg,
		updateChannel: make(chan interface{}),
		closeChannel:  make(chan interface{}),
		stopped:       false,
		mtx:           new(sync.Mutex),
		log:           log.WithPrefix("Discord"),
	}, nil
}

// Ping checks if the remote Tele service is alive.
func (h *DiscordHook) Ping() error {
	return nil
}

// Start starts the TeleHook.
func (h *DiscordHook) Start() error {
	go h.handleUpdateChan()

	return nil
}

// Stop stops the TeleHook.
func (h *DiscordHook) Stop() {
	h.closeChannel <- true

	for !h.stopped {
		time.Sleep(1 * time.Second)
	}

	// close channels
	close(h.updateChannel)
	close(h.closeChannel)

	h.log.Infof("TeleHook STOPPED!")
}

// Info pushes the given information message to the TeleHook.
func (h *DiscordHook) Info(msg interface{}) error {
	if !h.cfg.EnabledMessageTypes[common.InfoType] {
		return nil
	}
	return h.pushMessage(map[string]interface{}{"Type": common.InfoType, "Message": msg})
}

// Error pushes the given error message to the TeleHook.
func (h *DiscordHook) Error(msg interface{}) error {
	if !h.cfg.EnabledMessageTypes[common.ErrorType] {
		return nil
	}
	return h.pushMessage(map[string]interface{}{"Type": common.ErrorType, "Message": msg})
}

// Alert pushes the given alert message to the TeleHook.
func (h *DiscordHook) Alert(msg interface{}) error {
	if !h.cfg.EnabledMessageTypes[common.AlertType] {
		return nil
	}

	return h.pushMessage(msg)
}

func (h *DiscordHook) pushMessage(msg interface{}) error {
	if len(h.updateChannel) >= h.cfg.MessageQueueSize {
		h.log.Errorf("channel is full")
		return fmt.Errorf("channel is full")
	}

	h.updateChannel <- msg

	return nil
}

func (h *DiscordHook) handleUpdateChan() {
	stop := false
	for {
		select {
		case _ = <-h.closeChannel:
			h.mtx.Lock()
			h.stopped = true
			stop = true
			h.mtx.Unlock()
			break
		case msg := <-h.updateChannel:
			h.log.Infof("New message received: %v\n", common.MustFormatJson(msg))
			for _, subChannel := range h.cfg.SubChannels {
				msgStr, ok := msg.(string)
				if !ok {
					msgStr = common.MustFormatJson(msg)
				}
				if _, err := h.ChannelMessageSend(subChannel, msgStr); err != nil {
					h.log.Errorf("Error sending message to Discord.\nMessage: %v\nError: %v", msg, err)
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
