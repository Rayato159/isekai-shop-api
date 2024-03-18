package repository

import (
	entities "github.com/Rayato159/isekai-shop-api/entities"
)

type AdminRepository interface {
	InsertAdmin(adminEntity *entities.Admin) (string, error)
	FindAdminByID(adminID string) (*entities.Admin, error)
}
