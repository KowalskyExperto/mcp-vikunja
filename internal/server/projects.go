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

type ListProjectsInput struct{}

type ListProjectsOutput struct {
	Projects api.Projects `json:"projects"`
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

type CreateProjectInput struct {
	Title           string `json:"title" jsonschema:"The title of the project"`
	Description     string `json:"description,omitempty" jsonschema:"Project description"`
	HexColor        string `json:"hex_color,omitempty" jsonschema:"Project color in hex format, e.g. #ff0000"`
	IsFavorite      bool   `json:"is_favorite,omitempty" jsonschema:"Mark this project as a favorite"`
	ParentProjectID *int   `json:"parent_project_id,omitempty" jsonschema:"ID of the parent project to nest this project under"`
}

type CreateProjectOutput struct {
	Project api.Project `json:"project"`
}

func (s *VikunjaServer) CreateProject(ctx context.Context, req *mcp.CallToolRequest, input CreateProjectInput) (*mcp.CallToolResult, CreateProjectOutput, error) {
	newProject := api.ProjectInput{
		Title:           input.Title,
		Description:     input.Description,
		HexColor:        input.HexColor,
		IsFavorite:      input.IsFavorite,
		ParentProjectID: input.ParentProjectID,
	}
	project, err := s.client.CreateProject(newProject)
	if err != nil {
		result := &mcp.CallToolResult{}
		result.SetError(err)
		return result, CreateProjectOutput{}, nil
	}
	return nil, CreateProjectOutput{Project: project}, nil
}

type UpdateProjectInput struct {
	ProjectID       int    `json:"project_id" jsonschema:"ID of the project to update"`
	Title           string `json:"title,omitempty" jsonschema:"New title for the project"`
	Description     string `json:"description,omitempty" jsonschema:"New description for the project"`
	HexColor        string `json:"hex_color,omitempty" jsonschema:"Project color in hex format, e.g. #ff0000"`
	IsFavorite      bool   `json:"is_favorite,omitempty" jsonschema:"Mark this project as a favorite"`
	ParentProjectID *int   `json:"parent_project_id,omitempty" jsonschema:"Set the parent project ID, or 0 to remove the parent"`
}

type UpdateProjectOutput struct {
	Project api.Project `json:"project"`
}

func (s *VikunjaServer) UpdateProject(ctx context.Context, req *mcp.CallToolRequest, input UpdateProjectInput) (*mcp.CallToolResult, UpdateProjectOutput, error) {
	current, err := s.client.GetProject(input.ProjectID)
	if err != nil {
		result := &mcp.CallToolResult{}
		result.SetError(err)
		return result, UpdateProjectOutput{}, nil
	}

	currentParent := current.ParentProjectID
	updated := api.ProjectInput{
		Title:           current.Title,
		Description:     current.Description,
		IsFavorite:      current.IsFavorite,
		ParentProjectID: &currentParent,
	}
	if input.Title != "" {
		updated.Title = input.Title
	}
	if input.Description != "" {
		updated.Description = input.Description
	}
	if input.HexColor != "" {
		updated.HexColor = input.HexColor
	}
	if input.IsFavorite {
		updated.IsFavorite = input.IsFavorite
	}
	if input.ParentProjectID != nil {
		updated.ParentProjectID = input.ParentProjectID
	}

	project, err := s.client.UpdateProject(input.ProjectID, updated)
	if err != nil {
		result := &mcp.CallToolResult{}
		result.SetError(err)
		return result, UpdateProjectOutput{}, nil
	}
	return nil, UpdateProjectOutput{Project: project}, nil
}

type DeleteProjectInput struct {
	ProjectID int `json:"project_id" jsonschema:"ID of the project to delete"`
}

type DeleteProjectOutput struct {
	Message string `json:"message"`
}

func (s *VikunjaServer) DeleteProject(ctx context.Context, req *mcp.CallToolRequest, input DeleteProjectInput) (*mcp.CallToolResult, DeleteProjectOutput, error) {
	err := s.client.DeleteProject(input.ProjectID)
	if err != nil {
		result := &mcp.CallToolResult{}
		result.SetError(err)
		return result, DeleteProjectOutput{}, nil
	}
	return nil, DeleteProjectOutput{Message: "Project deleted successfully"}, nil
}
