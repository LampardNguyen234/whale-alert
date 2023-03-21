package logger

import "testing"

func TestZeroLogLoggerToConsole(t *testing.T) {
	log := NewZeroLogger("", "hello")

	log.Infof("This is my message: %v\n", "Hello world")
}

func TestZeroLogger_WithPrefix(t *testing.T) {
	log := NewZeroLogger("")
	log.Infof("This is my message: %v", "Hello world")

	log.WithPrefix("hello").Infof("This is my message: %v", "Hello world")
}
