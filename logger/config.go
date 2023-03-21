package logger

type LoggerConfig struct {
	LogPath string `json:"LogPath"`
	Level   int8   `json:"Level"`
	Color   bool   `json:"Color"`
}

func DefaultConfig() LoggerConfig {
	return LoggerConfig{
		LogPath: "",
		Level:   int8(LogLevelInfo),
		Color:   false,
	}
}

func (cfg LoggerConfig) IsValid() (bool, error) {
	return true, nil
}
