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
