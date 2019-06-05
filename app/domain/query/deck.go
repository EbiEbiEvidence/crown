package query

import "database/sql"

type DeckQueryModel struct {
	DeckID int64
	GameID int64
	CardID int64
}

func NewDeckQueryModel(deckID sql.NullInt64, gameID sql.NullInt64, cardID sql.NullInt64) *DeckQueryModel {
	return &DeckQueryModel{
		DeckID: deckID.Int64,
		GameID: gameID.Int64,
		CardID: cardID.Int64,
	}
}
