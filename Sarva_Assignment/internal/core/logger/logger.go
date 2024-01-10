package logger

import (
	"github.com/hashicorp/go-hclog"
)

// Logger represents the core logic for logging.
type Logger struct {
	logger hclog.Logger
}

// NewLogger creates a new instance of Logger.
func NewLogger(logger hclog.Logger) *Logger {
	return &Logger{logger: logger}
}

// Log logs a message.
func (l *Logger) Log(message string) {
	l.logger.Info(message)
}

// Trace logs a trace message.
func (l *Logger) Trace(message string) {
	l.logger.Trace(message)
}

// Debug logs a debug message.
func (l *Logger) Debug(message string) {
	l.logger.Debug(message)
}

// Metrics logs metrics.
func (l *Logger) Metrics(metrics map[string]interface{}) {
	// Convert the metrics map to a list of key-value pairs
	var keyValuePairs []interface{}
	for key, value := range metrics {
		keyValuePairs = append(keyValuePairs, key, value)
	}

	// Log the metrics using hclog.Logger
	l.logger.Info("Metrics logged", keyValuePairs...)
}
