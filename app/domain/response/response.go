package response

import "time"

type User struct {
	Name  string `json:"name"`
	Token string `json:"token"`
}

type HighScore struct {
	HighScoreID int64          `json:"highScoreID"`
	UserID      int64          `json:"userID"`
	Start       int64          `json:"start",omitempty`
	Age         int64          `json:"age",omitempty`
	CreatedAt   time.Time      `json:"created_at"`
	Scores      HighScoreScore `json:"scores"`
}

type HighScoreScore struct {
	ChurchScore    int64 `json:"church_score",omitempty`
	CommersScore   int64 `json:"commers_score",omitempty`
	MerchantsScore int64 `json:"merchants_score",omitempty`
	MilitaryScore  int64 `json:"military_score",omitempty`
	BonusScore     int64 `json:"bonus",omitempty`
}

type Game struct {
	GameID    int64     `json:"gameID"`
	UserID    int64     `json:"userID"`
	Score     int64     `json:"score"`
	CreatedAt time.Time `json:"createdAt"`
}
