package ai

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"os"
)

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type GroqRequest struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
}

type GroqResponse struct {
	Choices []struct {
		Message Message `json:"message"`
	} `json:"choices"`
}

func GetGroqResponse(userInput string) string {
	apiKey := os.Getenv("GROQ_API_KEY")

	body := GroqRequest{
		Model: "llama-3.3-70b-versatile",
		Messages: []Message{
			{
				Role: "system",
				Content: `
You are a clean chat assistant.

Rules:
- No markdown (** ## ---)
- Use simple text only
- Use "•" for bullet points
- Keep responses short and readable
`,
			},
			{Role: "user", Content: userInput},
		},
	}

	jsonData, _ := json.Marshal(body)

	req, _ := http.NewRequest(
		"POST",
		"https://api.groq.com/openai/v1/chat/completions",
		bytes.NewBuffer(jsonData),
	)

	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "API error"
	}
	defer resp.Body.Close()

	resBody, _ := io.ReadAll(resp.Body)

	var result GroqResponse
	json.Unmarshal(resBody, &result)

	if len(result.Choices) == 0 {
		return "No response"
	}

	return result.Choices[0].Message.Content
}