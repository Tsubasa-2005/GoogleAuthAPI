package main

import (
	"context"
	"log/slog"
	"os"

	"GoogleAuthAPI/cmd"
)

func main() {
	var slogOpts slog.HandlerOptions
	slogOpts.Level = slog.LevelDebug
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slogOpts))
	slog.SetDefault(logger)

	ctx := context.Background()

	if err := cmd.RootCmd(ctx).Execute(); err != nil {
		slog.ErrorContext(ctx, err.Error())
		os.Exit(1)
	}
}
