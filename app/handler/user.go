package handler

import (
	"crowns/app/domain/request"
	"crowns/app/domain/response"
	"crowns/app/usecase"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type UserHandler struct {
	userUseCase      *usecase.UserUseCase
	highScoreUseCase *usecase.HighScoreUseCase
}

func NewUserHandler(userUseCase *usecase.UserUseCase, highScoreUseCase *usecase.HighScoreUseCase) *UserHandler {
	return &UserHandler{
		userUseCase,
		highScoreUseCase,
	}
}

func (h *UserHandler) Get(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	req := request.GetUser{}
	err := unmarshallRequest(&req, w, r)
	if err != nil {
		marshallErrorResponse(err.Error(), w)
		return
	}

	if req.Token == "" && req.Name == "" {
		marshallErrorResponse("name or token must not be empty", w)
		return
	}

	userQueryModel, err := h.userUseCase.FindByTokenOrName(req.Token, req.Name)
	if err != nil {
		marshallErrorResponse(err.Error(), w)
		return
	}

	userRes := response.NewUser(userQueryModel.DisplayName)

	marshallResponse(userRes, w)
}

func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	req := request.CreateUser{}
	err := unmarshallRequest(&req, w, r)
	if err != nil {
		marshallErrorResponse(err.Error(), w)
		return
	}

	if req.Name == "" {
		marshallErrorResponse("name must not be empty", w)
		return
	}

	userQueryModel, err := h.userUseCase.Save(req.Name)
	if err != nil {
		marshallErrorResponse(err.Error(), w)
		return
	}

	userRes := response.NewUser(userQueryModel.DisplayName)

	marshallResponse(userRes, w)
}

func (h *UserHandler) GetHighScores(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	req := request.GetHighScores{}

	if err := unmarshallRequest(&req, w, r); err != nil {
		marshallErrorResponse(err.Error(), w)
		return
	}

	if req.Token == "" {
		marshallErrorResponse("token must not be empty", w)
		return
	}

	user, err := h.userUseCase.FindByTokenOrName(req.Token, "")
	if err != nil {
		marshallErrorResponse(err.Error(), w)
		return
	}

	highScores, err := h.highScoreUseCase.IndexUser(user.UserID)
	if err != nil {
		marshallErrorResponse(err.Error(), w)
		return
	}
	if highScores == nil {
		marshallErrorResponse("could not found highScores", w)
		return
	}

	marshallResponse(highScores, w)
}
