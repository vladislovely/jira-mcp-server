package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/mark3labs/mcp-go/server"

	"vladislove-mcp/internal/config"
	"vladislove-mcp/internal/registry"
)

func NewMCPServer() *server.MCPServer {
	hooks := &server.Hooks{}

	instructionsText := ""

	mcpServer := server.NewMCPServer(
		"Vladislove MCP Server",
		"1.0.0",
		server.WithResourceCapabilities(true, true),
		server.WithPromptCapabilities(true),
		server.WithToolCapabilities(true),
		server.WithLogging(),
		server.WithHooks(hooks),
		server.WithInstructions(instructionsText),
	)

	registry.RegisterAll(mcpServer)

	return mcpServer
}

func main() {
	var transport string
	flag.StringVar(&transport, "t", "stdio", "Transport type (stdio or http)")
	flag.StringVar(&transport, "transport", "stdio", "Transport type (stdio or http)")
	flag.Parse()

	cfg := config.LoadConfig()

	fmt.Printf("Transport: %s\n", transport)

	mcpServer := NewMCPServer()

	if transport == "http" {
		httpServer := server.NewStreamableHTTPServer(mcpServer)

		readTimeoutConst := 15
		idleTimeoutConst := 60

		srv := &http.Server{
			Addr:         ":" + cfg.McpOutputPort,
			Handler:      httpServer,
			ReadTimeout:  time.Duration(readTimeoutConst) * time.Second,
			WriteTimeout: time.Duration(readTimeoutConst) * time.Second,
			IdleTimeout:  time.Duration(idleTimeoutConst) * time.Second,
		}

		log.Printf("HTTP server listening on " + srv.Addr)

		if err := srv.ListenAndServe(); err != nil {
			log.Fatalf("Server error: %v", err)
		}
	} else {
		if err := server.ServeStdio(mcpServer); err != nil {
			log.Fatalf("Server error: %v", err)
		}
	}
}
