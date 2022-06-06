package main

import (
	"github.com/abyssparanoia/application-boilerplate/internal/server/application/api"
	"github.com/spf13/cobra"
)

// NewServerCmd ... new default HTTP cmd
func NewServerCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "server",
		Short: "cli http server",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				cmd.HelpFunc()(cmd, args)
			}
		},
	}
	cmd.AddCommand(&cobra.Command{
		Use:   "run",
		Short: "running http server",
		Run: func(cmd *cobra.Command, args []string) {
			api.Run()
		},
	})
	return cmd
}
