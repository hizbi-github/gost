package models

import (
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ApiGenericResponse struct {
	Message string `json:"message"`
	Data    any    `json:"data"`
}

type HttpRequest struct {
	Url     string
	Headers http.Header
	//Headerss map[string][]string
	Body []byte
}

type HttpResponse struct {
	Headers http.Header
	Body    []byte
}

type SomeMongoDocument struct {
	Id        primitive.ObjectID `bson:"_id"`
	SomeKey   string             `bson:"some_key"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
}

type OpenAiChatRequest struct {
	Model       string  `json:"model"`
	Temperature float64 `json:"temperature"`
	Messages    []struct {
		Role    string `json:"role"`
		Content string `json:"content"`
	} `json:"messages"`
}

type OpenAiChatResponse struct {
	ID      string `json:"id"`
	Object  string `json:"object"`
	Created int    `json:"created"`
	Model   string `json:"model"`
	Choices []struct {
		Index   int `json:"index"`
		Message struct {
			Role    string `json:"role"`
			Content string `json:"content"`
		} `json:"message"`
		Logprobs     any    `json:"logprobs"`
		FinishReason string `json:"finish_reason"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
	SystemFingerprint string `json:"system_fingerprint"`
}

type OpenAiImageRequest struct {
	Model          string `json:"model"`
	Prompt         string `json:"prompt"`
	N              int    `json:"n"`
	Size           string `json:"size"`
	Style          string `json:"style"`
	Quality        string `json:"quality"`
	ResponseFormat string `json:"response_format"`
}

type OpenAiImageResponse struct {
	Created int `json:"created"`
	Data    []struct {
		//B64JSON string `json:"b64_json"`
		ImageUrl string `json:"url"`
	} `json:"data"`
}
