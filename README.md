Cloud Feona ☁️✨
Cloud Feona is a high-performance AI search aggregator written in Go. It leverages Go's lightweight concurrency (Goroutines) to fetch, parse, and synthesise search results from multiple providers simultaneously using Large Language Models (LLMs).
🚀 Features

    Concurrent Retrieval: Uses Goroutines to query multiple search engines (Google, Tavily, Bing) in parallel for ultra-low latency.
    LLM Provider Agnostic: Native support for OpenAI, Anthropic, and Ollama (local LLMs) via structured Go interfaces.
    Streaming Responses: Utilises Server-Sent Events (SSE) to stream AI-generated answers to the client in real-time.
    Memory Efficient: Built with a minimal footprint, suitable for edge deployment or high-scale cloud environments.

🛠️ Tech Stack

    Language: Go (Golang) 1.21+
    Framework: Gin or Echo for the REST API.
    AI SDK: LangChainGo for LLM orchestration.
    Frontend: React/Next.js (connected via API).

📦 Getting Started
Prerequisites

    Go 1.21 or higher
    API Keys (OpenAI, Anthropic, or Tavily)

Installation

    Clone the Source
    bash

    git clone https://github.com
    cd cloud-feona

    Use code with caution.
    Install Dependencies
    bash

    go mod tidy

    Use code with caution.
    Configure Environment
    bash

    export OPENAI_API_KEY="your-key"
    export SEARCH_API_KEY="your-key"

    Use code with caution.
    Run the Server
    bash

    go run cmd/api/main.go

Use code with caution.
