package discord

import (
	"fmt"
	"github.com/LampardNguyen234/whale-alert/common"
	"testing"
	"time"
)

var hook *DiscordHook

const numTests = 10

func init() {
	cfg := &DiscordConfig{
		Token:               "",
		SubChannels:         []string{""},
		MessageQueueSize:    1024,
		EnabledMessageTypes: map[string]bool{common.InfoType: true, common.AlertType: true},
	}

	var err error
	hook, err = NewDiscordHook(*cfg)
	if err != nil {
		panic(err)
	}

	err = hook.Start()
	if err != nil {
		panic(err)
	}
}

func sleepAndExit(duration time.Duration) {
	time.Sleep(duration)
	hook.Stop()
}

func TestDiscordHook_Info(t *testing.T) {
	for i := 0; i < numTests; i++ {
		err := hook.Info(fmt.Sprintf("Hello World"))
		if err != nil {
			panic(err)
		}
		time.Sleep(200 * time.Millisecond)
	}

	sleepAndExit(5 * time.Second)
}

func TestDiscordHook_Alert(t *testing.T) {
	for i := 0; i < numTests; i++ {
		err := hook.Alert(fmt.Sprintf("Hello World"))
		if err != nil {
			panic(err)
		}
		time.Sleep(200 * time.Millisecond)
	}

	sleepAndExit(5 * time.Second)
}

func TestDiscordHook_Error(t *testing.T) {
	for i := 0; i < numTests; i++ {
		err := hook.Error(fmt.Sprintf("Hello World"))
		if err != nil {
			panic(err)
		}
		time.Sleep(200 * time.Millisecond)
	}

	sleepAndExit(5 * time.Second)
}
