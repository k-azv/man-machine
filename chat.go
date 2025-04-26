package main

import (
	"context"
	"fmt"
	"io"

	"github.com/k-azv/man-machine/config"
	"github.com/k-azv/man-machine/prompt"
	"github.com/sashabaranov/go-openai"
)

func initClient() *openai.Client {
	cfg := openai.DefaultConfig(config.Config.APIKey)
	cfg.BaseURL = config.Config.BaseURL
	c := openai.NewClientWithConfig(cfg)
	return c
}

// Chat handles a single message chat using command-line arguments.
func Chat(client *openai.Client, content string) error {
	messages := []openai.ChatCompletionMessage{
		{
			Role:    openai.ChatMessageRoleSystem,
			Content: prompt.Mam(),
		},
		{
			Role:    openai.ChatMessageRoleUser,
			Content: content,
		},
	}
	resp, err := createChatStream(client, messages)
	if err != nil {
		return fmt.Errorf("create chat stream: %w", err)
	}
	if resp != nil {
		displayResponse(resp, nil)
	} else {
		return fmt.Errorf("receive a response from chat")
	}
	return nil
}

// createChatStream creates a chat completion stream with given messages.
func createChatStream(client *openai.Client, messages []openai.ChatCompletionMessage) (*openai.ChatCompletionStream, error) {
	stream, err := client.CreateChatCompletionStream(
		context.Background(),
		openai.ChatCompletionRequest{
			Model:    config.Config.Model,
			Messages: messages,
		},
	)
	if err != nil {
		return nil, fmt.Errorf("create chat completion stream: %w", err)
	}
	return stream, nil
}

// displayResponse processes and displays the response from the chat stream.
func displayResponse(stream *openai.ChatCompletionStream, messages *[]openai.ChatCompletionMessage) error {
	defer stream.Close()

	var responseContent string
	for {
		recv, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("receiving from stream: %w", err)
		}

		if len(recv.Choices) > 0 {
			content := recv.Choices[0].Delta.Content
			fmt.Print(content)
			responseContent += content
		}
	}

	if messages != nil {
		*messages = append(*messages, openai.ChatCompletionMessage{
			Role:    openai.ChatMessageRoleAssistant,
			Content: responseContent,
		})
	}
	return nil
}
