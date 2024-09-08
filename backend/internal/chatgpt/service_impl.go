package chatgpt

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

const API_URL = "https://api.openai.com/v1/chat/completions"
const MODEL = "gpt-3.5-turbo"

type service struct{}

type GPTResponse struct {
	Choices []struct {
		Message struct {
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
}

func NewService() Service {
	return &service{}
}

func (s *service) Ask(question string) (string, error) {
	jsonData := map[string]interface{}{
		"model": MODEL,
		"messages": []map[string]string{
			{
				"role":    "user",
				"content": question,
			},
		},
	}

	jsonStr, err := json.Marshal(jsonData)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest(
		"POST",
		API_URL,
		bytes.NewBuffer(jsonStr),
	)
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+os.Getenv("OPEN_AI_API_KEY"))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var gptResponse GPTResponse
	err = json.Unmarshal(body, &gptResponse)
	if err != nil {
		return "", err
	}

	if len(gptResponse.Choices) > 0 {
		return gptResponse.Choices[0].Message.Content, nil
	}

	return "", fmt.Errorf("no response from GPT")
}
