package logger

import (
	"log/slog"
	"os"
)

type Logger struct {
	echo *slog.Logger
}

func NewLogger() *Logger {
	return &Logger{
		echo: slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelDebug,
		})),
	}
}

func (l *Logger) Info(msg string, args ...any) {
	l.echo.Info(msg, args...)
}

func (l *Logger) Debug(msg string, args ...any) {
	l.echo.Debug(msg, args...)
}
