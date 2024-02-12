package main

import (
	"github.com/Rayato159/isekai-shop-api/config"
	"github.com/Rayato159/isekai-shop-api/databases"
	_playerEntity "github.com/Rayato159/isekai-shop-api/modules/player/entity"
)

func main() {
	appConfig := config.GetAppConfig()
	database := databases.NewPostgresDatabase(appConfig.DatabaseConfig)

	uuidMigreate(database)
	playerMigrate(database)
}

func uuidMigreate(db databases.Database) {
	db.GetDb().Raw(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`).Scan(&struct{}{})
}

func playerMigrate(db databases.Database) {
	db.GetDb().Migrator().CreateTable(&_playerEntity.Player{})
}
