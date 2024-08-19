package slogger

import (
	"log/slog"
	"os"
)

func NewLogger() *slog.Logger {
	logger := slog.New(
		&ContextHandler{
			slog.NewJSONHandler(
				os.Stdout,
				&slog.HandlerOptions{
					AddSource: true,
					Level:     slog.LevelDebug,
				},
			),
		},
	)

	return logger
}

func Err(err error) slog.Attr {
	return slog.Attr{
		Key:   "error",
		Value: slog.StringValue(err.Error()),
	}
}
