package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/ij4l/foodCatalog/apps"
)

func InitializeAuthService(repo *apps.AppRepository) AuthService {
	ctx := &gin.Context{}
	authRepo := NewAuthRepository(repo)
	return NewAuthService(authRepo, ctx)
}
