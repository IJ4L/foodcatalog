package auth

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/ij4l/foodCatalog/graph/model"
	"github.com/ij4l/foodCatalog/util"
	"github.com/rs/zerolog/log"
)

type authRepositoryContract interface {
	create(md model.NewUser, ctx *gin.Context) error
	getByEmail(email string, ctx *gin.Context) (*model.User, error)
}

type AuthService struct {
	ctx  *gin.Context
	repo authRepositoryContract
}

func NewAuthService(repo authRepositoryContract, ctx *gin.Context) AuthService {
	return AuthService{repo: repo, ctx: ctx}
}

func (ar *AuthService) create(mu model.NewUser) (err error) {
	mu.Password, err = util.Hash(mu.Password)
	if err != nil {
		log.Error().Err(err).Msg("error hashing password")
		return
	}

	err = ar.repo.create(mu, ar.ctx)
	if err != nil {
		log.Error().Err(err).Msg("error creating user")
		return
	}

	return
}

func (ar *AuthService) login(email string, pass string) (ma *model.AuthPayload, err error) {
	ud, err := ar.repo.getByEmail(email, ar.ctx)
	if err != nil {
		log.Error().Err(err).Msg("error getting user by email")
		if err.Error() == "no rows in result set" {
			err = fmt.Errorf("user not found")
		}
		return
	}

	err = util.Verify(ud.Password, pass)
	if err != nil {
		log.Error().Err(err).Msg("error verifying password")
		err = fmt.Errorf("email and password not match")
		return
	}

	tokenJWT := util.NewJWT(ud.ID)
	token, err := tokenJWT.GenerateToken()
	ma = &model.AuthPayload{
		Token: &token,
		User: &model.User{
			ID:    ud.ID,
			Email: ud.Email,
		},
	}

	return
}
