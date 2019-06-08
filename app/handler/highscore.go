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
			Start:       highScore.Start,
			Age:         highScore.Age,
			CreatedAt:   highScore.CreatedAt,
			Scores: response.HighScoreScore{
				ChurchScore:    highScore.ChurchScore,
				CommersScore:   highScore.CommersScore,
				MerchantsScore: highScore.MerchantsScore,
				MilitaryScore:  highScore.MilitaryScore,
				BonusScore:     highScore.BonusScore,
			},
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

	if req.Start == 0 {
		marshallErrorResponse("start must not be empty", w)
		return
	}

	if req.Age == 0 {
		marshallErrorResponse("age must not be empty", w)
		return
	}

	if req.ChurchScore == 0 {
		marshallErrorResponse("churchScore must not be empty", w)
		return
	}

	if req.CommersScore == 0 {
		marshallErrorResponse("commersScore must not be empty", w)
		return
	}

	if req.MerchantsScore == 0 {
		marshallErrorResponse("merchantsScore must not be empty", w)
		return
	}

	if req.MilitaryScore == 0 {
		marshallErrorResponse("militaryScore must not be empty", w)
		return
	}

	if req.BonusScore == 0 {
		marshallErrorResponse("bonusScore must not be empty", w)
		return
	}

	user, err := h.userUseCase.FindByTokenOrName(req.Token, "")
	if err != nil {
		marshallErrorResponse(err.Error(), w)
		return
	}

	highScores, err := h.highScoreUseCase.Submit(
		user.UserID,
		req.Start,
		req.Age,
		req.ChurchScore+req.CommersScore+req.MerchantsScore+req.MilitaryScore+req.BonusScore,
		req.ChurchScore,
		req.CommersScore,
		req.MerchantsScore,
		req.MilitaryScore,
		req.BonusScore,
	)
	if err != nil {
		marshallErrorResponse(err.Error(), w)
		return
	}

	httpRes := []response.HighScore{}
	for _, highScore := range highScores {
		httpRes = append(httpRes, response.HighScore{
			HighScoreID: highScore.HighScoreID,
			UserID:      highScore.UserID,
			Start:       highScore.Start,
			Age:         highScore.Age,
			CreatedAt:   highScore.CreatedAt,
			Scores: response.HighScoreScore{
				ChurchScore:    highScore.ChurchScore,
				CommersScore:   highScore.CommersScore,
				MerchantsScore: highScore.MerchantsScore,
				MilitaryScore:  highScore.MilitaryScore,
				BonusScore:     highScore.BonusScore,
			},
		})
	}

	marshallResponse(httpRes, w)
}
