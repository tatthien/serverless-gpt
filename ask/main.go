package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	gogpt "github.com/sashabaranov/go-gpt3"
)

type Request events.APIGatewayProxyRequest
type Response events.APIGatewayProxyResponse

type RequestBody struct {
	Prompt string `json:"prompt"`
}

func Handler(req Request) (Response, error) {
	var data RequestBody
	if err := json.Unmarshal([]byte(req.Body), &data); err != nil {
		return Response{}, err
	}

	answer, err := getAnswer(data.Prompt)
	if err != nil {
		fmt.Println("[Error]", err.Error())
		return Response{}, err
	}

	body, _ := json.Marshal(map[string]string{"answer": answer})

	return Response{
		StatusCode: 200,
		Body:       string(body),
		Headers: map[string]string{
			"Content-Type":                     "application/json",
			"Access-Control-Allow-Origin":      "*",
			"Access-Control-Allow-Credentials": "true",
		},
	}, nil
}

func getAnswer(prompt string) (string, error) {
	c := gogpt.NewClient(os.Getenv("OPENAI_API_KEY"))
	req := gogpt.CompletionRequest{
		Model:       gogpt.GPT3TextDavinci002,
		MaxTokens:   200,
		Prompt:      prompt,
		Temperature: 0.5,
	}

	resp, err := c.CreateCompletion(context.TODO(), req)
	if err != nil {
		return "", err
	}

	answer := "no answers"
	if len(resp.Choices) > 0 {
		answer = strings.TrimSpace(resp.Choices[0].Text)
		answer = strings.Trim(answer, "\"")
	}

	return answer, nil
}

func main() {
	lambda.Start(Handler)
}
