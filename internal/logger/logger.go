package logger

// https://www.mountedthoughts.com/golang-logger-interface/

// A global variable so that log functions can be directly accessed
var log Logger

// Fields Type to pass when we want to call WithFields for structured logging
type Fields map[string]interface{}

const (
	// Debug has verbose message
	Debug = "debug"
	// Info is default log level
	Info = "info"
	// Warn is for logging messages about possible issues
	Warn = "warn"
	// Error is for logging errors
	Error = "error"
	// Fatal is for logging fatal messages. The system shuts down after logging the message.
	Fatal = "fatal"
)

type Logger interface {
	Debugf(format string, args ...interface{})

	Infof(format string, args ...interface{})

	Warnf(format string, args ...interface{})

	Errorf(format string, args ...interface{})

	Fatalf(format string, args ...interface{})

	Panicf(format string, args ...interface{})

	WithFields(keyValues Fields) Logger
}

// Configuration stores the config for the logger
type Configuration struct {
	Level string
}

// NewLogger returns an instance of logger
func NewLogger(config Configuration) (Logger, error) {
	logger, err := NewLogrusLogger(config)
	if err != nil {
		return nil, err
	}
	return logger, nil
}

// The default configuration
var defaultConfig = Configuration{
	Level: Info,
}

func init() {
	l, err := NewLogger(defaultConfig)

	if err != nil {
		panic(err)
	}

	log = l
}

func Debugf(format string, args ...interface{}) {
	log.Debugf(format, args...)
}

func Infof(format string, args ...interface{}) {
	log.Infof(format, args...)
}

func Warnf(format string, args ...interface{}) {
	log.Warnf(format, args...)
}

func Errorf(format string, args ...interface{}) {
	log.Errorf(format, args...)
}

func Fatalf(format string, args ...interface{}) {
	log.Fatalf(format, args...)
}

func Panicf(format string, args ...interface{}) {
	log.Panicf(format, args...)
}

func WithFields(keyValues Fields) Logger {
	return log.WithFields(keyValues)
}
