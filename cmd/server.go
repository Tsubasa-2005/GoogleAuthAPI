package cmd

import (
	"context"
	"log/slog"
	"net/http"

	"GoogleAuthAPI/pkg/infra"
	"GoogleAuthAPI/pkg/middleware"
	"GoogleAuthAPI/ui/api"
	"GoogleAuthAPI/usecase/handler"

	"github.com/spf13/cobra"
)

func serverCmd(ctx context.Context) *cobra.Command {
	return &cobra.Command{
		Use:   "serve",
		Short: "Start the HTTP server",
		RunE: func(cmd *cobra.Command, args []string) error {
			return runServer(ctx)
		},
	}
}

func runServer(ctx context.Context) error {
	dbConn := infra.ConnectDB(ctx)
	h := handler.NewHandler(dbConn)
	s := handler.NewSecurityHandler()

	srv, err := api.NewServer(h, s, api.WithMiddleware(middleware.Logging()))
	if err != nil {
		return err
	}

	corsHandler := middleware.Cors(srv)

	addr := ":8080"
	slog.InfoContext(ctx, "Starting HTTP Server", slog.String("addr", addr))
	if err := http.ListenAndServe(addr, corsHandler); err != nil {
		return err
	}

	return nil
}
