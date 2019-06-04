package repository

import (
	"crowns/app/domain/query"
	"database/sql"

	"github.com/lib/pq"
)

type UserRepository struct {
	db *sql.DB
}

func (repo *UserRepository) FindByName(name string) (query.UserQueryModel, error) {
	var userId *sql.NullInt64
	var email *sql.NullString
	var displayName *sql.NullString
	var token *sql.NullString
	var createdAt *pq.NullTime
	
	err := repo
		.db.QueryRow("SELECT user_id, email, display_name, token, created_at FROM users WHERE users.display_name = $1", name)
		.Scan(userId, email, displayName, token, createdAt)
}
