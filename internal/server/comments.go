package server

import (
	"context"

	"github.com/kowalskyexperto/mcp-vikunja/internal/api"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

type ListTaskCommentsInput struct {
	TaskID int `json:"task_id" jsonschema:"ID of the task to list comments from"`
}

type ListTaskCommentsOutput struct {
	Comments api.TaskComments `json:"comments"`
}

func (s *VikunjaServer) ListTaskComments(ctx context.Context, req *mcp.CallToolRequest, input ListTaskCommentsInput) (*mcp.CallToolResult, ListTaskCommentsOutput, error) {
	comments, err := s.client.ListTaskComments(input.TaskID)
	if err != nil {
		result := &mcp.CallToolResult{}
		result.SetError(err)
		return result, ListTaskCommentsOutput{}, nil
	}
	return nil, ListTaskCommentsOutput{Comments: comments}, nil
}

type CreateTaskCommentInput struct {
	TaskID  int    `json:"task_id" jsonschema:"ID of the task to comment on"`
	Comment string `json:"comment" jsonschema:"The comment text"`
}

type CreateTaskCommentOutput struct {
	Comment api.TaskComment `json:"comment"`
}

func (s *VikunjaServer) CreateTaskComment(ctx context.Context, req *mcp.CallToolRequest, input CreateTaskCommentInput) (*mcp.CallToolResult, CreateTaskCommentOutput, error) {
	comment, err := s.client.CreateTaskComment(input.TaskID, api.TaskCommentInput{Comment: input.Comment})
	if err != nil {
		result := &mcp.CallToolResult{}
		result.SetError(err)
		return result, CreateTaskCommentOutput{}, nil
	}
	return nil, CreateTaskCommentOutput{Comment: comment}, nil
}

type DeleteTaskCommentInput struct {
	TaskID    int `json:"task_id" jsonschema:"ID of the task that owns the comment"`
	CommentID int `json:"comment_id" jsonschema:"ID of the comment to delete"`
}

type DeleteTaskCommentOutput struct {
	Message string `json:"message"`
}

func (s *VikunjaServer) DeleteTaskComment(ctx context.Context, req *mcp.CallToolRequest, input DeleteTaskCommentInput) (*mcp.CallToolResult, DeleteTaskCommentOutput, error) {
	err := s.client.DeleteTaskComment(input.TaskID, input.CommentID)
	if err != nil {
		result := &mcp.CallToolResult{}
		result.SetError(err)
		return result, DeleteTaskCommentOutput{}, nil
	}
	return nil, DeleteTaskCommentOutput{Message: "Comment deleted successfully"}, nil
}
