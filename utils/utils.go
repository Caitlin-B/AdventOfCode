package utils

import (
    "go.uber.org/zap"
)

// Logger wraps zap logger
type Logger struct {
    *zap.Logger
}

// Error logs with the error log level.
func (l *Logger) Error(msg string, fields ...zap.Field) {
    l.Logger.Error(msg, fields...)
}
