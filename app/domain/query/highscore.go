package query

import (
	"time"
)

type HighScoreQueryModel struct {
	HighScoreID int64
	UserID      int64
	Score       int64
	CreatedAt   time.Time
}
