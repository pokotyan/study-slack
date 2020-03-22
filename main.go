package main

import (
	"flag"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/pokotyan/study-slack/server"

	cmd "github.com/pokotyan/study-slack/cmd/cobra"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Errorf("err %s", "load error .env")
	}

	flag.Parse()

	switch flag.Arg(0) {
	case "web":
		server.Init()
	default:
		cmd.Execute()

	}
}
