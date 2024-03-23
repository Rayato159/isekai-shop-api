package main

import (
	"github.com/Rayato159/isekai-shop-api/config"
	"github.com/Rayato159/isekai-shop-api/databases"
	entities "github.com/Rayato159/isekai-shop-api/entities"
)

func main() {
	conf := config.ConfigGetting()
	db := databases.NewPostgresDatabase(conf.DatabaseConfig)

	playerMigration(db)
	adminMigration(db)
	itemMigration(db)
	playerCoinMigration(db)
	inventoryMigration(db)
	purchaseHistoryMigration(db)
}

func playerMigration(db databases.Database) {
	db.ConnectionGetting().Migrator().CreateTable(&entities.Player{})
}

func adminMigration(db databases.Database) {
	db.ConnectionGetting().Migrator().CreateTable(&entities.Admin{})
}

func itemMigration(db databases.Database) {
	db.ConnectionGetting().Migrator().CreateTable(&entities.Item{})
}

func playerCoinMigration(db databases.Database) {
	db.ConnectionGetting().Migrator().CreateTable(&entities.PlayerCoin{})
}

func inventoryMigration(db databases.Database) {
	db.ConnectionGetting().Migrator().CreateTable(&entities.Inventory{})
}

func purchaseHistoryMigration(db databases.Database) {
	db.ConnectionGetting().Migrator().CreateTable(&entities.PurchaseHistory{})
}
