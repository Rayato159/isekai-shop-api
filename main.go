package main

import (
	"github.com/Rayato159/isekai-shop-api/config"
	"github.com/Rayato159/isekai-shop-api/databases"
	"github.com/Rayato159/isekai-shop-api/server"
)

func main() {
	appConfig := config.GetAppConfig()
	database := databases.NewAppDatabase(appConfig.DatabaseConfig)
	server := server.NewAppServer(appConfig, database.GetDb())

	server.Start()
}
