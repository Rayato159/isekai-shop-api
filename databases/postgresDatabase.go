package databases

import (
	"fmt"
	"log"
	"sync"

	"github.com/Rayato159/isekai-shop-api/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type postgresDatabase struct {
	*gorm.DB
}

var (
	postgresDatabaseInstace *postgresDatabase
	once                    sync.Once
)

func NewPostgresDatabase(cfg *config.DatabaseConfig) Database {
	once.Do(func() {
		dsn := fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%d sslmode=%s search_path=%s",
			cfg.Host,
			cfg.User,
			cfg.Password,
			cfg.DBName,
			cfg.Port,
			cfg.SSLMode,
			cfg.Schema,
		)

		conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			errMessage := fmt.Sprintf("failed to connect database: %s", err.Error())
			panic(errMessage)
		}

		log.Printf("Connected to database %s", cfg.DBName)

		postgresDatabaseInstace = &postgresDatabase{conn}
	})

	return postgresDatabaseInstace
}

func (db *postgresDatabase) GetDb() *gorm.DB {
	return db.DB
}
