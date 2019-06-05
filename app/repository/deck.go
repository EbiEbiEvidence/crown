package repository

import (
	"crowns/app/domain/query"

	"github.com/jmoiron/sqlx"
)

type DeckRepository struct {
	db *sqlx.DB
}

func NewDeckRepository(db *sqlx.DB) *DeckRepository {
	return &DeckRepository{
		db,
	}
}

func (repo *DeckRepository) Save(gameID int64, cardID int64) (deckQueryModel *query.DeckQueryModel, err error) {
	tx, err := repo.db.Beginx()
	if err != nil {
		return nil, err
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	deckQueryModel, err = repo.SaveTx(gameID, cardID, tx)
	if err != nil {
		return nil, err
	}
	return deckQueryModel, nil
}

func (repo *DeckRepository) SaveTx(gameID int64, cardID int64, tx *sqlx.Tx) (deckQueryModel *query.DeckQueryModel, err error) {
	deckQueryModel = &query.DeckQueryModel{}
	err = tx.QueryRowx("INSERT INTO decks (game_id, card_id) VALUES ($1, $2) RETURNING deck_id AS DeckID, game_id AS GameID, card_id AS CardID", gameID, cardID).StructScan(deckQueryModel)
	if err != nil {
		return nil, err
	}
	return deckQueryModel, nil
}

func (repo *DeckRepository) FindByID(deckID int64) (deckQueryModel *query.DeckQueryModel, err error) {
	deckQueryModel = &query.DeckQueryModel{}
	err = repo.db.QueryRowx("SELECT (deck_id, game_id, card_id) FROM decks WHERE deck_id = $1", deckID).StructScan(deckQueryModel)
	if err != nil {
		return nil, err
	}
	return deckQueryModel, nil
}
