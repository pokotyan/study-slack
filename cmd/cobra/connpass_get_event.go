package cmd

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"unicode/utf8"

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
		Use:   "search (<yyyymm> | <yyyymmdd>) <keyword>",
		Short: "キーワードで検索 ex: search 201209 python",
		RunE:  getEvent,
		Args:  cobra.MinimumNArgs(1),
	}

	return cmd
}

func setDate(reqParam *connpassEvent.ReqParam) func(arg string) {
	return func(arg string) {
		length := utf8.RuneCountInString(arg)
		param, _ := strconv.Atoi(arg)

		if length == 6 {
			reqParam.YmList = []int{param}
		}

		if length == 8 {
			reqParam.YmdList = []int{param}
		}
	}
}

func setKeyword(reqParam *connpassEvent.ReqParam) func(arg string) {
	return func(arg string) {
		reqParam.Keyword = arg
	}
}

func getEvent(cmd *cobra.Command, args []string) error {
	if len(args) == 0 {
		return errors.New("date is required")
	}

	u := usecase.AssignGetEventWithUsecase()
	ctx := context.Background()
	reqParam := connpassEvent.ReqParam{}
	setParamList := []func(arg string){setDate(&reqParam), setKeyword(&reqParam)}

	for _, arg := range args {
		setParam := setParamList[0]
		_ = setParamList[1:]

		setParam(arg)
	}

	res := u.GetEvent(ctx, reqParam)

	fmt.Printf("%v", res)

	return nil
}
