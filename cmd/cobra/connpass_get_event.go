package cmd

import (
	"context"
	"errors"
	"fmt"

	"github.com/spf13/cobra"

	connpassEvent "github.com/pokotyan/connpass-map-api/infrastructure/connpass/event"
	usecase "github.com/pokotyan/connpass-map-api/usecase/connpass/event"
)

func newConnpassCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "connpass",
		Short: "connpassの勉強会を検索",
		RunE:  getEvent,
	}

	cmd.AddCommand(
		newGetEventCmd(),
	)

	return cmd
}

func newGetEventCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "keyword <Keyword>",
		Short: "キーワードで検索",
		RunE:  getEvent,
	}

	return cmd
}

func getEvent(cmd *cobra.Command, args []string) error {
	if len(args) == 0 {
		return errors.New("keyword is required")
	}

	u := usecase.AssignGetEventWithUsecase()
	ctx := context.Background()
	keyword := args[0]

	res := u.GetEvent(ctx, connpassEvent.ReqParam{
		Keyword: keyword,
	})

	fmt.Printf("%v", res)

	return nil
}
