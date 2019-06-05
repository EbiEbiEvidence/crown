package repository

import (
	"crowns/app/domain/query"

	"github.com/jmoiron/sqlx"
)

type CardRepository struct {
	db *sqlx.DB
}

func NewCardRepository(db *sqlx.DB) *CardRepository {
	return &CardRepository{
		db,
	}
}

func (repo *CardRepository) FindRootRandomly() (cardQueryModel *query.CardQueryModel, err error) {
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

	cardQueryModel, err = repo.FindRootRandomlyTx(tx)
	if err != nil {
		return nil, err
	}
	return cardQueryModel, nil
}

func (repo *CardRepository) FindRootRandomlyTx(tx *sqlx.Tx) (cardQueryModel *query.CardQueryModel, err error) {
	cardQueryModel = &query.CardQueryModel{}
	err = tx.QueryRowx("SELECT card_id AS CardID, is_child AS IsChild, scenario_id AS ScenarioID, card_text AS CardText, card_image AS CardImage FROM cards ORDER BY RANDOM() LIMIT 1;").StructScan(cardQueryModel)
	if err != nil {
		return nil, err
	}

	return cardQueryModel, nil
}
