package usecase

import (
	"crowns/app/domain/command"
	"crowns/app/domain/query"
	"crowns/app/repository"
	"errors"

	"github.com/jmoiron/sqlx"
)

type HighScoreUseCase struct {
	highScoreRepo *repository.HighScoreRepository
}

func NewHighScoreUseCase(highScoreRepo *repository.HighScoreRepository) *HighScoreUseCase {
	return &HighScoreUseCase{
		highScoreRepo,
	}
}

func (uc *HighScoreUseCase) IndexUser(userID int64) (highScoreQueryModels []query.HighScoreQueryModel, err error) {
	highScoreQueryModels, err = uc.highScoreRepo.IndexUser(userID, nil)
	if err != nil {
		return nil, err
	}
	if highScoreQueryModels == nil {
		return nil, errors.New("Failed to index highScoreQueryModels")
	}

	return highScoreQueryModels, nil
}

func (uc *HighScoreUseCase) Submit(userID int64, score int64) (highScoreQueryModels []query.HighScoreQueryModel, err error) {
	ret, err := uc.highScoreRepo.ExecTx(
		func(tx *sqlx.Tx) (highScoreQueryModelsI interface{}, err error) {
			_, err = uc.highScoreRepo.Save(*command.NewHighScoreCommandModel(userID, score), tx)
			if err != nil {
				return nil, err
			}

			return uc.highScoreRepo.IndexUser(userID, tx)
		})

	highScoreQueryModels = ret.([]query.HighScoreQueryModel)

	if err != nil {
		return nil, err
	}
	if highScoreQueryModels == nil {
		return nil, errors.New("Failed to save highScoreQueryModel")
	}

	return ret.([]query.HighScoreQueryModel), nil
}
