package repository

import (
	"github.com/jmoiron/sqlx"

	"crowns/app/domain/query"
)

type GameRepository struct {
	db *sqlx.DB
}

func NewGameRepository(db *sqlx.DB) *GameRepository {
	return &GameRepository{
		db,
	}
}

type TxFunc func(*sqlx.Tx) (interface{}, error)

func (repo *GameRepository) ExecTx(f TxFunc) (ret interface{}, err error) {
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

func (repo *GameRepository) Save(userID int64) (gameQueryModel *query.GameQueryModel, err error) {
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

	gameQueryModel, err = repo.SaveTx(userID, tx)
	if err != nil {
		return nil, err
	}
	return gameQueryModel, nil
}

func (repo *GameRepository) SaveTx(userID int64, tx *sqlx.Tx) (gameQueryModel *query.GameQueryModel, err error) {
	gameQueryModel = &query.GameQueryModel{}
	err = tx.QueryRowx("INSERT INTO games (user_id) VALUES ($1) RETURNING game_id AS GameID, user_id AS UserID, created_at AS CreatedAt, score AS Score", userID).StructScan(gameQueryModel)
	if err != nil {
		return nil, err
	}
	return gameQueryModel, nil
}
