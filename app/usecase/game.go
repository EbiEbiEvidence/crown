package usecase

import (
	"crowns/app/domain/query"
	"crowns/app/repository"
	"errors"

	"github.com/jmoiron/sqlx"
)

type GameUseCase struct {
	gameRepo *repository.GameRepository
	cardRepo *repository.CardRepository
	deckRepo *repository.DeckRepository
	userRepo *repository.UserRepository
}

func NewGameUseCase(
	gameRepo *repository.GameRepository,
	cardRepo *repository.CardRepository,
	deckRepo *repository.DeckRepository,
	userRepo *repository.UserRepository,
) *GameUseCase {
	return &GameUseCase{
		gameRepo,
		cardRepo,
		deckRepo,
		userRepo,
	}
}

func (uc *GameUseCase) Start(token string) (gameQueryModel *query.GameQueryModel, err error) {
	ret, err := uc.gameRepo.ExecTx(
		func(tx *sqlx.Tx) (gameQueryModel interface{}, err error) {
			userQueryModel, err := uc.userRepo.FindByToken(token, tx)
			if err != nil {
				return nil, err
			}

			if userQueryModel == nil {
				return nil, errors.New("user not found")
			}

			gameQueryModel, err = uc.gameRepo.Save(userQueryModel.UserID, tx)
			if err != nil {
				return nil, err
			}

			if gameQueryModel == nil {
				return nil, errors.New("failed to save game")
			}

			for i := 0; i < 270; i++ {
				cardQueryModel, err := uc.cardRepo.FindRootRandomly(tx)
				if err != nil {
					return nil, err
				}

				uc.deckRepo.Save(gameQueryModel.(*query.GameQueryModel).GameID, cardQueryModel.CardID, tx)
			}

			return gameQueryModel, nil
		})

	if err != nil {
		return nil, err
	}
	return ret.(*query.GameQueryModel), err
}
