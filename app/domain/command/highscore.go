package command

type HighScoreCommandModel struct {
	UserID int64
	Score  int64
}

func NewHighScoreCommandModel(
	userID int64,
	score int64,
) *HighScoreCommandModel {
	return &HighScoreCommandModel{
		UserID: userID,
		Score:  score,
	}
}
