// Package llm for olaama
package llm

import (
	"bufio"
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

func (c *Client) Stream(ctx context.Context, prompt string, cb func(string)) error {
	body := map[string]any{
		"model":  "llama3",
		"promt":  prompt,
		"stream": true,
	}

	b, _ := json.Marshal(body)

	req, _ := http.NewRequestWithContext(ctx, "POST", c.url+"/api/generate", bytes.NewReader(b))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	defer func() {
		if err := resp.Body.Close(); err != nil {
			log.Printf("body did not close: %v", err)
		}
	}()

	scanner := bufio.NewScanner(resp.Body)

	for scanner.Scan() {
		var chunk struct {
			Response string `json:"response"`
		}

		if err := json.Unmarshal(scanner.Bytes(), &chunk); err != nil {
			log.Printf("unmarshal fialed: %v", err)
		}
		cb(chunk.Response)
	}
	return nil
}
