package main

import (
	"github.com/google/generative-ai-go/genai"
	"google.golang.org/genai"
)

const GenaiModel = "gemini-1.5-flash"

type App struct {
	client *genai.Client
	model  *genai.GenerativeModel
	cs     *genai.ChatSession
}

func main() {
	// err = godotenv.Load()
}
