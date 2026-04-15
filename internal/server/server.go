package server

import (
	"github.com/kowalskyexperto/mcp-vikunja/internal/api"
)

type VikunjaServer struct {
	client *api.Client
}

func NewVikunjaHandler(client *api.Client) *VikunjaServer {
	return &VikunjaServer{
		client: client,
	}
}
