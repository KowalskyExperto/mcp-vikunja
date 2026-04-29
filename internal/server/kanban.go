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

type CreateKanbanBucketInput struct {
	ProjectID int     `json:"project_id" jsonschema:"ID of the project"`
	ViewID    int     `json:"view_id" jsonschema:"ID of the kanban view"`
	Title     string  `json:"title" jsonschema:"Title of the new bucket"`
	Limit     int     `json:"limit,omitempty" jsonschema:"Max number of tasks in this bucket"`
	Position  float64 `json:"position,omitempty" jsonschema:"Position of the bucket"`
}

type CreateKanbanBucketOutput struct {
	Bucket api.Bucket `json:"bucket"`
}

func (s *VikunjaServer) CreateKanbanBucket(ctx context.Context, req *mcp.CallToolRequest, input CreateKanbanBucketInput) (*mcp.CallToolResult, CreateKanbanBucketOutput, error) {
	bucket, err := s.client.CreateKanbanBucket(input.ProjectID, input.ViewID, api.BucketInput{
		Title:    input.Title,
		Limit:    input.Limit,
		Position: input.Position,
	})
	if err != nil {
		result := &mcp.CallToolResult{}
		result.SetError(err)
		return result, CreateKanbanBucketOutput{}, nil
	}
	return nil, CreateKanbanBucketOutput{Bucket: bucket}, nil
}

type UpdateKanbanBucketInput struct {
	ProjectID int     `json:"project_id" jsonschema:"ID of the project"`
	ViewID    int     `json:"view_id" jsonschema:"ID of the kanban view"`
	BucketID  int     `json:"bucket_id" jsonschema:"ID of the bucket to update"`
	Title     string  `json:"title,omitempty" jsonschema:"New title for the bucket"`
	Limit     int     `json:"limit,omitempty" jsonschema:"New limit for the bucket"`
	Position  float64 `json:"position,omitempty" jsonschema:"New position for the bucket"`
}

type UpdateKanbanBucketOutput struct {
	Bucket api.Bucket `json:"bucket"`
}

func (s *VikunjaServer) UpdateKanbanBucket(ctx context.Context, req *mcp.CallToolRequest, input UpdateKanbanBucketInput) (*mcp.CallToolResult, UpdateKanbanBucketOutput, error) {
	updated := api.BucketInput{
		Title:    input.Title,
		Limit:    input.Limit,
		Position: input.Position,
	}

	bucket, err := s.client.UpdateKanbanBucket(input.ProjectID, input.ViewID, input.BucketID, updated)
	if err != nil {
		result := &mcp.CallToolResult{}
		result.SetError(err)
		return result, UpdateKanbanBucketOutput{}, nil
	}
	return nil, UpdateKanbanBucketOutput{Bucket: bucket}, nil
}

type DeleteKanbanBucketInput struct {
	ProjectID int `json:"project_id" jsonschema:"ID of the project"`
	ViewID    int `json:"view_id" jsonschema:"ID of the kanban view"`
	BucketID  int `json:"bucket_id" jsonschema:"ID of the bucket to delete"`
}

type DeleteKanbanBucketOutput struct {
	Message string `json:"message"`
}

func (s *VikunjaServer) DeleteKanbanBucket(ctx context.Context, req *mcp.CallToolRequest, input DeleteKanbanBucketInput) (*mcp.CallToolResult, DeleteKanbanBucketOutput, error) {
	err := s.client.DeleteKanbanBucket(input.ProjectID, input.ViewID, input.BucketID)
	if err != nil {
		result := &mcp.CallToolResult{}
		result.SetError(err)
		return result, DeleteKanbanBucketOutput{}, nil
	}
	return nil, DeleteKanbanBucketOutput{Message: "Bucket deleted successfully"}, nil
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
