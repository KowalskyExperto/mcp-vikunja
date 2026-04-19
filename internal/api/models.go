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

type ProjectInput struct {
	Description     string `json:"description,omitempty"`
	HexColor        string `json:"hex_color,omitempty"`
	IsFavorite      bool   `json:"is_favorite,omitempty"`
	ParentProjectID int    `json:"parent_project_id,omitempty"`
	Title           string `json:"title"`
}

type Tasks []Task

type TaskInput struct {
	Description string `json:"description,omitempty"`
	DueDate     string `json:"due_date,omitempty"`
	Priority    int    `json:"priority,omitempty"`
	Title       string `json:"title"`
}

type TaskDetail struct {
	Attachments []struct {
		Created   string `json:"created"`
		CreatedBy struct {
			Created  string `json:"created"`
			Email    string `json:"email"`
			ID       int    `json:"id"`
			Name     string `json:"name"`
			Updated  string `json:"updated"`
			Username string `json:"username"`
		} `json:"created_by"`
		File struct {
			Created string `json:"created"`
			ID      int    `json:"id"`
			Mime    string `json:"mime"`
			Name    string `json:"name"`
			Size    int    `json:"size"`
		} `json:"file"`
		ID     int `json:"id"`
		TaskID int `json:"task_id"`
	} `json:"attachments"`
	BucketID int `json:"bucket_id"`
	Buckets  []struct {
		Count     int    `json:"count"`
		Created   string `json:"created"`
		CreatedBy struct {
			Created  string `json:"created"`
			Email    string `json:"email"`
			ID       int    `json:"id"`
			Name     string `json:"name"`
			Updated  string `json:"updated"`
			Username string `json:"username"`
		} `json:"created_by"`
		ID            int `json:"id"`
		Limit         int `json:"limit"`
		Position      int `json:"position"`
		ProjectViewID int `json:"project_view_id"`
		Tasks         []struct {
		} `json:"tasks"`
		Title   string `json:"title"`
		Updated string `json:"updated"`
	} `json:"buckets"`
	CommentCount int `json:"comment_count"`
	Comments     []struct {
		Author struct {
			Created  string `json:"created"`
			Email    string `json:"email"`
			ID       int    `json:"id"`
			Name     string `json:"name"`
			Updated  string `json:"updated"`
			Username string `json:"username"`
		} `json:"author"`
		Comment   string `json:"comment"`
		Created   string `json:"created"`
		ID        int    `json:"id"`
		Reactions struct {
			Property1 []struct {
				Created  string `json:"created"`
				Email    string `json:"email"`
				ID       int    `json:"id"`
				Name     string `json:"name"`
				Updated  string `json:"updated"`
				Username string `json:"username"`
			} `json:"property1"`
			Property2 []struct {
				Created  string `json:"created"`
				Email    string `json:"email"`
				ID       int    `json:"id"`
				Name     string `json:"name"`
				Updated  string `json:"updated"`
				Username string `json:"username"`
			} `json:"property2"`
		} `json:"reactions"`
		Updated string `json:"updated"`
	} `json:"comments"`
	CoverImageAttachmentID int    `json:"cover_image_attachment_id"`
	Created                string `json:"created"`
	CreatedBy              struct {
		Created  string `json:"created"`
		Email    string `json:"email"`
		ID       int    `json:"id"`
		Name     string `json:"name"`
		Updated  string `json:"updated"`
		Username string `json:"username"`
	} `json:"created_by"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
	DoneAt      string `json:"done_at"`
	DueDate     string `json:"due_date"`
	EndDate     string `json:"end_date"`
	HexColor    string `json:"hex_color"`
	ID          int    `json:"id"`
	Identifier  string `json:"identifier"`
	Index       int    `json:"index"`
	IsFavorite  bool   `json:"is_favorite"`
	IsUnread    bool   `json:"is_unread"`
	Labels      []struct {
		Created   string `json:"created"`
		CreatedBy struct {
			Created  string `json:"created"`
			Email    string `json:"email"`
			ID       int    `json:"id"`
			Name     string `json:"name"`
			Updated  string `json:"updated"`
			Username string `json:"username"`
		} `json:"created_by"`
		Description string `json:"description"`
		HexColor    string `json:"hex_color"`
		ID          int    `json:"id"`
		Title       string `json:"title"`
		Updated     string `json:"updated"`
	} `json:"labels"`
	PercentDone int `json:"percent_done"`
	Position    int `json:"position"`
	Priority    int `json:"priority"`
	ProjectID   int `json:"project_id"`
	Reactions   struct {
		Property1 []struct {
			Created  string `json:"created"`
			Email    string `json:"email"`
			ID       int    `json:"id"`
			Name     string `json:"name"`
			Updated  string `json:"updated"`
			Username string `json:"username"`
		} `json:"property1"`
		Property2 []struct {
			Created  string `json:"created"`
			Email    string `json:"email"`
			ID       int    `json:"id"`
			Name     string `json:"name"`
			Updated  string `json:"updated"`
			Username string `json:"username"`
		} `json:"property2"`
	} `json:"reactions"`
	RelatedTasks struct {
		Property1 []struct {
		} `json:"property1"`
		Property2 []struct {
		} `json:"property2"`
	} `json:"related_tasks"`
	Reminders []struct {
		RelativePeriod int    `json:"relative_period"`
		RelativeTo     string `json:"relative_to"`
		Reminder       string `json:"reminder"`
	} `json:"reminders"`
	RepeatAfter  int    `json:"repeat_after"`
	RepeatMode   int    `json:"repeat_mode"`
	StartDate    string `json:"start_date"`
	Subscription struct {
		Created  string `json:"created"`
		Entity   int    `json:"entity"`
		EntityID int    `json:"entity_id"`
		ID       int    `json:"id"`
	} `json:"subscription"`
	Title   string `json:"title"`
	Updated string `json:"updated"`
}

type ProjectView struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	ViewKind  string `json:"view_kind"`
	ProjectID int    `json:"project_id"`
	Created   string `json:"created"`
	Updated   string `json:"updated"`
}

type ProjectViews []ProjectView

type Bucket struct {
	ID            int    `json:"id"`
	Title         string `json:"title"`
	Count         int    `json:"count"`
	Limit         int    `json:"limit"`
	ProjectViewID int    `json:"project_view_id"`
	Created       string `json:"created"`
	Updated       string `json:"updated"`
}

type Buckets []Bucket

type Label struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	HexColor    string `json:"hex_color"`
	Created     string `json:"created"`
	Updated     string `json:"updated"`
}

type Labels []Label

type LabelInput struct {
	Title       string `json:"title"`
	Description string `json:"description,omitempty"`
	HexColor    string `json:"hex_color,omitempty"`
}

type TaskComment struct {
	ID      int    `json:"id"`
	Comment string `json:"comment"`
	Author  struct {
		ID       int    `json:"id"`
		Username string `json:"username"`
		Name     string `json:"name"`
	} `json:"author"`
	Created string `json:"created"`
	Updated string `json:"updated"`
}

type TaskComments []TaskComment

type TaskCommentInput struct {
	Comment string `json:"comment"`
}

type TaskLabel struct {
	Description string `json:"description,omitempty"`
	Title       string `json:"title,omitempty"`
}

type TaskUpdate struct {
	Description string      `json:"description,omitempty"`
	Done        *bool       `json:"done,omitempty"`
	DueDate     string      `json:"due_date,omitempty"`
	EndDate     string      `json:"end_date,omitempty"`
	HexColor    string      `json:"hex_color,omitempty"`
	IsFavorite  *bool       `json:"is_favorite,omitempty"`
	Labels      []TaskLabel `json:"labels,omitempty"`
	PercentDone int         `json:"percent_done,omitempty"`
	Priority    int         `json:"priority,omitempty"`
	ProjectID   int         `json:"project_id,omitempty"`
	StartDate   string      `json:"start_date,omitempty"`
	Title       string      `json:"title,omitempty"`
}
