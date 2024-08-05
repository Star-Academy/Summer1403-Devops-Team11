package helper

import (
	"log/slog"
	"os"
)

func InitLogger() *slog.Logger {
	logHandler := slog.NewJSONHandler(
		os.Stdout,
		&slog.HandlerOptions{
			AddSource: true,
		},
	)
	logger := slog.New(logHandler)

	return logger
}
