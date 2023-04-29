package openai

import (
	"context"
	"fmt"
)

type Summarizer struct {
	client *Client
}

func NewSummarizer(client *Client) *Summarizer {
	return &Summarizer{
		client: client,
	}
}

func (s *Summarizer) GetSummary(ctx context.Context, url string) (string, error) {
	r := CreateCompletionsRequest{
		Model: "gpt-3.5-turbo",
		Messages: []Message{
			{
				Role:    "user",
				Content: fmt.Sprintf("Summarize article at %s", url),
			},
		},
		Temperature: 0.7,
	}

	completions, err := s.client.CreateCompletions(ctx, r)
	if err != nil {
		return "", fmt.Errorf("could not create completions: %w", err)
	}
	if completions.Error.Message != "" {
		return "", fmt.Errorf("could not create completions: %s", completions.Error.Message)
	}

	for i := range completions.Choices {
		return completions.Choices[i].Message.Content, nil
	}

	return "", nil
}
