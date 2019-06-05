package command

import "database/sql"

type DeckCommandModel struct {
	GameID int64
	CardID int64
}

func NewDeckCommandModel(gameID sql.NullInt64, cardID sql.NullInt64) *DeckCommandModel {
	return &DeckCommandModel{
		GameID: gameID.Int64,
		CardID: cardID.Int64,
	}
}
