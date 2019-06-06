package server

import (
	"github.com/julienschmidt/httprouter"

	"crowns/app/handler"
	"crowns/app/repository"
	"crowns/app/usecase"
)

func (s *SimpleServer) Router() *httprouter.Router {
	r := httprouter.New()
	userRepo := repository.NewUserRepository(s.db)
	gameRepo := repository.NewGameRepository(s.db)
	cardRepo := repository.NewCardRepository(s.db)
	deckRepo := repository.NewDeckRepository(s.db)
	highScoreRepo := repository.NewHighScoreRepository(s.db)

	userUseCase := usecase.NewUserUseCase(userRepo)
	highScoreUseCase := usecase.NewHighScoreUseCase(highScoreRepo)
	gameUseCase := usecase.NewGameUseCase(gameRepo, cardRepo, deckRepo, userRepo)

	userHandler := handler.NewUserHandler(userUseCase, highScoreUseCase)
	highScoreHandler := handler.NewHighScoreHandler(userUseCase, highScoreUseCase)
	gameHandler := handler.NewGameHandler(gameUseCase)

	r.POST("/user/get", userHandler.Get)
	r.POST("/user/create", userHandler.Create)
	r.POST("/user/highscores", userHandler.GetHighScores)
	r.POST("/game/new", gameHandler.Start)
	r.POST("/highscores/submit", highScoreHandler.Submit)

	return r
}
