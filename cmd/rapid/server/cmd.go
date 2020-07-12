package server

import (
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
	cmd.AddCommand(newServerRunCmd())
	return cmd
}

func newServerRunCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "run",
		Short: "running http server",
		Run: func(cmd *cobra.Command, args []string) {
			run()
		},
	}
	return cmd
}
