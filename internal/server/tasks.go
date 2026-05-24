package server

import (
	"context"

	"github.com/kowalskyexperto/mcp-vikunja/internal/api"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

type ListTasksBySearchInput struct {
	Search string `json:"search" jsonschema:"The text query to search tasks by name/title"`
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
	ProjectID   int    `json:"project_id" jsonschema:"The unique ID of the Vikunja project to create the task in"`
	Description string `json:"description,omitempty" jsonschema:"A detailed description of the task"`
	DueDate     string `json:"due_date,omitempty" jsonschema:"The task due date in ISO 8601 format (YYYY-MM-DDTHH:MM:SSZ)"`
	Priority    int    `json:"priority,omitempty" jsonschema:"The task priority level (1 for low, 2 for medium, 3 for high, 4 for urgent, 5 for critical)"`
	Title       string `json:"title" jsonschema:"The title or summary of the task"`
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
	TaskID int `json:"task_id" jsonschema:"The unique ID of the Vikunja task to retrieve"`
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
	TaskID int `json:"id" jsonschema:"The unique ID of the Vikunja task to update"`

	ClearFields []string        `json:"clear_fields,omitempty" jsonschema:"List of optional fields to clear/reset. Supported: description, due_date, start_date, end_date, hex_color"`
	Description string          `json:"description,omitempty" jsonschema:"The new detailed description for the task"`
	Done        *bool           `json:"done,omitempty" jsonschema:"Whether the task is marked as completed/done"`
	DueDate     string          `json:"due_date,omitempty" jsonschema:"The new due date for the task in ISO 8601 format (YYYY-MM-DDTHH:MM:SSZ)"`
	EndDate     string          `json:"end_date,omitempty" jsonschema:"The new end date for the task in ISO 8601 format (YYYY-MM-DDTHH:MM:SSZ)"`
	HexColor    string          `json:"hex_color,omitempty" jsonschema:"The task color in hexadecimal format (e.g. '#ff0000', max 7 characters)"`
	IsFavorite  *bool           `json:"is_favorite,omitempty" jsonschema:"Whether the task is marked as a favorite"`
	Labels      []api.TaskLabel `json:"labels,omitempty" jsonschema:"New labels to associate with this task, each with a title and optional description"`
	PercentDone *float64        `json:"percent_done,omitempty" jsonschema:"Completion percentage of the task, from 0.0 (0%) to 1.0 (100%)"`
	Priority    *int            `json:"priority,omitempty" jsonschema:"The new task priority level (0 for none, 1 for low, 2 for medium, 3 for high, 4 for urgent, 5 for critical)"`
	ProjectID   int             `json:"project_id,omitempty" jsonschema:"The ID of the project this task should belong to (useful to move the task to a different project)"`
	StartDate   string          `json:"start_date,omitempty" jsonschema:"The new start date for the task in ISO 8601 format (YYYY-MM-DDTHH:MM:SSZ)"`
	Title       string          `json:"title,omitempty" jsonschema:"The new title or summary for the task"`
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

	clearSet := make(map[string]bool)
	for _, f := range input.ClearFields {
		clearSet[f] = true
	}

	if clearSet["description"] {
		updateTask.Description = ""
	} else if input.Description != "" {
		updateTask.Description = input.Description
	}
	if input.Title != "" {
		updateTask.Title = input.Title
	}
	if input.Done != nil {
		updateTask.Done = input.Done
	}
	if clearSet["due_date"] {
		updateTask.DueDate = ""
	} else if input.DueDate != "" {
		updateTask.DueDate = input.DueDate
	}
	if clearSet["end_date"] {
		updateTask.EndDate = ""
	} else if input.EndDate != "" {
		updateTask.EndDate = input.EndDate
	}
	if clearSet["hex_color"] {
		updateTask.HexColor = ""
	} else if input.HexColor != "" {
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
	if clearSet["start_date"] {
		updateTask.StartDate = ""
	} else if input.StartDate != "" {
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
	TaskID int `json:"task_id" jsonschema:"The unique ID of the Vikunja task to delete"`
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
