# MCP Vikunja — Tools Checklist

Tools already implemented are marked as done. The rest are pending.

## Existing Tools

- [x] `list_projects` — List all projects
- [x] `list_tasks_by_project` — List tasks in a project
- [x] `search_tasks` — Search tasks by name
- [x] `create_task` — Create a task in a project
- [x] `get_task` — Get task details by ID
- [x] `update_task` — Update a task by ID
- [x] `delete_task` — Delete a task by ID

---

## Projects

- [x] `create_project` — `PUT /projects` — Create a new project
- [x] `update_project` — `POST /projects/{id}` — Update project title, description, or archive it
- [x] `delete_project` — `DELETE /projects/{id}` — Delete a project

## Task Comments

- [ ] `list_task_comments` — `GET /tasks/{id}/comments` — List all comments on a task
- [ ] `create_task_comment` — `PUT /tasks/{id}/comments` — Add a comment to a task
- [ ] `delete_task_comment` — `DELETE /tasks/{id}/comments/{commentID}` — Remove a comment

## Labels

- [ ] `list_labels` — `GET /labels` — List all labels available to the user
- [ ] `create_label` — `PUT /labels` — Create a new global label
- [ ] `delete_label` — `DELETE /labels/{id}` — Delete a label
- [ ] `add_label_to_task` — `PUT /tasks/{id}/labels` — Assign a label to a task
- [ ] `remove_label_from_task` — `DELETE /tasks/{id}/labels/{label}` — Remove a label from a task

## Kanban

- [ ] `list_project_views` — `GET /projects/{id}/views` — List views of a project (list, kanban, gantt, etc.)
- [ ] `list_kanban_buckets` — `GET /projects/{id}/views/{view}/buckets` — List kanban columns (buckets)
- [ ] `move_task_to_bucket` — `POST /projects/{project}/views/{view}/buckets/{bucket}/tasks` — Move a task to a kanban column

## Users and Notifications

- [ ] `get_current_user` — `GET /user` — Get the current authenticated user
- [ ] `search_users` — `GET /users?s=` — Search users by username (needed for assignees)
- [ ] `list_task_assignees` — `GET /tasks/{id}/assignees` — List assignees of a task
- [ ] `add_task_assignee` — `PUT /tasks/{id}/assignees` — Assign a user to a task
- [ ] `remove_task_assignee` — `DELETE /tasks/{id}/assignees/{userID}` — Remove an assignee from a task
- [ ] `list_notifications` — `GET /notifications` — List unread notifications for the current user
- [ ] `mark_notifications_read` — `POST /notifications` — Mark all notifications as read

## Task Relations

- [ ] `create_task_relation` — `PUT /tasks/{id}/relations` — Create a relation between two tasks (subtask, blocked by, duplicate, etc.)
- [ ] `delete_task_relation` — `DELETE /tasks/{id}/relations/{kind}/{otherTaskID}` — Remove a relation between two tasks
