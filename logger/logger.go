package logger

// LogLevel defines log priorities.
type LogLevel int8

// LogFields holds the detail to be logged.
type LogFields map[string]interface{}

const (
	LogLevelDebug LogLevel = iota
	LogLevelInfo
	LogLevelError
	LogLevelPanic
	LogDisabled
)

type Logger interface {
	SetLogLevel(level LogLevel)
	GetLogLevel() LogLevel

	Panic(message string)
	Panicf(format string, values ...interface{})
	Error(message string)
	Errorf(format string, values ...interface{})
	Info(message string)
	Infof(format string, values ...interface{})
	Debug(message string)
	Debugf(format string, values ...interface{})

	WithInterface(key string, value interface{}) Logger
	WithFields(fields LogFields) Logger
	WithPrefix(prefix string) Logger
}
