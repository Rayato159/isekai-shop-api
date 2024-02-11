package main

import (
	"github.com/Rayato159/isekai-shop-api/config"
	"github.com/Rayato159/isekai-shop-api/databases"
	_oauth2Entity "github.com/Rayato159/isekai-shop-api/modules/oauth2/entity"
	_playerEntity "github.com/Rayato159/isekai-shop-api/modules/player/entity"
)

func main() {
	appConfig := config.GetAppConfig()
	database := databases.NewPostgresDatabase(appConfig.DatabaseConfig)

	uuidMigreate(database)
	playerMigrate(database)
	oauth2Migrate(database)
}

func uuidMigreate(db databases.Database) {
	db.GetDb().Raw(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`).Scan(&struct{}{})
}

func oauth2Migrate(db databases.Database) {
	db.GetDb().Migrator().CreateTable(&_oauth2Entity.Passport{})
}

func playerMigrate(db databases.Database) {
	db.GetDb().Migrator().CreateTable(&_playerEntity.Player{})
}
