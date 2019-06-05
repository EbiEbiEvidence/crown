package server

import (
	"github.com/julienschmidt/httprouter"

	"crowns/app/handler"
	"crowns/app/repository"
	"crowns/app/usecase"
)

func (s *SimpleServer) Router() *httprouter.Router {
	h := httprouter.New()
	userRepository := repository.NewUserRepository(s.db)
	userUseCase := usecase.NewUserUseCase(userRepository)
	userHandler := handler.NewUserHandler(userUseCase)
	h.POST("/user/get", userHandler.Get)
	h.POST("/user/create", userHandler.Create)
	return h
}
