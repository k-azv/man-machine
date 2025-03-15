package chat

import (
	"context"
	"fmt"
	"go-man/internal/config"
	"os"

	"github.com/sashabaranov/go-openai"
)

// StartChat send a message to the model and prints the response
func StartChat(client *openai.Client) {
	if len(os.Args) < 2 {
		fmt.Println("Please provide a message to chat with the model")
		return
	}
	content := os.Args[1]
	resp := chat(client, content)
	showResponse(resp)
}

func chat(client *openai.Client, content string) openai.ChatCompletionResponse {
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: config.Config.Model,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleSystem,
					Content: config.Config.Prompt,
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: content,
				},
			},
		},
	)
	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return openai.ChatCompletionResponse{}
	}

	return resp
}

func showResponse(resp openai.ChatCompletionResponse) {
	fmt.Println(resp.Choices[0].Message.Content)

}
