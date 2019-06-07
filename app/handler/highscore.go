package handler

import (
	"crowns/app/domain/request"
	"crowns/app/domain/response"
	"crowns/app/usecase"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type HighScoreHandler struct {
	userUseCase      *usecase.UserUseCase
	highScoreUseCase *usecase.HighScoreUseCase
}

func NewHighScoreHandler(userUseCase *usecase.UserUseCase, highScoreUseCase *usecase.HighScoreUseCase) *HighScoreHandler {
	return &HighScoreHandler{
		userUseCase,
		highScoreUseCase,
	}
}

func (h *HighScoreHandler) GetHighScores(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	highScores, err := h.highScoreUseCase.Index()
	if err != nil {
		marshallErrorResponse(err.Error(), w)
		return
	}
	if highScores == nil {
		marshallErrorResponse("could not found highScores", w)
		return
	}

	if err != nil {
		marshallErrorResponse(err.Error(), w)
		return
	}

	httpRes := []response.HighScore{}
	for _, highScore := range highScores {
		httpRes = append(httpRes, response.HighScore{
			HighScoreID: highScore.HighScoreID,
			UserID:      highScore.UserID,
			Score:       highScore.Score,
			CreatedAt:   highScore.CreatedAt,
		})
	}

	marshallResponse(httpRes, w)
}

func (h *HighScoreHandler) Submit(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	req := request.SubmitHighScores{}
	err := unmarshallRequest(&req, w, r)
	if err != nil {
		marshallErrorResponse(err.Error(), w)
		return
	}

	if req.Token == "" {
		marshallErrorResponse("token must not be empty", w)
		return
	}

	if req.Score == 0 {
		marshallErrorResponse("score must not be empty", w)
		return
	}

	user, err := h.userUseCase.FindByTokenOrName(req.Token, "")
	if err != nil {
		marshallErrorResponse(err.Error(), w)
		return
	}

	highScores, err := h.highScoreUseCase.Submit(user.UserID, req.Score)
	if err != nil {
		marshallErrorResponse(err.Error(), w)
		return
	}

	httpRes := []response.HighScore{}
	for _, highScore := range highScores {
		httpRes = append(httpRes, response.HighScore{
			HighScoreID: highScore.HighScoreID,
			UserID:      highScore.UserID,
			Score:       highScore.Score,
			CreatedAt:   highScore.CreatedAt,
		})
	}

	marshallResponse(httpRes, w)
}
