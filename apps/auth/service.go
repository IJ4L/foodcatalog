package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/ij4l/foodCatalog/graph/model"
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

func (ar *AuthService) Create(md model.NewUser) error {
	if err := ar.repo.create(md, ar.ctx); err != nil {
		return err
	}
	return nil
}

func (ar *AuthService) Login(email string) (*model.User, error) {
	user, err := ar.repo.getByEmail(email, ar.ctx)
	if err != nil {
		return nil, err
	}

	return user, nil
}
