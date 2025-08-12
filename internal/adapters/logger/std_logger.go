package logger

import (
	"log"
	"userregisterapi/internal/app/ports"
)

// StdLogger is a minimal logger using the standard library.
var _ ports.Logger = (*StdLogger)(nil)

type StdLogger struct{}

func NewStdLogger() *StdLogger { return &StdLogger{} }

func (l *StdLogger) Info(msg string, fields map[string]any) {
	log.Printf("[INFO] %s %v", msg, fields)
}

func (l *StdLogger) Error(msg string, fields map[string]any) {
	log.Printf("[ERROR] %s %v", msg, fields)
}
