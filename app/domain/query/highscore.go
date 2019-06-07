package query

import (
	"time"
)

type HighScoreQueryModel struct {
	HighScoreID    int64
	UserID         int64
	Start          int64
	Age            int64
	Score          int64
	ChurchScore    int64
	CommersScore   int64
	MerchantsScore int64
	MilitaryScore  int64
	BonusScore     int64
	CreatedAt      time.Time
}
