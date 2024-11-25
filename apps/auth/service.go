package auth

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/ij4l/foodCatalog/graph/model"
	"github.com/ij4l/foodCatalog/util"
)

type authRepositoryContract interface {
	create(md model.NewUser, ctx *gin.Context) (ur *model.UserResponse, err error)
	getByEmail(email string, ctx *gin.Context) (ur *model.User, err error)
}

type AuthService struct {
	ctx  *gin.Context
	repo authRepositoryContract
}

func NewAuthService(repo authRepositoryContract, ctx *gin.Context) AuthService {
	return AuthService{repo: repo, ctx: ctx}
}

func (ar *AuthService) create(mu model.NewUser) (ur *model.UserResponse, err error) {
	mu.Password, err = util.Hash(mu.Password)
	if err != nil {
		return
	}

	ur, err = ar.repo.create(mu, ar.ctx)
	if err != nil {
		return
	}

	return
}

func (ar *AuthService) login(email string, pass string) (ma *model.AuthPayload, err error) {
	ud, err := ar.repo.getByEmail(email, ar.ctx)
	if err != nil {
		if err.Error() == "no rows in result set" {
			err = fmt.Errorf("user not found")
		}
		return
	}

	err = util.Verify(ud.Password, pass)
	if err != nil {
		err = fmt.Errorf("email and password not match")
		return
	}

	tokenJWT := util.NewJWT(ud.ID)
	token, err := tokenJWT.GenerateToken()
	ma = &model.AuthPayload{
		Token: &token,
		User: &model.UserResponse{
			ID:    ud.ID,
			Email: ud.Email,
			CratedAt: ud.CratedAt,
			UpdatedAt: ud.UpdatedAt,
		},
	}

	return
}
