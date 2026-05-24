package main

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/kowalskyexperto/mcp-vikunja/internal/api"
	"github.com/kowalskyexperto/mcp-vikunja/internal/server"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func main() {
	godotenv.Load()
	vikunjaClient, err := api.NewClient(os.Getenv("API_URL"), os.Getenv("API_TOKEN"))
	if err != nil {
		log.Fatalf("Error initializing client: %v", err)
	}
	handler := server.NewVikunjaHandler(vikunjaClient)
	server := mcp.NewServer(&mcp.Implementation{Name: "MCP Vikunja", Version: "v1.0.0"}, nil)
	mcp.AddTool(server, &mcp.Tool{Name: "list_projects", Description: "Retrieve a list of all Vikunja projects, including their titles, descriptions, unique IDs, parent projects, and archived/favorite statuses."}, handler.ListProjects)
	mcp.AddTool(server, &mcp.Tool{Name: "create_project", Description: "Create a new project in Vikunja. Supports specifying title, description, favorite status, parent project ID (to nest), and hex color."}, handler.CreateProject)
	mcp.AddTool(server, &mcp.Tool{Name: "update_project", Description: "Update an existing project's metadata by its ID. Supports title, description, favorite status, parent project ID, and hex color. Senders can use 'clear_fields' to explicitly reset the description."}, handler.UpdateProject)
	mcp.AddTool(server, &mcp.Tool{Name: "delete_project", Description: "Delete a Vikunja project and all its nested tasks permanently by its unique ID. This is a destructive action."}, handler.DeleteProject)
	mcp.AddTool(server, &mcp.Tool{Name: "list_tasks_by_project", Description: "List all tasks belonging to a specific project by its project ID."}, handler.ListTasksByProject)
	mcp.AddTool(server, &mcp.Tool{Name: "search_tasks", Description: "Search across all Vikunja tasks globally by a text keyword or name query."}, handler.ListTasksBySearch)
	mcp.AddTool(server, &mcp.Tool{Name: "create_task", Description: "Create a new task inside a specific Vikunja project. Supports title, description, due date (ISO 8601), and numeric priority (1-5)."}, handler.CreateTask)
	mcp.AddTool(server, &mcp.Tool{Name: "get_task", Description: "Get the complete detailed information of a single task by its task ID, including attached labels, comments, and reminders."}, handler.GetTask)
	mcp.AddTool(server, &mcp.Tool{Name: "list_project_views", Description: "List all views of a project (list, kanban, gantt, etc.). To move a task between kanban columns, first call this tool to get the view ID, then call list_kanban_buckets to get the column IDs, then call move_task_to_bucket."}, handler.ListProjectViews)
	mcp.AddTool(server, &mcp.Tool{Name: "list_kanban_buckets", Description: "List all columns (buckets) within a project's Kanban view. Required before moving tasks to a specific column."}, handler.ListKanbanBuckets)
	mcp.AddTool(server, &mcp.Tool{Name: "create_kanban_bucket", Description: "Create a new column (bucket) inside a project's Kanban view. Supports title, max task capacity limit, and float position."}, handler.CreateKanbanBucket)
	mcp.AddTool(server, &mcp.Tool{Name: "update_kanban_bucket", Description: "Update the title, limit, or position of an existing Kanban column (bucket) by its bucket ID."}, handler.UpdateKanbanBucket)
	mcp.AddTool(server, &mcp.Tool{Name: "delete_kanban_bucket", Description: "Permanently delete an existing Kanban column (bucket) by its bucket ID."}, handler.DeleteKanbanBucket)
	mcp.AddTool(server, &mcp.Tool{Name: "move_task_to_bucket", Description: "Move a specific task to a different Kanban column (bucket) in a project view."}, handler.MoveTaskToBucket)
	mcp.AddTool(server, &mcp.Tool{Name: "list_labels", Description: "List all global labels/tags available to the user, including their titles, descriptions, and hex colors."}, handler.ListLabels)
	mcp.AddTool(server, &mcp.Tool{Name: "get_label", Description: "Retrieve the details of a single global label by its label ID."}, handler.GetLabel)
	mcp.AddTool(server, &mcp.Tool{Name: "create_label", Description: "Create a new global label/tag with a title, description, and custom hex color."}, handler.CreateLabel)
	mcp.AddTool(server, &mcp.Tool{Name: "update_label", Description: "Update the title, description, or color of an existing global label by its label ID."}, handler.UpdateLabel)
	mcp.AddTool(server, &mcp.Tool{Name: "delete_label", Description: "Permanently delete a global label by its ID, removing it from all associated tasks."}, handler.DeleteLabel)
	mcp.AddTool(server, &mcp.Tool{Name: "add_label_to_task", Description: "Attach an existing label to a specific task by their respective IDs."}, handler.AddLabelToTask)
	mcp.AddTool(server, &mcp.Tool{Name: "remove_label_from_task", Description: "Dettach an existing label from a specific task by their respective IDs."}, handler.RemoveLabelFromTask)
	mcp.AddTool(server, &mcp.Tool{Name: "list_task_comments", Description: "List all user comments and notes posted under a specific task by its ID."}, handler.ListTaskComments)
	mcp.AddTool(server, &mcp.Tool{Name: "create_task_comment", Description: "Post a new user comment or note under a specific task."}, handler.CreateTaskComment)
	mcp.AddTool(server, &mcp.Tool{Name: "delete_task_comment", Description: "Delete a specific comment from a task by task ID and comment ID."}, handler.DeleteTaskComment)
	mcp.AddTool(server, &mcp.Tool{Name: "update_task", Description: "Update a task's properties (title, description, done, due/start/end dates, color, priority, favorite status, labels, percent completed) by its ID. Senders can use 'clear_fields' to reset optional properties to blank."}, handler.UpdateTask)
	mcp.AddTool(server, &mcp.Tool{Name: "delete_task", Description: "Permanently delete a specific task by its ID."}, handler.DeleteTask)
	mcp.AddTool(server, &mcp.Tool{Name: "create_task_relation", Description: "Create a relationship/dependency between two tasks. Supported kinds: subtask, parenttask, related, duplicateof, duplicates, blocking, blocked, precedes, follows."}, handler.CreateTaskRelation)
	mcp.AddTool(server, &mcp.Tool{Name: "delete_task_relation", Description: "Remove an existing relationship/dependency between two tasks."}, handler.DeleteTaskRelation)
	server.Run(context.Background(), &mcp.StdioTransport{})
}
