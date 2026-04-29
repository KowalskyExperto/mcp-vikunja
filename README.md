# mcp-vikunja

An MCP server that connects Claude to a self-hosted [Vikunja](https://vikunja.io) instance, allowing you to manage projects and tasks directly from Claude or Claude Code.

## Setup

1. Copy `.env.example` to `.env` and fill in your values:

```env
API_URL=https://your-vikunja-instance.com/api/v1/
API_TOKEN=your_api_token_here
```

2. Run the server:

```bash
go run .
```

3. To test with the MCP Inspector:

```bash
npx @modelcontextprotocol/inspector go run .
```

## Available Tools

### Projects

| Tool | Description |
|------|-------------|
| `list_projects` | List all projects |
| `create_project` | Create a new project |
| `update_project` | Update a project by ID. Use `clear_fields` to clear optional fields (e.g. `["description"]`) |
| `delete_project` | Delete a project by ID |
| `list_tasks_by_project` | List all tasks in a project |

### Tasks

| Tool | Description |
|------|-------------|
| `search_tasks` | Search tasks by name |
| `create_task` | Create a task in a project |
| `get_task` | Get full task details by ID |
| `update_task` | Update a task by ID. Use `clear_fields` to clear optional fields (e.g. `["description", "due_date"]`) |
| `delete_task` | Delete a task by ID |

### Comments

| Tool | Description |
|------|-------------|
| `list_task_comments` | List all comments on a task |
| `create_task_comment` | Add a comment to a task |
| `delete_task_comment` | Delete a comment from a task |

### Labels

| Tool | Description |
|------|-------------|
| `list_labels` | List all available labels |
| `get_label` | Get a label by ID |
| `create_label` | Create a new label |
| `update_label` | Update an existing label by ID |
| `delete_label` | Delete a label by ID |
| `add_label_to_task` | Add a label to a task |
| `remove_label_from_task` | Remove a label from a task |

### Kanban

| Tool | Description |
|------|-------------|
| `list_project_views` | List all views of a project (list, kanban, gantt, etc.) |
| `list_kanban_buckets` | List all columns of a kanban view |
| `create_kanban_bucket` | Create a new kanban bucket on a project view |
| `update_kanban_bucket` | Update an existing kanban bucket |
| `delete_kanban_bucket` | Delete an existing kanban bucket |
| `move_task_to_bucket` | Move a task to a different kanban column |


To manage kanban columns: `list_project_views` → `list_kanban_buckets` → `create_kanban_bucket` / `update_kanban_bucket` / `delete_kanban_bucket`.
To move a task between columns: `list_project_views` → `list_kanban_buckets` → `move_task_to_bucket`.

### Task Relations

| Tool | Description |
|------|-------------|
| `create_task_relation` | Create a relation between two tasks |
| `delete_task_relation` | Remove a relation between two tasks |

Available relation kinds: `subtask`, `parenttask`, `related`, `duplicateof`, `duplicates`, `blocking`, `blocked`, `precedes`, `follows`.

## Features & Improvements

- **Rich Error Reporting:** The server captures and returns detailed error messages from the Vikunja API, including validation failures, to simplify debugging.
- **Full Resource Management:** Support for complete CRUD operations on Projects, Tasks, Labels, and Kanban Buckets.
- **Smart Updates:** `update_task` and `update_project` support a `clear_fields` parameter to explicitly reset optional fields to empty values.
