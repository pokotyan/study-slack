package cmd

import (
	"fmt"
	"os"

	"github.com/pokotyan/study-slack/cmd/cobra/connpass"
	"github.com/spf13/cobra"
)

var (
	rootCmd = newRootCmd()
)

func newRootCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "cmd",
		Short: "コマンドツールです",
		Long:  "コマンドツールです",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("コマンドツールです")
			return nil
		},
	}
}

func init() {
	rootCmd.AddCommand(connpass.NewConnpassCmd())
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
