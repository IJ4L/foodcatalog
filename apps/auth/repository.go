package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/ij4l/foodCatalog/apps"
	db "github.com/ij4l/foodCatalog/database/postgres/sqlc"
	"github.com/ij4l/foodCatalog/graph/model"
)

type authRepository struct {
	repo apps.AppRepository
}

func NewAuthRepository(repo *apps.AppRepository) authRepository {
	return authRepository{repo: *repo}
}

func (ap authRepository) create(md model.NewUser, ctx *gin.Context) error {
	arg := db.InsertUserParams{
		Email:    md.Email,
		Password: md.Password,
	}

	if err := ap.repo.InsertUser(ctx, arg); err != nil {
		return err
	}

	return nil
}

func (ap authRepository) getByEmail(email string, ctx *gin.Context) (*model.User, error) {
	auth, err := ap.repo.SelectUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	rsp := &model.User{
		ID:       int(auth.ID),
		Email:    auth.Email,
		Password: auth.Password,
	}

	return rsp, nil
}
