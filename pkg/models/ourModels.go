package models

type User struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Photo     string `json:"photo"`
	Status    bool   `json:"status"`
	Timestamp string `json:"timestamp"`
	Books     []Book `json:"books"`
}

type Book struct {
	ID        int    `json:"id"`
	AuthorID  int    `json:"author_id"`
	BookName  string `json:"book_name"`
	CreatedAt string `json:"created_at"`
	Author    User   `json:"author"`
}
