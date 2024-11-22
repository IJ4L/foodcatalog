package auth

import (
	"log"

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

func (ap authRepository) create(mu model.NewUser, ctx *gin.Context) (err error) {
	arg := db.InsertUserParams{
		Email:    mu.Email,
		Password: mu.Password,
	}

	if err = ap.repo.InsertUser(ctx, arg); err != nil {
		log.Println("error inserting user: ", err)
		return
	}

	return
}

func (ap authRepository) getByEmail(email string, ctx *gin.Context) (mu *model.User, err error) {
	auth, err := ap.repo.SelectUserByEmail(ctx, email)
	if err != nil {
		log.Println("error selecting user by email: ", err)
		return
	}

	mu = &model.User{
		ID:       int(auth.ID),
		Email:    auth.Email,
		Password: auth.Password,
	}

	return
}
