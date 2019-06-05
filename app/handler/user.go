package handler

import (
	"crowns/app/domain/request"
	"crowns/app/domain/response"
	"crowns/app/usecase"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type UserHandler struct {
	userUseCase *usecase.UserUseCase
}

func NewUserHandler(userUseCase *usecase.UserUseCase) *UserHandler {
	return &UserHandler{
		userUseCase,
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
