package repository

import (
	"crowns/app/domain/command"
	"crowns/app/domain/query"
	"database/sql"
	"errors"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{
		db,
	}
}

func (repo *UserRepository) FindByName(name string) (userQueryModel *query.UserQueryModel, err error) {
	userID := &sql.NullInt64{}
	email := &sql.NullString{}
	displayName := &sql.NullString{}
	token := &sql.NullString{}
	createdAt := &pq.NullTime{}

	err = repo.db.QueryRow("SELECT user_id, email, display_name, token, created_at FROM users WHERE users.display_name = $1", name).Scan(userID, email, displayName, token, createdAt)
	if err != nil {
		return nil, err
	}

	if userID == nil || email == nil || displayName == nil || token == nil || createdAt == nil {
		return nil, errors.New("Cannot allow null object")
	}

	userQueryModel = query.NewUserQueryModel(*userID, *email, *displayName, *token, *createdAt)

	return userQueryModel, nil
}

func (repo *UserRepository) FindByNameCaseInsensitive(name string, tx *sqlx.Tx) (userQueryModel *query.UserQueryModel, err error) {
	userID := &sql.NullInt64{}
	email := &sql.NullString{}
	displayName := &sql.NullString{}
	token := &sql.NullString{}
	createdAt := &pq.NullTime{}
	if tx == nil {
		err = repo.db.QueryRow("SELECT user_id, email, display_name, token, created_at FROM users WHERE LOWER(users.display_name) = LOWER($1)", name).Scan(userID, email, displayName, token, createdAt)
	} else {
		err = tx.QueryRow("SELECT user_id, email, display_name, token, created_at FROM users WHERE LOWER(users.display_name) = LOWER($1)", name).Scan(userID, email, displayName, token, createdAt)
	}
	if err != nil {
		return nil, err
	}

	if userID == nil || email == nil || displayName == nil || token == nil || createdAt == nil {
		return nil, errors.New("Cannot allow null object")
	}

	userQueryModel = query.NewUserQueryModel(*userID, *email, *displayName, *token, *createdAt)

	return userQueryModel, nil
}

func (repo *UserRepository) FindByToken(token string, tx *sqlx.Tx) (userQueryModel *query.UserQueryModel, err error) {
	userID := &sql.NullInt64{}
	email := &sql.NullString{}
	displayName := &sql.NullString{}
	tokenToFind := &sql.NullString{}
	createdAt := &pq.NullTime{}

	if tx == nil {
		err = repo.db.QueryRow("SELECT user_id, email, display_name, token, created_at FROM users WHERE users.token = $1", token).Scan(userID, email, displayName, tokenToFind, createdAt)
	} else {
		err = tx.QueryRow("SELECT user_id, email, display_name, token, created_at FROM users WHERE users.token = $1", token).Scan(userID, email, displayName, tokenToFind, createdAt)
	}

	if err != nil {
		return nil, err
	}

	if userID == nil || email == nil || displayName == nil || tokenToFind == nil || createdAt == nil {
		return nil, errors.New("Cannot allow null object")
	}

	userQueryModel = query.NewUserQueryModel(*userID, *email, *displayName, *tokenToFind, *createdAt)

	return userQueryModel, nil
}

func (repo *UserRepository) FindByID(userID int64, tx *sqlx.Tx) (userQueryModel *query.UserQueryModel, err error) {
	userIDToFind := &sql.NullInt64{}
	email := &sql.NullString{}
	displayName := &sql.NullString{}
	tokenToFind := &sql.NullString{}
	createdAt := &pq.NullTime{}

	if tx == nil {
		err = repo.db.QueryRow("SELECT user_id, email, display_name, token, created_at FROM users WHERE users.user_id = $1", userID).Scan(userIDToFind, email, displayName, tokenToFind, createdAt)
	} else {
		err = tx.QueryRow("SELECT user_id, email, display_name, token, created_at FROM users WHERE users.user_id = $1", userID).Scan(userIDToFind, email, displayName, tokenToFind, createdAt)
	}

	if err != nil {
		return nil, err
	}

	if userIDToFind == nil || email == nil || displayName == nil || tokenToFind == nil || createdAt == nil {
		return nil, errors.New("Cannot allow null object")
	}

	userQueryModel = query.NewUserQueryModel(*userIDToFind, *email, *displayName, *tokenToFind, *createdAt)

	return userQueryModel, nil
}

func (repo *UserRepository) Save(userCommandModel *command.UserCommandModel, tx *sqlx.Tx) (userID *int64, err error) {
	userIDToFind := &sql.NullInt64{}
	if tx == nil {
		err = repo.db.QueryRow("INSERT INTO users (email, display_name, token) VALUES ($1, $1, $2) RETURNING user_id", userCommandModel.Name, userCommandModel.Token).Scan(userIDToFind)
	} else {
		err = tx.QueryRow("INSERT INTO users (email, display_name, token) VALUES ($1, $1, $2) RETURNING user_id", userCommandModel.Name, userCommandModel.Token).Scan(userIDToFind)
	}

	if err != nil {
		return nil, err
	}

	if !userIDToFind.Valid {
		return nil, errors.New("failed to insert")
	}

	return &userIDToFind.Int64, nil
}
