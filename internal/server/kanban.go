package server

import (
	"context"

	"github.com/kowalskyexperto/mcp-vikunja/internal/api"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

type ListProjectViewsInput struct {
	ProjectID int `json:"project_id" jsonschema:"ID of the project"`
}

type ListProjectViewsOutput struct {
	Views api.ProjectViews `json:"views"`
}

func (s *VikunjaServer) ListProjectViews(ctx context.Context, req *mcp.CallToolRequest, input ListProjectViewsInput) (*mcp.CallToolResult, ListProjectViewsOutput, error) {
	views, err := s.client.ListProjectViews(input.ProjectID)
	if err != nil {
		result := &mcp.CallToolResult{}
		result.SetError(err)
		return result, ListProjectViewsOutput{}, nil
	}
	return nil, ListProjectViewsOutput{Views: views}, nil
}

type ListKanbanBucketsInput struct {
	ProjectID int `json:"project_id" jsonschema:"ID of the project"`
	ViewID    int `json:"view_id" jsonschema:"ID of the kanban view (use list_project_views to find it)"`
}

type ListKanbanBucketsOutput struct {
	Buckets api.Buckets `json:"buckets"`
}

func (s *VikunjaServer) ListKanbanBuckets(ctx context.Context, req *mcp.CallToolRequest, input ListKanbanBucketsInput) (*mcp.CallToolResult, ListKanbanBucketsOutput, error) {
	buckets, err := s.client.ListKanbanBuckets(input.ProjectID, input.ViewID)
	if err != nil {
		result := &mcp.CallToolResult{}
		result.SetError(err)
		return result, ListKanbanBucketsOutput{}, nil
	}
	return nil, ListKanbanBucketsOutput{Buckets: buckets}, nil
}

type MoveTaskToBucketInput struct {
	ProjectID int `json:"project_id" jsonschema:"ID of the project"`
	ViewID    int `json:"view_id" jsonschema:"ID of the kanban view"`
	BucketID  int `json:"bucket_id" jsonschema:"ID of the destination bucket (column)"`
	TaskID    int `json:"task_id" jsonschema:"ID of the task to move"`
}

type MoveTaskToBucketOutput struct {
	Message string `json:"message"`
}

func (s *VikunjaServer) MoveTaskToBucket(ctx context.Context, req *mcp.CallToolRequest, input MoveTaskToBucketInput) (*mcp.CallToolResult, MoveTaskToBucketOutput, error) {
	err := s.client.MoveTaskToBucket(input.ProjectID, input.ViewID, input.BucketID, input.TaskID)
	if err != nil {
		result := &mcp.CallToolResult{}
		result.SetError(err)
		return result, MoveTaskToBucketOutput{}, nil
	}
	return nil, MoveTaskToBucketOutput{Message: "Task moved successfully"}, nil
}
