package main

import (
	"github.com/Rayato159/isekai-shop-api/config"
	"github.com/Rayato159/isekai-shop-api/databases"
	"github.com/Rayato159/isekai-shop-api/server"
)

func main() {
	appConfig := config.GetAppConfig()
	database := databases.NewPostgresDatabase(appConfig.DatabaseConfig)
	server := server.NewEchoServer(appConfig, database.GetDb())

	server.Start()
}
