# mcp-vikunja

An MCP (Model Context Protocol) server that connects Claude and other MCP clients to a self-hosted [Vikunja](https://vikunja.io) instance. This allows AI assistants like Claude Desktop, Cursor, or Claude Code to manage your projects, tasks, labels, comments, and Kanban boards directly through natural language.

---

## Features

- **Full Project Management:** Create, list, update, and delete projects, with support for parent-child project nesting.
- **Robust Task CRUD:** Search, create, read, update, and delete tasks. Schedulable properties like due dates, priorities, and description fields are fully supported.
- **Smart Updates:** Explicitly reset optional fields (like `description`, `due_date`, etc.) to empty values using the special `clear_fields` argument.
- **Kanban Support:** Manage Kanban columns (buckets) within project views and easily move tasks between columns.
- **Labels & Tags:** Categorize tasks by creating, deleting, and attaching/removing labels.
- **Task Relations:** Define dependencies between tasks (e.g., subtask, parenttask, blocking, blocked, related, duplicates, etc.).
- **Comments Management:** Retrieve, add, and delete comments on individual tasks.
- **Rich Error Reporting:** Detailed error messages from the Vikunja API (including payload validation issues) are piped back to the LLM to make debugging and error recovery seamless.
- **Embedded API Reference:** The official Swagger 2.0 Vikunja API specification is included locally in [docs/swagger.json](file:///mnt/kanji/Development/Projects/mcp-vikunja/docs/swagger.json), serving as an offline reference for developers and a rich context source for AI agents.

---

## Requirements

Before setting up `mcp-vikunja`, ensure you have the following installed:

1. **Go (Golang):** Version **1.26** or higher. You can verify your version with:
   ```bash
   go version
   ```
2. **Vikunja Instance:** A running self-hosted Vikunja instance (API version 1).
3. **Vikunja API Token:** A Personal API Token generated from your Vikunja user settings.
   - Go to your Vikunja web interface.
   - Click on your avatar / **Settings** -> **API Tokens**.
   - Create a new token with appropriate read/write scopes for **Projects**, **Tasks**, **Labels**, etc.
4. **Node.js (Optional):** Required only if you want to use the MCP Inspector for interactive testing.

---

## Installation

### 1. Clone the Repository
Clone the repository to your local machine and navigate into the project directory:
```bash
git clone https://github.com/kowalskyexperto/mcp-vikunja.git
cd mcp-vikunja
```

### 2. Build the Server
Build the production-ready binary using Go:
```bash
go build -o mcp-vikunja .
```
This compiles the code into an executable called `mcp-vikunja` in the root of the project.

---

## Configuration

The server expects two environment variables for its configuration:

- `API_URL`: The URL of your self-hosted Vikunja API v1 endpoint (must end with a trailing slash, e.g., `https://your-vikunja-instance.com/api/v1/`).
- `API_TOKEN`: Your personal Vikunja API token (generated in your user settings).

> [!NOTE]
> Since this is a Model Context Protocol server, you do not need to create or maintain a `.env` file. The environment variables are injected directly by your MCP client (such as Claude Desktop or Cursor) during server startup.

---

## MCP Integration Setup

To use this server with your favorite AI assistant, configure it to launch the compiled binary or run the project through Go.

### Claude Desktop
Add the following server configuration to your `claude_desktop_config.json`:

* **On Linux:** `~/.config/Claude/claude_desktop_config.json`
* **On macOS:** `~/Library/Application Support/Claude/claude_desktop_config.json`
* **On Windows:** `%APPDATA%\Claude\claude_desktop_config.json`

#### Option A: Running the precompiled binary (Recommended)
```json
{
  "mcpServers": {
    "vikunja": {
      "command": "/absolute/path/to/mcp-vikunja/mcp-vikunja",
      "env": {
        "API_URL": "https://your-vikunja-instance.com/api/v1/",
        "API_TOKEN": "your_personal_api_token_here"
      }
    }
  }
}
```

#### Option B: Running via Go source code
```json
{
  "mcpServers": {
    "vikunja": {
      "command": "go",
      "args": [
        "run",
        "/absolute/path/to/mcp-vikunja/main.go"
      ],
      "env": {
        "API_URL": "https://your-vikunja-instance.com/api/v1/",
        "API_TOKEN": "your_personal_api_token_here"
      }
    }
  }
}
```

---

## Development and Testing

### Run the Server Locally
To start the server locally in stdio mode:
```bash
go run .
```

### Test with MCP Inspector
You can test the available tools and view JSON schemas interactively using the official Model Context Protocol Inspector:
```bash
npx @modelcontextprotocol/inspector go run .
```
This command spins up a web-based testing utility where you can send tool requests and inspect raw JSON payloads.

---

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
| `update_task` | Update a task by ID. Use `clear_fields` to clear optional fields (e.g. `["description", "due_date", "start_date", "end_date", "hex_color"]`) |
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

* **Workflow to manage Kanban columns:** `list_project_views` &rarr; `list_kanban_buckets` &rarr; `create_kanban_bucket` / `update_kanban_bucket` / `delete_kanban_bucket`.
* **Workflow to move a task:** `list_project_views` &rarr; `list_kanban_buckets` &rarr; `move_task_to_bucket`.

### Task Relations

| Tool | Description |
|------|-------------|
| `create_task_relation` | Create a relation between two tasks |
| `delete_task_relation` | Remove a relation between two tasks |

* **Supported relation kinds:** `subtask`, `parenttask`, `related`, `duplicateof`, `duplicates`, `blocking`, `blocked`, `precedes`, `follows`.

