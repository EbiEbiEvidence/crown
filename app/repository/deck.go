package repository

import (
	"crowns/app/domain/query"
	"database/sql"
)

type DeckRepository struct {
	db *sql.DB
}

func (repo *DeckRepository) Save(name string) (userQueryModel *query.UserQueryModel, err error) {
	return nil, nil
}
