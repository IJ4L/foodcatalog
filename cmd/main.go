package main

import (
	"context"
	"log"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/ij4l/foodCatalog/apps"
	"github.com/ij4l/foodCatalog/apps/auth"
	database "github.com/ij4l/foodCatalog/database/postgres"
	"github.com/ij4l/foodCatalog/graph"
	"github.com/ij4l/foodCatalog/util"
	"github.com/jackc/pgx/v5"
)

func main() {
	config := loadConfig()

	pgxConn := connectDatabase(config)
	defer pgxConn.Close(context.Background())

	repo := apps.NewRepository(pgxConn)
	graphServer := initializeGraphQLServer(&repo)

	startServer(config, repo, graphServer)
}

func loadConfig() util.Config {
	config, err := util.LoadConfig(".env")
	if err != nil {
		log.Fatal("Cannot load configuration:", err)
	}
	return config
}

func connectDatabase(config util.Config) *pgx.Conn {
	pgxConn, err := database.ConnectPostgreSql(config)
	if err != nil {
		log.Fatal("Cannot connect to database:", err)
	}
	return pgxConn
}

func initializeGraphQLServer(repo *apps.AppRepository) *handler.Server {
	graphConfig := graphConfig(repo)
	return handler.NewDefaultServer(graph.NewExecutableSchema(graphConfig))
}

func startServer(config util.Config, repo apps.AppRepository, graphServer *handler.Server) {
	server, err := apps.NewServer(config, repo, graphServer)
	if err != nil {
		log.Fatal("Cannot create server:", err)
	}

	if err := server.Start(config.ServerAddress); err != nil {
		log.Fatal("Cannot start server:", err)
	}
}

func graphConfig(repo *apps.AppRepository) graph.Config {
	resolver := graph.Resolver{
		AuthService: auth.InitializeAuthService(repo),
	}

	return graph.Config{Resolvers: &resolver}
}
