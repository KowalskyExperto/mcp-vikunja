package main

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/kowalskyexperto/mcp-vikunja/internal/api"
	"github.com/kowalskyexperto/mcp-vikunja/internal/server"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env")
	}
	vikunjaClient, err := api.NewClient(os.Getenv("API_URL"), os.Getenv("API_TOKEN"))
	if err != nil {
		log.Fatalf("Error initializing client: %v", err)
	}
	handler := server.NewVikunjaHandler(vikunjaClient)
	server := mcp.NewServer(&mcp.Implementation{Name: "MCP Vikunja", Version: "v1.0.0"}, nil)
	mcp.AddTool(server, &mcp.Tool{Name: "list_projects", Description: "List all Vikunja Projects"}, handler.ListProjects)
	mcp.AddTool(server, &mcp.Tool{Name: "list_tasks_by_project", Description: "List all Vikunja Project Tasks"}, handler.ListTasksByProject)
	mcp.AddTool(server, &mcp.Tool{Name: "search_tasks", Description: "Search tasks by name"}, handler.ListTasksBySearch)
	mcp.AddTool(server, &mcp.Tool{Name: "create_task", Description: "Create a task for a project"}, handler.CreateTask)
	mcp.AddTool(server, &mcp.Tool{Name: "get_task", Description: "Get the details of a task by ID"}, handler.GetTask)
	mcp.AddTool(server, &mcp.Tool{Name: "update_task", Description: "Update a task by ID"}, handler.UpdateTask)
	mcp.AddTool(server, &mcp.Tool{Name: "delete_task", Description: "Delete a task by ID"}, handler.DeleteTask)
	server.Run(context.Background(), &mcp.StdioTransport{})
}
