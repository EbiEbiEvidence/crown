package repository

import (
	"crowns/app/domain/query"

	"github.com/jmoiron/sqlx"
)

type DeckRepository struct {
	Repository
}

func NewDeckRepository(db *sqlx.DB) *DeckRepository {
	return &DeckRepository{
		Repository{
			db,
		},
	}
}

func (repo *DeckRepository) Save(gameID int64, cardID int64, tx *sqlx.Tx) (deckQueryModel *query.DeckQueryModel, err error) {
	deckQueryModel = &query.DeckQueryModel{}
	if tx == nil {
		err = repo.db.QueryRowx("INSERT INTO decks (game_id, card_id) VALUES ($1, $2) RETURNING deck_id AS DeckID, game_id AS GameID, card_id AS CardID", gameID, cardID).StructScan(deckQueryModel)
	} else {
		err = tx.QueryRowx("INSERT INTO decks (game_id, card_id) VALUES ($1, $2) RETURNING deck_id AS DeckID, game_id AS GameID, card_id AS CardID", gameID, cardID).StructScan(deckQueryModel)
	}

	if err != nil {
		return nil, err
	}
	return deckQueryModel, nil
}

func (repo *DeckRepository) FindByID(deckID int64, tx *sqlx.Tx) (deckQueryModel *query.DeckQueryModel, err error) {
	deckQueryModel = &query.DeckQueryModel{}
	if tx == nil {
		err = repo.db.QueryRowx("SELECT (deck_id, game_id, card_id) FROM decks WHERE deck_id = $1", deckID).StructScan(deckQueryModel)
	} else {
		err = tx.QueryRowx("SELECT (deck_id, game_id, card_id) FROM decks WHERE deck_id = $1", deckID).StructScan(deckQueryModel)
	}
	if err != nil {
		return nil, err
	}
	return deckQueryModel, nil
}
