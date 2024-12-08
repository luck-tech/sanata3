package router

import "github.com/murasame29/go-httpserver-template/internal/usecase/interactor"

type di struct {
	login *interactor.Login
	user  *interactor.User
	skill *interactor.Skill
	room  *interactor.Room
	chat  *interactor.Chat
}

func NewDI(
	login *interactor.Login,
	user *interactor.User,
	skill *interactor.Skill,
	room *interactor.Room,
	chat *interactor.Chat,
) *di {
	return &di{
		login: login,
		user:  user,
		skill: skill,
		room:  room,
		chat:  chat,
	}
}
