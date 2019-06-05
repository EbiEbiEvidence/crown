package usecase

import (
	"crowns/app/domain/command"
	"crowns/app/domain/query"
	"crowns/app/repository"
	"errors"
	"strings"

	"github.com/google/uuid"
)

type UserUseCase struct {
	userRepo *repository.UserRepository
}

func NewUserUseCase(userRepo *repository.UserRepository) *UserUseCase {
	return &UserUseCase{
		userRepo,
	}
}

func (uc *UserUseCase) FindByTokenOrName(token string, name string) (userQueryModel *query.UserQueryModel, err error) {
	if token == "" && name == "" {
		return nil, errors.New("name or token must not be empty")
	}

	if token != "" {
		userQueryModel, err = uc.userRepo.FindByToken(token)
	} else {
		userQueryModel, err = uc.userRepo.FindByName(name)
	}

	if err != nil || userQueryModel == nil {
		return nil, errors.New("not found")
	}

	return userQueryModel, nil
}

func (uc *UserUseCase) Save(name string) (userQueryModel *query.UserQueryModel, err error) {
	if name == "" {
		return nil, errors.New("name must not be empty")
	}

	formerUser, err := uc.userRepo.FindByNameCaseInsensitive(name)
	if err != nil && !strings.Contains(err.Error(), "no rows in result set") {
		return nil, err
	}
	if formerUser != nil {
		return nil, errors.New("user already exists")
	}

	tokenToSave, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}

	userID, err := uc.userRepo.Save(command.NewUserCommandModel(name, tokenToSave.String()))
	if err != nil {
		return nil, err
	}
	if userID == nil {
		return nil, errors.New("Failed to get userID")
	}

	return uc.userRepo.FindByID(*userID)
}
