package repository

import (
	_adminEntity "github.com/Rayato159/isekai-shop-api/domains/admin/entity"
)

type AdminRepository interface {
	InsertAdmin(adminEntity *_adminEntity.Admin) (string, error)
	FindAdminByID(adminID string) (*_adminEntity.Admin, error)
}
