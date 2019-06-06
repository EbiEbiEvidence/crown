package repository

import (
	"crowns/app/domain/query"

	"github.com/jmoiron/sqlx"
)

type CardRepository struct {
	Repository
}

func NewCardRepository(db *sqlx.DB) *CardRepository {
	return &CardRepository{
		Repository{
			db,
		},
	}
}

func (repo *CardRepository) FindRootRandomly(tx *sqlx.Tx) (cardQueryModel *query.CardQueryModel, err error) {
	cardQueryModel = &query.CardQueryModel{}
	if tx == nil {
		err = repo.db.QueryRowx("SELECT card_id AS CardID, is_child AS IsChild, scenario_id AS ScenarioID, card_text AS CardText, card_image AS CardImage FROM cards ORDER BY RANDOM() LIMIT 1;").StructScan(cardQueryModel)
	} else {
		err = tx.QueryRowx("SELECT card_id AS CardID, is_child AS IsChild, scenario_id AS ScenarioID, card_text AS CardText, card_image AS CardImage FROM cards ORDER BY RANDOM() LIMIT 1;").StructScan(cardQueryModel)
	}
	if err != nil {
		return nil, err
	}

	return cardQueryModel, nil
}
