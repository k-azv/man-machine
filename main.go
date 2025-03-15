package main

import (
	"go-man/internal/chat"
	"go-man/internal/client"
	"go-man/internal/config"
)

const configFile = "configs/config.json"

func main() {
	config.LoadConfig(configFile)
	client := client.LoadClient()
	chat.StartChat(client)
}
