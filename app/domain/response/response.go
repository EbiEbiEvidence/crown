package response

import "time"

type User struct {
	Name  string `json:"name"`
	Token string `json:"token"`
}

type HighScore struct {
	HighScoreID int64     `json:"highScoreID"`
	UserID      int64     `json:"userID"`
	Score       int64     `json:"score"`
	CreatedAt   time.Time `json:"createdAt"`
}

type Game struct {
	GameID    int64     `json:"gameID"`
	UserID    int64     `json:"userID"`
	Score     int64     `json:"score"`
	CreatedAt time.Time `json:"createdAt"`
}
