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

func (ap authRepository) create(mu model.NewUser, ctx *gin.Context) (ur *model.UserResponse, err error) {
	arg := db.InsertUserParams{
		Email:    mu.Email,
		Password: mu.Password,
	}

	user, err := ap.repo.InsertUser(ctx, arg)
	if err != nil {
		return
	}

	ur = &model.UserResponse{
		ID:        int(user.ID),
		Email:     user.Email,
		CratedAt:  user.CreatedAt.Time.String(),
		UpdatedAt: user.UpdatedAt.Time.String(),
	}

	return
}

func (ap authRepository) getByEmail(email string, ctx *gin.Context) (ur *model.User, err error) {
	auth, err := ap.repo.SelectUserByEmail(ctx, email)
	if err != nil {
		return
	}

	ur = &model.User{
		ID:        int(auth.ID),
		Email:     auth.Email,
		Password:  auth.Password,
		CratedAt:  auth.CreatedAt.Time.String(),
		UpdatedAt: auth.UpdatedAt.Time.String(),
	}

	return
}
