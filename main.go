package main

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/pokotyan/study-slack/server"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Errorf("err %s", "load error .env")
	}

	server.Init()
}
