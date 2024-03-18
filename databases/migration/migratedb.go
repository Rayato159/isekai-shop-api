package main

import (
	"github.com/Rayato159/isekai-shop-api/config"
	"github.com/Rayato159/isekai-shop-api/databases"
	entities "github.com/Rayato159/isekai-shop-api/entities"
)

func main() {
	appConfig := config.GetAppConfig()
	database := databases.NewPostgresDatabase(appConfig.DatabaseConfig)

	uuidMigreate(database)
	playerMigrate(database)
	adminMigrate(database)
	itemMigrate(database)
	balancingMigrate(database)
	inventoryMigrate(database)
	purchasingMigrate(database)
}

func uuidMigreate(db databases.Database) {
	db.GetDb().Raw(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`).Scan(&struct{}{})
}

func playerMigrate(db databases.Database) {
	db.GetDb().Migrator().CreateTable(&entities.Player{})
}

func adminMigrate(db databases.Database) {
	db.GetDb().Migrator().CreateTable(&entities.Admin{})
}

func itemMigrate(db databases.Database) {
	db.GetDb().Migrator().CreateTable(&entities.Item{})
}

func balancingMigrate(db databases.Database) {
	db.GetDb().Migrator().CreateTable(&entities.Balancing{})
}

func inventoryMigrate(db databases.Database) {
	db.GetDb().Migrator().CreateTable(&entities.Inventory{})
}

func purchasingMigrate(db databases.Database) {
	db.GetDb().Migrator().CreateTable(&entities.PurchasingHistory{})
}
