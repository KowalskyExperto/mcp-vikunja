package server

import (
	"context"

	"github.com/kowalskyexperto/mcp-vikunja/internal/api"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

type CreateTaskRelationInput struct {
	TaskID       int    `json:"task_id" jsonschema:"ID of the source task"`
	OtherTaskID  int    `json:"other_task_id" jsonschema:"ID of the task to relate to"`
	RelationKind string `json:"relation_kind" jsonschema:"Type of relation: subtask, parenttask, related, duplicateof, duplicates, blocking, blocked, precedes, follows"`
}

type CreateTaskRelationOutput struct {
	Message string `json:"message"`
}

func (s *VikunjaServer) CreateTaskRelation(ctx context.Context, req *mcp.CallToolRequest, input CreateTaskRelationInput) (*mcp.CallToolResult, CreateTaskRelationOutput, error) {
	err := s.client.CreateTaskRelation(input.TaskID, api.TaskRelationInput{
		OtherTaskID:  input.OtherTaskID,
		RelationKind: input.RelationKind,
	})
	if err != nil {
		result := &mcp.CallToolResult{}
		result.SetError(err)
		return result, CreateTaskRelationOutput{}, nil
	}
	return nil, CreateTaskRelationOutput{Message: "Task relation created successfully"}, nil
}

type DeleteTaskRelationInput struct {
	TaskID       int    `json:"task_id" jsonschema:"ID of the source task"`
	OtherTaskID  int    `json:"other_task_id" jsonschema:"ID of the related task"`
	RelationKind string `json:"relation_kind" jsonschema:"Type of relation to remove: subtask, parenttask, related, duplicateof, duplicates, blocking, blocked, precedes, follows"`
}

type DeleteTaskRelationOutput struct {
	Message string `json:"message"`
}

func (s *VikunjaServer) DeleteTaskRelation(ctx context.Context, req *mcp.CallToolRequest, input DeleteTaskRelationInput) (*mcp.CallToolResult, DeleteTaskRelationOutput, error) {
	err := s.client.DeleteTaskRelation(input.TaskID, input.RelationKind, input.OtherTaskID)
	if err != nil {
		result := &mcp.CallToolResult{}
		result.SetError(err)
		return result, DeleteTaskRelationOutput{}, nil
	}
	return nil, DeleteTaskRelationOutput{Message: "Task relation deleted successfully"}, nil
}
