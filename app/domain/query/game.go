package query

import (
	"database/sql"
	"time"

	"github.com/lib/pq"
)

type GameQueryModel struct {
	GameID    int64
	UserID    int64
	Score     int64
	CreatedAt time.Time
}

func NewGameQueryModel(
	gameID sql.NullInt64,
	userID sql.NullInt64,
	score sql.NullInt64,
	createdAt pq.NullTime,
) *GameQueryModel {
	return &GameQueryModel{
		GameID:    gameID.Int64,
		UserID:    userID.Int64,
		Score:     score.Int64,
		CreatedAt: createdAt.Time,
	}
}
