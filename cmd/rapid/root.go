package main

import (
	"fmt"
	"os"

	"github.com/abyssparanoia/application-boilerplate/cmd/rapid/server"
	"github.com/spf13/cobra"
)

func newCmdRoot() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "rapid",
		Short: "cli tool for rapid-go",
		Long:  "cli tool for rapid-go",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				cmd.HelpFunc()(cmd, args)
			}
		},
	}
	cmd.AddCommand(server.NewServerCmd())
	return cmd
}

func execute() {
	cmd := newCmdRoot()
	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
