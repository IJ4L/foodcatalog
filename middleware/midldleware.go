package middleware

import (
	"context"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/ij4l/foodCatalog/util"
	"github.com/rs/zerolog/log"
)

func GinContextToContextMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		bearer := c.GetHeader("Authorization")

		if bearer == "" {
			log.Error().Msg("no token provided")
			return
		}

		tokSlice := strings.Split(bearer, "Bearer ")
		if len(tokSlice) != 2 {
			log.Error().Msg("invalid token format")
			return
		}

		tokString := tokSlice[1]
		token, err := util.VerifyToken(tokString)
		if err != nil {
			log.Error().Err(err).Msg("failed to verify token")
			return
		}

		ctx := context.WithValue(c.Request.Context(), "AUTH_ID", token.Id)
		c.Request = c.Request.WithContext(ctx)

		c.Next()
	}
}

func GinContextFromContext(ctx context.Context) (*int, error) {
	authID := ctx.Value("AUTH_ID")
	if authID == nil {
		return nil, fmt.Errorf("could not retrieve AUTH_ID from context")
	}

	authIDInt, ok := authID.(int)
	if !ok {
		return nil, fmt.Errorf("could not convert AUTH_ID to int")
	}

	return &authIDInt, nil
}
