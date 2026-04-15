package server

import (
	"context"

	"github.com/kowalskyexperto/mcp-vikunja/internal/api"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

type VikunjaServer struct {
	client *api.Client
}

type ListProjectsInput struct {
}

type ListProjectsOutput struct {
	Projects api.Projects `json:"projects"`
}

func NewVikunjaHandler(client *api.Client) *VikunjaServer {
	return &VikunjaServer{
		client: client,
	}
}

func (s *VikunjaServer) ListProjects(ctx context.Context, req *mcp.CallToolRequest, input ListProjectsInput) (*mcp.CallToolResult, ListProjectsOutput, error) {
	projects, err := s.client.GetProjects()
	if err != nil {
		result := &mcp.CallToolResult{}
		result.SetError(err)
		return result, ListProjectsOutput{}, nil
	}
	return nil, ListProjectsOutput{Projects: projects}, nil
}
