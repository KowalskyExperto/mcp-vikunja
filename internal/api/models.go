package api

type Project struct {
	Created         string `json:"created"`
	Description     string `json:"description"`
	ID              int    `json:"id"`
	Identifier      string `json:"identifier"`
	IsArchived      bool   `json:"is_archived"`
	IsFavorite      bool   `json:"is_favorite"`
	ParentProjectID int    `json:"parent_project_id"`
	Title           string `json:"title"`
	Updated         string `json:"updated"`
}

type Projects []Project

type Task struct {
	CommentCount int    `json:"comment_count"`
	Created      string `json:"created"`
	Description  string `json:"description"`
	Done         bool   `json:"done"`
	DueDate      string `json:"due_date"`
	EndDate      string `json:"end_date"`
	ID           int    `json:"id"`
	Identifier   string `json:"identifier"`
	IsFavorite   bool   `json:"is_favorite"`
	ProjectID    int    `json:"project_id"`
	StartDate    string `json:"start_date"`
	Title        string `json:"title"`
	Updated      string `json:"updated"`
}

type Tasks []Task
