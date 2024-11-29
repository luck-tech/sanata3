package router

import "github.com/murasame29/go-httpserver-template/internal/usecase/interactor"

type di struct {
	login *interactor.Login
}

func NewDI(
	login *interactor.Login,
) *di {
	return &di{
		login: login,
	}
}
