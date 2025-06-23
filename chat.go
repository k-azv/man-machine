package main

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/k-azv/man-machine/config"
	"github.com/peterh/liner"
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

// ContinuousChat handles a continuous chat.
func ContinuousChat(client *openai.Client, content string, pg *PromptGenerator, cfg config.Config) error {
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
		displayResponse(resp, &messages)
	} else {
		return fmt.Errorf("receive a response from chat")
	}

	pg.UpdatePromptForContinuousChat()
	messages[0].Content = pg.Mam()

	// Initialize liner
	line := liner.NewLiner()
	defer line.Close()

	// Allow user to quit continuous chat using Ctrl+C
	line.SetCtrlCAborts(true)

	// Exit continuous chat using "exit"
	for {
		fmt.Println()
		userInput, err := line.Prompt("> ")
		if err != nil {
			if err == liner.ErrPromptAborted || err == io.EOF {
				break
			}
			return fmt.Errorf("read input: %w", err)
		}

		userInput = strings.TrimSpace(userInput)
		if userInput == "exit" {
			break
		} else if userInput == "" {
			continue
		}

		messages = append(messages, openai.ChatCompletionMessage{
			Role:    openai.ChatMessageRoleUser,
			Content: userInput,
		})

		resp, err = createChatStream(client, cfg, messages)
		if err != nil {
			return fmt.Errorf("create chat stream: %w", err)
		}

		if err := displayResponse(resp, &messages); err != nil {
			return fmt.Errorf("display response: %w", err)
		}
	}
	return nil
}

// createChatStream creates a chat completion stream with given messages.
func createChatStream(client *openai.Client, cfg config.Config, messages []openai.ChatCompletionMessage) (*openai.ChatCompletionStream, error) {
	for {
		stream, err := client.CreateChatCompletionStream(
			context.Background(),
			openai.ChatCompletionRequest{
				Model:    cfg.Model,
				Messages: messages,
			},
		)

		if err == nil {
			return stream, err
		}

		var apiErr *openai.APIError
		// If http status code is 429, ask user whether to retry
		if errors.As(err, &apiErr) && apiErr.HTTPStatusCode == 429 {

			allowRetry, confirmErr := confirmRetry(os.Stdout, os.Stdin, *apiErr)
			if confirmErr != nil {
				return nil, fmt.Errorf("%v: confirm retry: %w", apiErr, confirmErr)
			}

			if !allowRetry {
				return nil, fmt.Errorf("user aborted retry: %w", apiErr)
			}

			continue
		}

		return nil, fmt.Errorf("create chat completion stream: %w", err)
	}
}

func confirmRetry(w io.Writer, r io.Reader, err openai.APIError) (bool, error) {
	fmt.Fprintf(w, "%v\n", err)
	fmt.Fprint(w, "Press Enter to retry:\n")

	scanner := bufio.NewScanner(r)
	scanner.Scan()
	return scanner.Text() == "", fmt.Errorf("confirm whether retry: %w", scanner.Err())
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
