package logger

import (
	"log/slog"
	"os"

	"github.com/google/uuid"
)

func NewLogger() *slog.Logger {
	return slog.New(slog.NewJSONHandler(os.Stderr, nil)).With("trace_id", uuid.NewString())
}

