package command

type GameCommandModel struct {
	UserID int64
	Score  int64
}

func NewGameCommandModel(
	userID int64,
	score int64,
) *GameCommandModel {
	return &GameCommandModel{
		UserID: userID,
		Score:  score,
	}
}
