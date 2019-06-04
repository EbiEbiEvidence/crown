package query

import (
	"database/sql"
	"time"

	"github.com/lib/pq"
)

type UserQueryModel struct {
	UserID      int64
	Email       string
	DisplayName string
	Token       string
	CreatedAt   time.Time
}

func NewUserQueryModel(
	userID sql.NullInt64,
	email sql.NullString,
	displayName sql.NullString,
	token sql.NullString,
	createdAt pq.NullTime,
) *UserQueryModel {
	return &UserQueryModel{
		UserID:      userID.Int64,
		Email:       email.String,
		DisplayName: displayName.String,
		Token:       token.String,
		CreatedAt:   createdAt.Time,
	}
}
