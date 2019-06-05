package usecase

import (
	"crowns/app/repository"
)

type CardUseCase struct {
	cardRepo *repository.CardRepository
}

func NewCardUseCase(cardRepo *repository.CardRepository) *CardUseCase {
	return &CardUseCase{
		cardRepo,
	}
}
