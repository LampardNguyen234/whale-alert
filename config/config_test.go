package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"testing"
)

func TestDefaultConfig(t *testing.T) {
	cfg := DefaultConfig()

	jsb, _ := json.MarshalIndent(cfg, "", "\t")

	err := ioutil.WriteFile("../example_config.json", jsb, os.ModePerm)
	if err != nil {
		panic(err)
	}
}
