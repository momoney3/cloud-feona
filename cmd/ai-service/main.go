package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"

	"google.golang.org/grpc"
)

func main() {
	// ---- Ollama request (your feature) ----
	body := map[string]any{
		"model":  "llama3",
		"prompt": "Explain kubernetes simply",
	}

	b, err := json.Marshal(body)
	if err != nil {
		log.Fatalf("failed to marshal request: %v", err)
	}

	resp, err := http.Post(
		"http://localhost:11434/api/generate",
		"application/json",
		bytes.NewReader(b),
	)
	if err != nil {
		log.Fatalf("failed to call ollama: %v", err)
	}
	defer resp.Body.Close()

	var r map[string]any
	if err := json.NewDecoder(resp.Body).Decode(&r); err != nil {
		log.Printf("failed to decode response: %v", err)
	}

	fmt.Println("response:", r)

	// ---- gRPC server (new base) ----
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}

	server := grpc.NewServer()

	log.Println("RAG server running on :50051")

	if err := server.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
