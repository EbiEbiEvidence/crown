package repository

import (
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	db *sqlx.DB
}

func (repo *Repository) ExecTx(f func(*sqlx.Tx) (interface{}, error)) (ret interface{}, err error) {
	tx, err := repo.db.Beginx()
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()
	return f(tx)
}
