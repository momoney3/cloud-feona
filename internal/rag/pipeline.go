// Package rag for for prompt
package rag

import (
	"context"
	"fmt"
	"strings"
)

type Pipeline struct {
	Embedder Embedder
	Store    VectorStore
	LLM      LLM
}

func (p *Pipeline) Ask(ctx context.Context, question string, stream func(string)) error {
	vec, err := p.Embedder.Embed(ctx, question)
	if err != nil {
		return err
	}

	docs, err := p.Store.Search(ctx, vec, 5)
	if err != nil {
		return err
	}

	var parts []string
	for _, d := range docs {
		parts = append(parts, d.Text)
	}

	prompt := fmt.Sprintf(`Use the documentation context to answer. Context: %s Question %s`, strings.Join(parts, "\n\n"), question)
	return p.LLM.Stream(ctx, prompt, stream)
}
