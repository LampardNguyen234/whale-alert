package logger

import (
	"fmt"
	"github.com/rs/zerolog"
	"log"
	"os"
	"strings"
	"time"
)

// ZeroLogger implements Logger interface and provide logging service
// using zerolog.
type ZeroLogger struct {
	instance *zerolog.Logger
	prefix   string
	writer   zerolog.ConsoleWriter
}

func NewZeroLogger(logFile string, prefix ...string) *ZeroLogger {
	tmpPrefix := ""
	if len(prefix) > 0 {
		tmpPrefix = prefix[0]
	}

	output := os.Stdout
	var err error
	if logFile != "" {
		if _, err := os.Stat(logFile); os.IsNotExist(err) {
			tmpStrings := strings.Split(logFile, "/")
			if len(tmpStrings) > 1 {
				directory := strings.Replace(logFile, tmpStrings[len(tmpStrings)-1], "", -1)
				err = os.MkdirAll(directory, os.ModePerm)
				if err != nil {
					fmt.Printf("make directory %v error: %v\n", directory, err)
					os.Exit(1)
				}
			}
		}

		output, err = os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			log.Println("Error opening file:", err)
			os.Exit(1)
		}
	}

	writer := zerolog.ConsoleWriter{
		Out:        output,
		NoColor:    true,
		TimeFormat: time.RFC3339,
		FormatLevel: func(i interface{}) string {
			if tmpPrefix != "" {
				return strings.ToUpper(fmt.Sprintf("[%v-%v]", i, tmpPrefix))
			}
			return strings.ToUpper(fmt.Sprintf("[%s]", i))
		},
		FormatMessage: func(i interface{}) string {
			return fmt.Sprintf("| %s", i)
		},
	}
	logger := zerolog.New(writer)

	return &ZeroLogger{
		instance: &logger,
		prefix:   tmpPrefix,
		writer:   writer,
	}
}

func NewZeroLoggerWithColor(logFile string, prefix ...string) *ZeroLogger {
	tmpPrefix := ""
	if len(prefix) > 0 {
		tmpPrefix = prefix[0]
	}

	output := os.Stdout
	var err error
	if logFile != "" {
		if _, err := os.Stat(logFile); os.IsNotExist(err) {
			tmpStrings := strings.Split(logFile, "/")
			if len(tmpStrings) > 1 {
				directory := strings.Replace(logFile, tmpStrings[len(tmpStrings)-1], "", -1)
				err = os.MkdirAll(directory, os.ModePerm)
				if err != nil {
					fmt.Printf("make directory %v error: %v\n", directory, err)
					os.Exit(1)
				}
			}
		}

		output, err = os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			log.Println("Error opening file:", err)
			os.Exit(1)
		}
	}

	writer := zerolog.ConsoleWriter{
		Out:        output,
		NoColor:    false,
		TimeFormat: time.RFC3339,
		FormatLevel: func(i interface{}) string {
			if tmpPrefix != "" {
				return strings.ToUpper(fmt.Sprintf("[%v-%v]", i, tmpPrefix))
			}
			return strings.ToUpper(fmt.Sprintf("[%s]", i))
		},
		FormatMessage: func(i interface{}) string {
			return fmt.Sprintf("| %s", i)
		},
	}
	logger := zerolog.New(writer)

	return &ZeroLogger{
		instance: &logger,
		prefix:   tmpPrefix,
		writer:   writer,
	}
}

func (logger *ZeroLogger) SetLogLevel(level LogLevel) {
	updatedLogger := logger.instance.Level(logLevelToZerologLevel(level))

	logger.instance = &updatedLogger
}

func logLevelToZerologLevel(logLevel LogLevel) zerolog.Level {
	switch logLevel {
	case LogDisabled:
		return zerolog.Disabled
	case LogLevelPanic:
		return zerolog.PanicLevel
	case LogLevelError:
		return zerolog.ErrorLevel
	case LogLevelInfo:
		return zerolog.InfoLevel
	case LogLevelDebug:
		return zerolog.DebugLevel
	default:
		panic(fmt.Sprintf("Unsupported log level %v", logLevel))
	}
}

func (logger *ZeroLogger) GetLogLevel() LogLevel {
	return zeroLogLevelToLogLevel(logger.instance.GetLevel())
}

func zeroLogLevelToLogLevel(logLevel zerolog.Level) LogLevel {
	switch logLevel {
	case zerolog.Disabled:
		return LogDisabled
	case zerolog.PanicLevel, zerolog.FatalLevel:
		return LogLevelPanic
	case zerolog.ErrorLevel:
		return LogLevelError
	case zerolog.WarnLevel, zerolog.InfoLevel:
		return LogLevelInfo
	case zerolog.DebugLevel, zerolog.TraceLevel:
		return LogLevelDebug
	default:
		panic(fmt.Sprintf("Unsupported log level %v", logLevel))
	}

	return LogLevelDebug
}

func (logger *ZeroLogger) Panic(message string) {
	logger.instance.Panic().Timestamp().Msg(message)
}

func (logger *ZeroLogger) Panicf(format string, values ...interface{}) {
	logger.instance.Panic().Timestamp().Msgf(format, values...)
}

func (logger *ZeroLogger) Error(message string) {
	logger.instance.Error().Timestamp().Msg(message)
}

func (logger *ZeroLogger) Errorf(format string, values ...interface{}) {
	logger.instance.Error().Timestamp().Msgf(format, values...)
}

func (logger *ZeroLogger) Info(message string) {
	logger.instance.Info().Timestamp().Msg(message)
}

func (logger *ZeroLogger) Infof(format string, values ...interface{}) {
	logger.instance.Info().Timestamp().Msgf(format, values...)
}

func (logger *ZeroLogger) Debug(message string) {
	logger.instance.Debug().Timestamp().Msg(message)
}

func (logger *ZeroLogger) Debugf(format string, values ...interface{}) {
	logger.instance.Debug().Timestamp().Msgf(format, values...)
}

func (logger *ZeroLogger) WithInterface(key string, value interface{}) Logger {
	instance := logger.instance.With().Interface(key, value).Logger()
	return &ZeroLogger{
		instance: &instance,
		prefix:   logger.prefix,
		writer:   logger.writer,
	}
}

func (logger *ZeroLogger) WithFields(fields LogFields) Logger {
	instance := logger.instance.With().Fields(fields).Logger()
	return &ZeroLogger{
		instance: &instance,
		prefix:   logger.prefix,
		writer:   logger.writer,
	}
}

func (logger *ZeroLogger) WithPrefix(prefix string) Logger {
	writer := logger.writer
	writer.FormatLevel = func(i interface{}) string {
		if prefix != "" {
			return strings.ToUpper(fmt.Sprintf("[%v-%v]", i, prefix))
		}
		return strings.ToUpper(fmt.Sprintf("[%s]", i))
	}

	instance := zerolog.New(writer)
	return &ZeroLogger{
		instance: &instance,
		prefix:   prefix,
		writer:   writer,
	}
}
