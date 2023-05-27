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
					Content: "Hello!",
				},
			},
		},
	)
	if err != nil {
		log.Println("CreateChatCompletion error:", err)
		return
	}

	log.Println(resp)
}
