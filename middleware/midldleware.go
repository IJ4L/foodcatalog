package middleware

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/ij4l/foodCatalog/util"
)

func GinContextToContextMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		bearer := c.GetHeader("Authorization")

		if bearer == "" {
			return
		}

		token, err := util.VerifyToken(bearer)
		if err != nil {
			return
		}

		ctx := context.WithValue(c.Request.Context(), "AUTH_ID", token.Id)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

func GinFromContext(ctx context.Context) (*int, error) {
	authID := ctx.Value("AUTH_ID")
	if authID == nil {
		return nil, fmt.Errorf("barer token not found")
	}

	authIDInt, ok := authID.(int)
	if !ok {
		return nil, fmt.Errorf("could not convert AUTH_ID to int")
	}

	return &authIDInt, nil
}
