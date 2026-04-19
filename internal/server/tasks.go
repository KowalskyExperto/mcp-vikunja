package server

import (
	"context"

	"github.com/kowalskyexperto/mcp-vikunja/internal/api"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

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

type GetTaskInput struct {
	TaskID int `json:"task_id" jsonschema:"ID Task to get"`
}

type GetTaskOutput struct {
	Task api.TaskDetail `json:"task"`
}

func (s *VikunjaServer) GetTask(ctx context.Context, req *mcp.CallToolRequest, input GetTaskInput) (*mcp.CallToolResult, GetTaskOutput, error) {
	task, err := s.client.GetTask(input.TaskID)
	if err != nil {
		result := &mcp.CallToolResult{}
		result.SetError(err)
		return result, GetTaskOutput{}, nil
	}
	return nil, GetTaskOutput{Task: task}, nil
}

type UpdateTaskInput struct {
	TaskID int `json:"id" jsonschema:"ID Task to update"`

	Description string          `json:"description,omitempty" jsonschema:"Task description"`
	Done        *bool           `json:"done,omitempty" jsonschema:"Whether a task is done or not"`
	DueDate     string          `json:"due_date,omitempty" jsonschema:"When this task is due in format ISO 8601 YYYY-MM-DDTHH:MM:SSZ"`
	EndDate     string          `json:"end_date,omitempty" jsonschema:"When this task ends"`
	HexColor    string          `json:"hex_color,omitempty" jsonschema:"The task color in hex <= 7 characters"`
	IsFavorite  *bool           `json:"is_favorite,omitempty" jsonschema:"True if a task is a favorite task"`
	Labels      []api.TaskLabel `json:"labels,omitempty" jsonschema:"Labels associated with this task, each with a title and optional description"`
	PercentDone *float64        `json:"percent_done,omitempty" jsonschema:"How far the task is from being done, from 0.0 (0%) to 1.0 (100%)"`
	Priority    *int            `json:"priority,omitempty" jsonschema:"The task priority, 0 for none, 1 for low, 5 for critic"`
	ProjectID   int             `json:"project_id,omitempty" jsonschema:"The project this task belongs to"`
	StartDate   string          `json:"start_date,omitempty" jsonschema:"When this task starts"`
	Title       string          `json:"title,omitempty" jsonschema:"The title of the task"`
}

type UpdateTaskOutput struct {
	Task api.Task `json:"task"`
}

func (s *VikunjaServer) UpdateTask(ctx context.Context, req *mcp.CallToolRequest, input UpdateTaskInput) (*mcp.CallToolResult, UpdateTaskOutput, error) {
	current, err := s.client.GetTask(input.TaskID)
	if err != nil {
		result := &mcp.CallToolResult{}
		result.SetError(err)
		return result, UpdateTaskOutput{}, nil
	}

	currentPercentDone := current.PercentDone
	currentPriority := current.Priority
	done := current.Done
	isFavorite := current.IsFavorite
	updateTask := api.TaskUpdate{
		Description: current.Description,
		Done:        &done,
		DueDate:     current.DueDate,
		EndDate:     current.EndDate,
		HexColor:    current.HexColor,
		IsFavorite:  &isFavorite,
		PercentDone: &currentPercentDone,
		Priority:    &currentPriority,
		ProjectID:   current.ProjectID,
		StartDate:   current.StartDate,
		Title:       current.Title,
	}

	if input.Title != "" {
		updateTask.Title = input.Title
	}
	if input.Description != "" {
		updateTask.Description = input.Description
	}
	if input.Done != nil {
		updateTask.Done = input.Done
	}
	if input.DueDate != "" {
		updateTask.DueDate = input.DueDate
	}
	if input.EndDate != "" {
		updateTask.EndDate = input.EndDate
	}
	if input.HexColor != "" {
		updateTask.HexColor = input.HexColor
	}
	if input.IsFavorite != nil {
		updateTask.IsFavorite = input.IsFavorite
	}
	if input.PercentDone != nil {
		updateTask.PercentDone = input.PercentDone
	}
	if input.Priority != nil {
		updateTask.Priority = input.Priority
	}
	if input.ProjectID != 0 {
		updateTask.ProjectID = input.ProjectID
	}
	if input.StartDate != "" {
		updateTask.StartDate = input.StartDate
	}
	if len(input.Labels) > 0 {
		labels := make([]api.TaskLabel, len(input.Labels))
		for i, l := range input.Labels {
			labels[i] = api.TaskLabel{
				Description: l.Description,
				Title:       l.Title,
			}
		}
		updateTask.Labels = labels
	}

	task, err := s.client.UpdateTask(input.TaskID, updateTask)
	if err != nil {
		result := &mcp.CallToolResult{}
		result.SetError(err)
		return result, UpdateTaskOutput{}, nil
	}
	return nil, UpdateTaskOutput{Task: task}, nil
}

type DeleteTaskInput struct {
	TaskID int `json:"task_id" jsonschema:"ID of the task to delete"`
}

type DeleteTaskOutput struct {
	Message string `json:"message"`
}

func (s *VikunjaServer) DeleteTask(ctx context.Context, req *mcp.CallToolRequest, input DeleteTaskInput) (*mcp.CallToolResult, DeleteTaskOutput, error) {
	err := s.client.DeleteTask(input.TaskID)
	if err != nil {
		result := &mcp.CallToolResult{}
		result.SetError(err)
		return result, DeleteTaskOutput{}, nil
	}
	return nil, DeleteTaskOutput{Message: "Task deleted successfully"}, nil
}
