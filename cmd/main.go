package main

import (
	"context"
	"log"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/ij4l/foodCatalog/apps"
	database "github.com/ij4l/foodCatalog/database/postgres"
	"github.com/ij4l/foodCatalog/graph"
	"github.com/ij4l/foodCatalog/util"
)

func main() {
	config, err := util.LoadConfig(".env")
	if err != nil {
		log.Fatal("Cannot load configuration :", err)
	}

	pgxConn, err := database.ConnectPostgreSql(config)
	if err != nil {
		log.Fatal("Cannot connect to database :", err)
	}
	defer pgxConn.Close(context.Background())

	graphConfig := graph.Config{Resolvers: &graph.Resolver{}}
	executableSchema := graph.NewExecutableSchema(graphConfig)
	graphServer := handler.NewDefaultServer(executableSchema)
	repo := apps.NewRepository(pgxConn)

	server, err := apps.NewServer(config, repo, graphServer)
	if err != nil {
		log.Fatal("cannot create server:", err)
	}

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
