package main

import (
	"go-man/internal/chat"
	"go-man/internal/client"
	"go-man/internal/config"
)

func main() {
	config.LoadConfig()
	client := client.LoadClient()
	chat.StartChat(client)
}
