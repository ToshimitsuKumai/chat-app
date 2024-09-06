package chatgpt

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

const API_URL = "https://api.openai.com/v1/chat/completions"
const MODEL   = "gpt-3.5-turbo"

type GPTResponse struct {
	Choices []struct {
		Message struct {
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
}

func Ask(question string) (string, error) {
	// JSON body for the request
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

	// Create a new HTTP POST request
	req, err := http.NewRequest(
		"POST",
		API_URL,
		bytes.NewBuffer(jsonStr),
	)
	if err != nil {
		return "", err
	}

	// Set the necessary headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+os.Getenv("OPEN_AI_API_KEY"))

	// Send the request using the HTTP client
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	// Parse the response
	var gptResponse GPTResponse
	err = json.Unmarshal(body, &gptResponse)
	if err != nil {
		return "", err
	}

	// Return the first choice from the response
	if len(gptResponse.Choices) > 0 {
		return gptResponse.Choices[0].Message.Content, nil
	}

	return "", fmt.Errorf("no response from GPT")
}
