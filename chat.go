package main

import (
	"context"
	"fmt"
	"io"

	"github.com/k-azv/man-machine/config"
	"github.com/sashabaranov/go-openai"
)

// initClient initializes an go-openai client with the given config.Config.
func initClient(cfg config.Config) *openai.Client {
	clientConfig := openai.DefaultConfig(cfg.APIKey)
	clientConfig.BaseURL = cfg.BaseURL
	return openai.NewClientWithConfig(clientConfig)

}

// Chat handles a single message chat using command-line arguments.
func Chat(client *openai.Client, content string, pg *PromptGenerator, cfg config.Config) error {
	messages := []openai.ChatCompletionMessage{
		{
			Role:    openai.ChatMessageRoleSystem,
			Content: pg.Mam(),
		},
		{
			Role:    openai.ChatMessageRoleUser,
			Content: content,
		},
	}
	resp, err := createChatStream(client, cfg, messages)
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
func createChatStream(client *openai.Client, cfg config.Config, messages []openai.ChatCompletionMessage) (*openai.ChatCompletionStream, error) {
	stream, err := client.CreateChatCompletionStream(
		context.Background(),
		openai.ChatCompletionRequest{
			Model:    cfg.Model,
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
