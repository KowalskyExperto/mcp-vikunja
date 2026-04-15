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
