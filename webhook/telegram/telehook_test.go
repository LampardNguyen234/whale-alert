package telegram

import (
	"fmt"
	"github.com/LampardNguyen234/whale-alert/logger"
	"testing"
	"time"
)

var teleHook *TeleHook

const numTests = 10

func init() {
	cfg := DefaultConfig()
	var err error
	teleHook, err = NewTeleHook(cfg, logger.NewZeroLogger(""))
	if err != nil {
		panic(err)
	}

	err = teleHook.Start()
	if err != nil {
		panic(err)
	}
}

func sleepAndExit(duration time.Duration) {
	time.Sleep(duration)
	teleHook.Stop()
}

func TestTeleHook_Info(t *testing.T) {
	for i := 0; i < numTests; i++ {
		err := teleHook.Info(fmt.Sprintf("Hello World"))
		if err != nil {
			panic(err)
		}
		time.Sleep(200 * time.Millisecond)
	}

	sleepAndExit(5 * time.Second)
}

func TestTeleHook_Alert(t *testing.T) {
	for i := 0; i < numTests; i++ {
		err := teleHook.Alert(fmt.Sprintf("Hello World"))
		if err != nil {
			panic(err)
		}
		time.Sleep(200 * time.Millisecond)
	}

	sleepAndExit(5 * time.Second)
}

func TestTeleHook_Error(t *testing.T) {
	for i := 0; i < numTests; i++ {
		err := teleHook.Error(fmt.Sprintf("Hello World"))
		if err != nil {
			panic(err)
		}
		time.Sleep(200 * time.Millisecond)
	}

	sleepAndExit(5 * time.Second)
}
