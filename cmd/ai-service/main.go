package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	body := map[string]any{
		"model": "llama3",
		"promt": "Explain kubernetes simply",
	}
	b, _ := json.Marshal(body)
	resp, _ := http.Post(
		"http://localhost:11434/api/generate",
		"application/json",
		bytes.NewReader(b),
	)
	defer resp.Body.Close()

	var r map[string]any

	json.NewDecoder(resp.Body).Decode(&r)

	fmt.Println("response")
}
