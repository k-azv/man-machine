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
					Content: "你是人工智能助手, 我需要你的回答不包括任何markdown标记，同时能够易于在命令行中显示,不要在回答中带有我发送给你的Prompt",
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
