package gptAPI

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

// GPTClient is a struct that deals with GPT-completions related chat
type GPTClient struct {
	key string
}

// NewGPTClient creates a new GPTClient
func NewGPTClient(key string) *GPTClient {
	return &GPTClient{key: key}
}

// GPTRequest is the struct that contains the required fields for making a request to OpenAI API
type GPTRequest struct {
	Model    string `json:"model"`
	Messages []struct {
		Role    string `json:"role"`
		Content string `json:"content"`
	} `json:"messages"`
}

// GPTResponse is the returned response from OpenAI
type GPTResponse struct {
	Id      string `json:"id"`
	Object  string `json:"object"`
	Created int    `json:"created"`
	Model   string `json:"model"`
	Usage   struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
	Choices []struct {
		Message struct {
			Role    string `json:"role"`
			Content string `json:"content"`
		} `json:"message"`
		FinishReason string `json:"finish_reason"`
		Index        int    `json:"index"`
	} `json:"choices"`
}

// NewGPTRequest creates a new GPTRequest struct
func NewGPTRequest(message string) *GPTRequest {
	return &GPTRequest{
		Model: "gpt-3.5-turbo",
		Messages: []struct {
			Role    string `json:"role"`
			Content string `json:"content"`
		}{
			{
				Role:    "user",
				Content: message,
			},
		},
	}
}

// GetAnswers gets the AI generated answer from OpenAI
func (g *GPTClient) GetAnswers(message string) (string, error) {

	// Set up the API request
	apiUrl := "https://api.openai.com/v1/chat/completions"
	gptRequest := NewGPTRequest(message)
	requestBody, err := json.Marshal(*gptRequest)
	if err != nil {
		return "", fmt.Errorf("failed to marshal JSON request body: %v", err)
	}

	// Send the API request
	httpClient := &http.Client{}
	request, err := http.NewRequest("POST", apiUrl, bytes.NewBuffer(requestBody))
	if err != nil {
		return "", fmt.Errorf("failed to create JSON request: %v", err)
	}
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", g.key))

	response, err := httpClient.Do(request)
	if err != nil {
		return "", fmt.Errorf("failed to send HTTP request: %v", err)
	}
	defer func(r *http.Response) {
		if err := r.Body.Close(); err != nil {
			fmt.Printf("failed to close HTTP response body: %v", err)
		}
	}(response)

	// Parse the API response into type GPTResponse
	responseBody, _ := io.ReadAll(response.Body)
	var responseJson GPTResponse
	if err := json.Unmarshal(responseBody, &responseJson); err != nil {
		return "", fmt.Errorf("failed to unmarshall HTTP response body: %v", err)
	}
	var answers string
	// Get answers from gpt response
	if len(responseJson.Choices) > 0 {
		answers = strings.TrimSpace(responseJson.Choices[0].Message.Content)
	}
	return answers, nil
}
