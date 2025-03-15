package client

import (
	"go-man/internal/config"

	"github.com/sashabaranov/go-openai"
)


func LoadClient() *openai.Client {
	cltConfig := openai.DefaultConfig(config.Config.APIKey)
	cltConfig.BaseURL = config.Config.BaseURL
	client := openai.NewClientWithConfig(cltConfig)

	return client
}
