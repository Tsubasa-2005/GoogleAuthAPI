package cmd

import (
	"context"

	"github.com/spf13/cobra"
)

func RootCmd(ctx context.Context) *cobra.Command {
	c := &cobra.Command{
		Use: "cli",
	}

	c.AddCommand(
		serverCmd(ctx),
	)

	return c
}
