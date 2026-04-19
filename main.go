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
	mcp.AddTool(server, &mcp.Tool{Name: "create_project", Description: "Create a new Vikunja project"}, handler.CreateProject)
	mcp.AddTool(server, &mcp.Tool{Name: "update_project", Description: "Update an existing Vikunja project by ID"}, handler.UpdateProject)
	mcp.AddTool(server, &mcp.Tool{Name: "delete_project", Description: "Delete a Vikunja project by ID"}, handler.DeleteProject)
	mcp.AddTool(server, &mcp.Tool{Name: "list_tasks_by_project", Description: "List all Vikunja Project Tasks"}, handler.ListTasksByProject)
	mcp.AddTool(server, &mcp.Tool{Name: "search_tasks", Description: "Search tasks by name"}, handler.ListTasksBySearch)
	mcp.AddTool(server, &mcp.Tool{Name: "create_task", Description: "Create a task for a project"}, handler.CreateTask)
	mcp.AddTool(server, &mcp.Tool{Name: "get_task", Description: "Get the details of a task by ID"}, handler.GetTask)
	mcp.AddTool(server, &mcp.Tool{Name: "list_labels", Description: "List all labels available to the user"}, handler.ListLabels)
	mcp.AddTool(server, &mcp.Tool{Name: "create_label", Description: "Create a new label"}, handler.CreateLabel)
	mcp.AddTool(server, &mcp.Tool{Name: "delete_label", Description: "Delete a label by ID"}, handler.DeleteLabel)
	mcp.AddTool(server, &mcp.Tool{Name: "add_label_to_task", Description: "Add an existing label to a task"}, handler.AddLabelToTask)
	mcp.AddTool(server, &mcp.Tool{Name: "remove_label_from_task", Description: "Remove a label from a task"}, handler.RemoveLabelFromTask)
	mcp.AddTool(server, &mcp.Tool{Name: "list_task_comments", Description: "List all comments on a task"}, handler.ListTaskComments)
	mcp.AddTool(server, &mcp.Tool{Name: "create_task_comment", Description: "Add a comment to a task"}, handler.CreateTaskComment)
	mcp.AddTool(server, &mcp.Tool{Name: "delete_task_comment", Description: "Delete a comment from a task"}, handler.DeleteTaskComment)
	mcp.AddTool(server, &mcp.Tool{Name: "update_task", Description: "Update a task by ID"}, handler.UpdateTask)
	mcp.AddTool(server, &mcp.Tool{Name: "delete_task", Description: "Delete a task by ID"}, handler.DeleteTask)
	server.Run(context.Background(), &mcp.StdioTransport{})
}
