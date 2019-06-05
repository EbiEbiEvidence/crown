package query

import (
	"database/sql"
	"time"

	"github.com/lib/pq"
)

type HighScoreQueryModel struct {
	UserID    int64
	Score     int64
	CreatedAt time.Time
}

func NewHighScoreQueryModel(
	userID sql.NullInt64,
	score sql.NullInt64,
	createdAt pq.NullTime,
) *HighScoreQueryModel {
	return &HighScoreQueryModel{
		UserID:    userID.Int64,
		Score:     score.Int64,
		CreatedAt: createdAt.Time,
	}
}
