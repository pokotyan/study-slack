package connpass

import (
	"github.com/spf13/cobra"
)

func NewConnpassCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "connpass",
		Short: "connpassの勉強会を検索",
		RunE:  getEvent,
	}

	cmd.AddCommand(
		newGetEventCmd(),
		newPostSlackCmd(),
	)

	return cmd
}

func newGetEventCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "search (<yyyymm> | <yyyymmdd>) <keyword>",
		Short: "キーワードで検索 ex: search 201209 python",
		RunE:  getEvent,
		Args:  cobra.MinimumNArgs(1),
	}

	return cmd
}

func newPostSlackCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "slack",
		Short: "slackに勉強会を通知",
		RunE:  PostSlack,
	}

	return cmd
}
