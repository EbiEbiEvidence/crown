package repository

import (
	"github.com/jmoiron/sqlx"

	"crowns/app/domain/query"
)

type GameRepository struct {
	Repository
}

func NewGameRepository(db *sqlx.DB) *GameRepository {
	return &GameRepository{
		Repository{
			db,
		},
	}
}

func (repo *GameRepository) Save(userID int64, tx *sqlx.Tx) (gameQueryModel *query.GameQueryModel, err error) {
	gameQueryModel = &query.GameQueryModel{}
	if tx == nil {
		err = repo.db.QueryRowx("INSERT INTO games (user_id) VALUES ($1) RETURNING game_id AS GameID, user_id AS UserID, created_at AS CreatedAt, score AS Score", userID).StructScan(gameQueryModel)
	} else {
		err = tx.QueryRowx("INSERT INTO games (user_id) VALUES ($1) RETURNING game_id AS GameID, user_id AS UserID, created_at AS CreatedAt, score AS Score", userID).StructScan(gameQueryModel)
	}
	if err != nil {
		return nil, err
	}
	return gameQueryModel, nil
}
