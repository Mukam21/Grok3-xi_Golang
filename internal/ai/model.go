package ai

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type AIClient struct {
	apiKey  string
	baseURL string
}

func NewAIClient(apiKey, baseURL string) *AIClient {
	return &AIClient{apiKey: apiKey, baseURL: baseURL}
}

func (c *AIClient) Predict(text string) (string, error) {
	payload := map[string]string{"input": text}
	jsonData, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", c.baseURL+"/chat_xi", bytes.NewBuffer(jsonData))
	if err != nil {
		return "", err
	}
	req.Header.Set("Authorization", "Bearer "+c.apiKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var result struct {
		Output string `json:"output"`
	}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return "", err
	}

	return result.Output, nil
}
