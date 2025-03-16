package chat

import (
	"context"
	"fmt"
	"go-man/internal/config"
	"io"
	"os"

	"github.com/chzyer/readline"
	"github.com/sashabaranov/go-openai"
)

// StartChat determines the mode (single or interactive) and starts the chat.
func StartChat(client *openai.Client) (err error) {
	if len(os.Args) > 1 {
		return singleChat(client, os.Args[1])
	}
	return interactiveChat(client)
}

// singleChat handles a single message chat using command-line arguments.
func singleChat(client *openai.Client, content string) error {
	messages := []openai.ChatCompletionMessage{
		{
			Role:    openai.ChatMessageRoleSystem,
			Content: config.Config.Prompt,
		},
		{
			Role:    openai.ChatMessageRoleUser,
			Content: content,
		},
	}

	resp := createChatStream(client, messages)
	if resp != nil {
		displayResponse(resp, nil)
	}
	return nil
}

// interactiveChat handles an interactive chat session with context.
func interactiveChat(client *openai.Client) error {
	messages := []openai.ChatCompletionMessage{
		{
			Role:    openai.ChatMessageRoleSystem,
			Content: config.Config.Prompt,
		},
	}

	rl, err := readline.New("> ")
	if err != nil {
		return fmt.Errorf("failed to initialize readline: %w", err)
	}
	defer rl.Close()

	for {
		content, err := rl.Readline()
		if err != nil {
			return fmt.Errorf("failed to read input: %w", err)
		}
		if content == "exit" {
			break
		}

		messages = append(messages, openai.ChatCompletionMessage{
			Role:    openai.ChatMessageRoleUser,
			Content: content,
		})

		resp := createChatStream(client, messages)
		if resp == nil {
			continue
		}

		displayResponse(resp, &messages)
	}
	return nil
}

// createChatStream creates a chat completion stream with the given messages.
func createChatStream(client *openai.Client, messages []openai.ChatCompletionMessage) *openai.ChatCompletionStream {
	stream, err := client.CreateChatCompletionStream(
		context.Background(),
		openai.ChatCompletionRequest{
			Model:    config.Config.Model,
			Messages: messages,
		},
	)
	if err != nil {
		fmt.Printf("stream chat error: %v\n", err)
		return nil
	}
	return stream
}

// displayResponse processes and displays the response from the chat stream.
func displayResponse(stream *openai.ChatCompletionStream, messages *[]openai.ChatCompletionMessage) {
	defer stream.Close()

	var responseContent string
	for {
		recv, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Printf("Stream chat error: %v\n", err)
			return
		}

		if len(recv.Choices) > 0 {
			content := recv.Choices[0].Delta.Content
			fmt.Print(content)
			responseContent += content
		}
	}
	fmt.Println()

	if messages != nil {
		*messages = append(*messages, openai.ChatCompletionMessage{
			Role:    openai.ChatMessageRoleAssistant,
			Content: responseContent,
		})
	}
}
