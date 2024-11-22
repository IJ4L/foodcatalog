package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/ij4l/foodCatalog/apps"
)

func InitializeAuthHandler(repo *apps.AppRepository) (authHandler AuthHandler) {
	ctx := &gin.Context{}
	authRepo := NewAuthRepository(repo)
	authService := NewAuthService(authRepo, ctx)
	authHandler = NewAuthHandler(authService)
	return
}
