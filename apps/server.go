package apps

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"github.com/ij4l/foodCatalog/middleware"
	"github.com/ij4l/foodCatalog/util"
)

type Server struct {
	config util.Config
	repo   AppRepository
	router *gin.Engine
	srv    *handler.Server
}

func NewServer(config util.Config, repo AppRepository, srv *handler.Server) (*Server, error) {
	server := &Server{repo: repo, config: config, srv: srv}

	util.InitToken(config.TokenSymmetricKey, config.AccessTokenDuration)

	server.setupRouter()

	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()

	router.Use(CORSMiddleware())

	router.GET("/", server.graphQLPlayground)

	router.Use(middleware.GinContextToContextMiddleware())

	router.POST("/query", server.graphQLHandler)

	server.router = router
}

func (server *Server) graphQLPlayground(c *gin.Context) {
	playground.Handler("/query", "GraphQL Playground").ServeHTTP(c.Writer, c.Request)
}

func (server *Server) graphQLHandler(c *gin.Context) {
	server.srv.ServeHTTP(c.Writer, c.Request)
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
