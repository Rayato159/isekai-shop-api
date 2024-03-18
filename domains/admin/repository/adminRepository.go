package repository

import (
	entities "github.com/Rayato159/isekai-shop-api/entities"
)

type AdminRepository interface {
	Creating(adminEntity *entities.Admin) (string, error)
	FindByID(adminID string) (*entities.Admin, error)
}
