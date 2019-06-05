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

	userUseCase := usecase.NewUserUseCase(userRepo)
	gameUseCase := usecase.NewGameUseCase(gameRepo, cardRepo, deckRepo, userRepo)

	userHandler := handler.NewUserHandler(userUseCase)
	gameHandler := handler.NewGameHandler(gameUseCase)

	r.POST("/user/get", userHandler.Get)
	r.POST("/user/create", userHandler.Create)
	r.POST("/game/new", gameHandler.Start)

	return r
}
