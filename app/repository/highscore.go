package repository

import (
	"crowns/app/domain/command"
	"crowns/app/domain/query"
	"strings"

	"github.com/jmoiron/sqlx"
)

type HighScoreRepository struct {
	Repository
}

func NewHighScoreRepository(db *sqlx.DB) *HighScoreRepository {
	return &HighScoreRepository{
		Repository{
			db,
		},
	}
}

func (repo *HighScoreRepository) Save(highScoreCommand command.HighScoreCommandModel, tx *sqlx.Tx) (highScoreQueryModel *query.HighScoreQueryModel, err error) {
	const sqlQueryToSave = "INSERT INTO high_scores (user_id, score) VALUES ($1, $2) RETURNING high_score_id AS HighScoreID, user_id AS UserID, created_at AS CreatedAt, score AS Score"
	const sqlQueryToRemoveLowRank = `
		WITH
		ranked_high_scores AS (
			SELECT
				*,
				RANK() OVER(PARTITION BY user_id ORDER BY score) AS rank
			FROM high_scores
		),
		low_rank_highscore_ids AS (
			SELECT
				high_score_id
			FROM ranked_high_scores
			WHERE rank > 5
		)
		DELETE FROM high_scores
		WHERE high_score_id IN (SELECT * FROM low_rank_highscore_ids)
	`

	highScoreQueryModel = &query.HighScoreQueryModel{}
	if tx == nil {
		ret, err := repo.ExecTx(func(txx *sqlx.Tx) (interface{}, error) {
			return repo.Save(highScoreCommand, txx)
		})
		if err != nil {
			return nil, err
		}
		return ret.(*query.HighScoreQueryModel), nil
	}

	err = tx.QueryRowx(sqlQueryToSave, highScoreCommand.UserID, highScoreCommand.Score).StructScan(highScoreQueryModel)
	if err != nil {
		return nil, err
	}

	err = tx.QueryRowx(sqlQueryToRemoveLowRank).Scan()
	if err != nil {
		if strings.Contains(err.Error(), "no rows in result set") {
			return highScoreQueryModel, nil
		}
		return nil, err
	}

	return highScoreQueryModel, nil
}

func (repo *HighScoreRepository) IndexUser(userID int64, tx *sqlx.Tx) (highScoreQueryModels []query.HighScoreQueryModel, err error) {
	const sqlQuery = "SELECT high_score_id AS HighScoreID, user_id AS UserID, created_at AS CreatedAt, score AS Score FROM high_scores WHERE user_id = $1 ORDER BY score LIMIT 5"
	highScoreQueryModels = []query.HighScoreQueryModel{}
	if tx == nil {
		err = repo.db.Select(&highScoreQueryModels, sqlQuery, userID)
	} else {
		err = tx.Select(&highScoreQueryModels, sqlQuery, userID)
	}
	if err != nil {
		return nil, err
	}
	return highScoreQueryModels, nil
}

func (repo *HighScoreRepository) Index(tx *sqlx.Tx) (highScoreQueryModels []query.HighScoreQueryModel, err error) {
	const sqlQuery = "SELECT high_score_id AS HighScoreID, user_id AS UserID, created_at AS CreatedAt, score AS Score FROM high_scores ORDER BY score LIMIT 10"
	highScoreQueryModels = []query.HighScoreQueryModel{}
	if tx == nil {
		err = repo.db.Select(&highScoreQueryModels, sqlQuery)
	} else {
		err = tx.Select(&highScoreQueryModels, sqlQuery)
	}
	if err != nil {
		return nil, err
	}
	return highScoreQueryModels, nil
}
