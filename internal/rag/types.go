package rag

import "context"

type Embedder interface {
	Embed(ctx context.Context, text string) ([]float32, error)
}

type VectorStore interface {
	Search(ctx context.Context, vector []float32, limit int) ([]Document, error)
}

type LLM interface {
	Stream(ctx context.Context, prompt string, cb func(string)) error
}

type Document struct {
	Text string
}
