package handler

import (
	"crowns/app/domain/request"
	"crowns/app/usecase"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type GameHandler struct {
	gameUseCase *usecase.GameUseCase
}

func NewGameHandler(gameUseCase *usecase.GameUseCase) *GameHandler {
	return &GameHandler{
		gameUseCase,
	}
}

func (h *GameHandler) Start(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	req := request.StartGame{}
	err := unmarshallRequest(&req, w, r)
	if err != nil {
		marshallErrorResponse(err.Error(), w)
		return
	}

	if req.Token == "" {
		marshallErrorResponse("token must not be empty", w)
		return
	}

	gameQueryModel, err := h.gameUseCase.Start(req.Token)
	if err != nil {
		marshallErrorResponse(err.Error(), w)
		return
	}

	marshallResponse(gameQueryModel.GameID, w)
}
