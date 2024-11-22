package auth

import "github.com/ij4l/foodCatalog/graph/model"

type AuthHandler struct {
	as AuthService
}

func NewAuthHandler(as AuthService) AuthHandler {
	return AuthHandler{as: as}
}

func (ah AuthHandler) Register(md model.NewUser) (err error) {
	err = ah.as.create(md)
	return
}

func (ah AuthHandler) Login(email string, pass string) (ma *model.AuthPayload, err error) {
	ma, err = ah.as.login(email, pass)
	return
}
