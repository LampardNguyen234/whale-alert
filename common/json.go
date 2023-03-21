package common

import (
	"encoding/json"
	"log"
)

// FormatJson returns the json-encoded of the given msg.
func FormatJson(msg interface{}) (string, error) {
	jsb, err := json.Marshal(msg)
	if err != nil {
		return "", err
	}

	return string(jsb), nil
}

// MustFormatJson is the same as FormatJson, but it will panic if the given msg cannot be marshalled.
func MustFormatJson(msg interface{}) string {
	jsb, err := json.Marshal(msg)
	if err != nil {
		log.Panicf("[Error] cannot marshal %v\n", msg)
	}

	return string(jsb)
}
