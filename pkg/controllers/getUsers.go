package controllers

import (
	"database/sql"

	"github.com/anes011/gop/pkg/models"
)

type Pagination struct {
	Offset int
	Limit  int
}

func GetUsers(db *sql.DB, value *Pagination) ([]models.User, error) {
	users := []models.User{}

	rows, err := db.Query("SELECT id, name, email, password, photo, status, timestamp FROM users OFFSET $1 LIMIT $2", value.Offset, value.Limit)

	if err != nil {
		return users, err
	}

	defer rows.Close()

	for rows.Next() {
		user := models.User{}

		if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Photo,
			&user.Status, &user.Timestamp); err != nil {
			return users, err
		}

		//Get books of each user
		rows, err := db.Query("SELECT id, author_id, book_name, created_at FROM book WHERE author_id = $1", user.ID)

		if err != nil {
			return users, err
		}

		defer rows.Close()

		for rows.Next() {
			book := models.Book{}

			if err := rows.Scan(&book.ID, &book.AuthorID, &book.BookName, &book.CreatedAt); err != nil {
				return users, err
			}

			user.Books = append(user.Books, book)
		}

		users = append(users, user)
	}

	return users, nil
}
