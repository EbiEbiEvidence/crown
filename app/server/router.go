package server

import (
	"github.com/julienschmidt/httprouter"
)

func (s *SimpleServer) Router() *httprouter.Router {
	h := httprouter.New()
	return h
}
