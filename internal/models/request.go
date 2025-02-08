package models

type AIRequest struct {
	Prompt string `json:"prompt"`
}

type AIResponse struct {
	Response string `json:"response"`
	Error    string `json:"error,omitempty"`
}
