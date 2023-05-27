package main

import (
	"context"
	"log"

	openai "github.com/sashabaranov/go-openai"

	"chat-gpt/build"
)

func main() {
	cfg := build.MakeConfig()

	cli := openai.NewClient(cfg.Flags.OpenAIToken)

	resp, err := cli.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: "Hello! This is an API integration made in GoLang",
				},
			},
		},
	)
	if err != nil {
		log.Println("CreateChatCompletion error:", err)
		return
	}

	log.Println(resp.Choices[0].Message.Content)
}
