package server

import (
	"context"

	"github.com/kowalskyexperto/mcp-vikunja/internal/api"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

type ListLabelsInput struct{}

type ListLabelsOutput struct {
	Labels api.Labels `json:"labels"`
}

func (s *VikunjaServer) ListLabels(ctx context.Context, req *mcp.CallToolRequest, input ListLabelsInput) (*mcp.CallToolResult, ListLabelsOutput, error) {
	labels, err := s.client.ListLabels()
	if err != nil {
		result := &mcp.CallToolResult{}
		result.SetError(err)
		return result, ListLabelsOutput{}, nil
	}
	return nil, ListLabelsOutput{Labels: labels}, nil
}

type CreateLabelInput struct {
	Title       string `json:"title" jsonschema:"The title of the label"`
	Description string `json:"description,omitempty" jsonschema:"Label description"`
	HexColor    string `json:"hex_color,omitempty" jsonschema:"Label color in hex format, e.g. #ff0000"`
}

type CreateLabelOutput struct {
	Label api.Label `json:"label"`
}

func (s *VikunjaServer) CreateLabel(ctx context.Context, req *mcp.CallToolRequest, input CreateLabelInput) (*mcp.CallToolResult, CreateLabelOutput, error) {
	label, err := s.client.CreateLabel(api.LabelInput{
		Title:       input.Title,
		Description: input.Description,
		HexColor:    input.HexColor,
	})
	if err != nil {
		result := &mcp.CallToolResult{}
		result.SetError(err)
		return result, CreateLabelOutput{}, nil
	}
	return nil, CreateLabelOutput{Label: label}, nil
}

type DeleteLabelInput struct {
	LabelID int `json:"label_id" jsonschema:"ID of the label to delete"`
}

type DeleteLabelOutput struct {
	Message string `json:"message"`
}

func (s *VikunjaServer) DeleteLabel(ctx context.Context, req *mcp.CallToolRequest, input DeleteLabelInput) (*mcp.CallToolResult, DeleteLabelOutput, error) {
	err := s.client.DeleteLabel(input.LabelID)
	if err != nil {
		result := &mcp.CallToolResult{}
		result.SetError(err)
		return result, DeleteLabelOutput{}, nil
	}
	return nil, DeleteLabelOutput{Message: "Label deleted successfully"}, nil
}

type AddLabelToTaskInput struct {
	TaskID  int `json:"task_id" jsonschema:"ID of the task"`
	LabelID int `json:"label_id" jsonschema:"ID of the label to add"`
}

type AddLabelToTaskOutput struct {
	Label api.Label `json:"label"`
}

func (s *VikunjaServer) AddLabelToTask(ctx context.Context, req *mcp.CallToolRequest, input AddLabelToTaskInput) (*mcp.CallToolResult, AddLabelToTaskOutput, error) {
	label, err := s.client.AddLabelToTask(input.TaskID, input.LabelID)
	if err != nil {
		result := &mcp.CallToolResult{}
		result.SetError(err)
		return result, AddLabelToTaskOutput{}, nil
	}
	return nil, AddLabelToTaskOutput{Label: label}, nil
}

type RemoveLabelFromTaskInput struct {
	TaskID  int `json:"task_id" jsonschema:"ID of the task"`
	LabelID int `json:"label_id" jsonschema:"ID of the label to remove"`
}

type RemoveLabelFromTaskOutput struct {
	Message string `json:"message"`
}

func (s *VikunjaServer) RemoveLabelFromTask(ctx context.Context, req *mcp.CallToolRequest, input RemoveLabelFromTaskInput) (*mcp.CallToolResult, RemoveLabelFromTaskOutput, error) {
	err := s.client.RemoveLabelFromTask(input.TaskID, input.LabelID)
	if err != nil {
		result := &mcp.CallToolResult{}
		result.SetError(err)
		return result, RemoveLabelFromTaskOutput{}, nil
	}
	return nil, RemoveLabelFromTaskOutput{Message: "Label removed from task successfully"}, nil
}
