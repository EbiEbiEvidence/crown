package repository

import (
	"crowns/app/domain/query"
	"database/sql"
)

type CardRepository struct {
	db *sql.DB
}

func (repo *CardRepository) Save(name string) (userQueryModel *query.UserQueryModel, err error) {
	return nil, nil
}
