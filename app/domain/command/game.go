package command

type GameCommandModel struct {
	UserID       int64
	Score        int64
	EmperorImina string
	EmperorGengo string
}

func NewGameCommandModel(
	userID int64,
	score int64,
	emperorImina string,
	emperorGengo string,
) *GameCommandModel {
	return &GameCommandModel{
		UserID:       userID,
		Score:        score,
		EmperorImina: emperorImina,
		EmperorGengo: emperorGengo,
	}
}
