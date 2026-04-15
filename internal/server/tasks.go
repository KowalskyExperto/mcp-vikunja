package server

import (
	"context"

	"github.com/kowalskyexperto/mcp-vikunja/internal/api"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

type ListTasksByProjectInput struct {
	ProjectID int `json:"project_id" jsonschema:"ID Project of Vikunja"`
}

type ListTasksByProjectOutput struct {
	Tasks api.Tasks `json:"tasks"`
}

func (s *VikunjaServer) ListTasksByProject(ctx context.Context, req *mcp.CallToolRequest, input ListTasksByProjectInput) (*mcp.CallToolResult, ListTasksByProjectOutput, error) {
	tasks, err := s.client.GetTasksByProject(input.ProjectID)
	if err != nil {
		result := &mcp.CallToolResult{}
		result.SetError(err)
		return result, ListTasksByProjectOutput{}, nil
	}
	return nil, ListTasksByProjectOutput{Tasks: tasks}, nil
}

type ListTasksBySearchInput struct {
	Search string `json:"search" jsonschema:"Task name to search of Vikunja"`
}
type ListTasksBySearchOutput struct {
	Tasks api.Tasks `json:"tasks"`
}

func (s *VikunjaServer) ListTasksBySearch(ctx context.Context, req *mcp.CallToolRequest, input ListTasksBySearchInput) (*mcp.CallToolResult, ListTasksBySearchOutput, error) {
	tasks, err := s.client.SearchTasks(input.Search)
	if err != nil {
		result := &mcp.CallToolResult{}
		result.SetError(err)
		return result, ListTasksBySearchOutput{}, nil
	}
	return nil, ListTasksBySearchOutput{Tasks: tasks}, nil
}

type CreateTaskInput struct {
	ProjectID   int    `json:"project_id" jsonschema:"ID Project of Vikunja"`
	Description string `json:"description,omitempty" jsonschema:"Task description"`
	DueDate     string `json:"due_date,omitempty" jsonschema:"Task due date in format ISO 8601 YYYY-MM-DDTHH:MM:SSZ"`
	Priority    int    `json:"priority,omitempty" jsonschema:"The task priority, 1 for low, 5 for critic"`
	Title       string `json:"title" jsonschema:"The title of the task"`
}
type CreateTaskOutput struct {
	Task api.Task `json:"task"`
}

func (s *VikunjaServer) CreateTask(ctx context.Context, req *mcp.CallToolRequest, input CreateTaskInput) (*mcp.CallToolResult, CreateTaskOutput, error) {
	newTask := api.TaskInput{
		Description: input.Description,
		DueDate:     input.DueDate,
		Priority:    input.Priority,
		Title:       input.Title,
	}
	task, err := s.client.CreateTask(input.ProjectID, newTask)
	if err != nil {
		result := &mcp.CallToolResult{}
		result.SetError(err)
		return result, CreateTaskOutput{}, nil
	}
	return nil, CreateTaskOutput{Task: task}, nil
}
