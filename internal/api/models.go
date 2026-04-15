package api

type Projects []struct {
	BackgroundBlurHash    string `json:"background_blur_hash"`
	BackgroundInformation any    `json:"background_information"`
	Created               string `json:"created"`
	Description           string `json:"description"`
	HexColor              string `json:"hex_color"`
	ID                    int    `json:"id"`
	Identifier            string `json:"identifier"`
	IsArchived            bool   `json:"is_archived"`
	IsFavorite            bool   `json:"is_favorite"`
	MaxPermission         int    `json:"max_permission"`
	Owner                 struct {
		Created  string `json:"created"`
		Email    string `json:"email"`
		ID       int    `json:"id"`
		Name     string `json:"name"`
		Updated  string `json:"updated"`
		Username string `json:"username"`
	} `json:"owner"`
	ParentProjectID int `json:"parent_project_id"`
	Position        int `json:"position"`
	Subscription    struct {
		Created  string `json:"created"`
		Entity   int    `json:"entity"`
		EntityID int    `json:"entity_id"`
		ID       int    `json:"id"`
	} `json:"subscription"`
	Title   string `json:"title"`
	Updated string `json:"updated"`
	Views   []struct {
		BucketConfiguration []struct {
			Filter struct {
				Filter             string   `json:"filter"`
				FilterIncludeNulls bool     `json:"filter_include_nulls"`
				OrderBy            []string `json:"order_by"`
				S                  string   `json:"s"`
				SortBy             []string `json:"sort_by"`
			} `json:"filter"`
			Title string `json:"title"`
		} `json:"bucket_configuration"`
		BucketConfigurationMode string `json:"bucket_configuration_mode"`
		Created                 string `json:"created"`
		DefaultBucketID         int    `json:"default_bucket_id"`
		DoneBucketID            int    `json:"done_bucket_id"`
		Filter                  struct {
			Filter             string   `json:"filter"`
			FilterIncludeNulls bool     `json:"filter_include_nulls"`
			OrderBy            []string `json:"order_by"`
			S                  string   `json:"s"`
			SortBy             []string `json:"sort_by"`
		} `json:"filter"`
		ID        int    `json:"id"`
		Position  int    `json:"position"`
		ProjectID int    `json:"project_id"`
		Title     string `json:"title"`
		Updated   string `json:"updated"`
		ViewKind  string `json:"view_kind"`
	} `json:"views"`
}
