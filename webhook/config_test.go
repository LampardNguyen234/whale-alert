package webhook

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestDefaultConfig(t *testing.T) {
	cfg := DefaultConfig()
	jsb, err := json.MarshalIndent(cfg, "", "\t")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(jsb))
}
