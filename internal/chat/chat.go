package chat

import (
	"context"
	"fmt"
	"go-man/internal/config"
	"io"
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

func chat(client *openai.Client, content string) *openai.ChatCompletionStream {
	stream, err := client.CreateChatCompletionStream(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: "ep-m-20250314205037-rqnsn",
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
		fmt.Printf("stream chat error: %v\n", err)
		return &openai.ChatCompletionStream{}
	}
	return stream
}

func showResponse(stream *openai.ChatCompletionStream) {
	defer stream.Close()

	for {
		recv, err := stream.Recv()
		if err == io.EOF {
			return
		}
		if err != nil {
			fmt.Printf("Stream chat error: %v\n", err)
			return
		}

		if len(recv.Choices) > 0 {
			fmt.Print(recv.Choices[0].Delta.Content)
		}
	}
}
