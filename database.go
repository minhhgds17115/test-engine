package databasego

type UserInfo struct {
	UserID    int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

type Question struct {
	questions   []Question `json:"questions"`
	ID          int64      `json:"id"`
	UserID      int64      `json:"user_id"`
	multiple    bool       `json:"multiple"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Type        string     `json:"type"`
	Answers     []Answer   `json:"answers"`
}

type Answer struct {
}
