// Package embed is used to ?
package embed

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"net/http"
)

type Client struct {
	url string
}

func New(url string) *Client {
	return &Client{url: url}
}

func (c *Client) Embed(ctx context.Context, text string) ([]float32, error) {
	body := map[string]any{
		"model":  "nomic-embed-text",
		"prompt": text,
	}

	b, _ := json.Marshal(body)

	req, _ := http.NewRequestWithContext(ctx, "POST", c.url+"/api/embeddings", bytes.NewReader(b))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer func() {
		if err := resp.Body.Close(); err != nil {
			log.Printf("failed to close: %v", err)
		}
	}()

	var r struct {
		Embedding []float32 `json:"embedding"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&r); err != nil {
		return nil, err
	}
	return r.Embedding, nil
}
