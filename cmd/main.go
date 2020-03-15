package main

import (
	"fmt"

	"github.com/joho/godotenv"

	cmd "github.com/pokotyan/study-slack/cmd/cobra"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Errorf("err %s", "load error .env")

		return
	}

	cmd.Execute()
}
