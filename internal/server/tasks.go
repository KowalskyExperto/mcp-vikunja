package server

import (
	"context"

	"github.com/kowalskyexperto/mcp-vikunja/internal/api"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

type ListTasksByProjectInput struct {
	ProjectID int `json:"project_id" jsonschema:"El ID del proyecto de Vikunja"`
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
