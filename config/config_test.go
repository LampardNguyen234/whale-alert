package config

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestDefaultConfig(t *testing.T) {
	cfg := DefaultConfig()
	jsb, _ := json.MarshalIndent(cfg, "", "\t")
	fmt.Println(string(jsb))
}
