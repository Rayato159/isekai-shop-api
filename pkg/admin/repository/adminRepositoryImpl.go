package repository

import (
	entities "github.com/Rayato159/isekai-shop-api/entities"
	_adminExpception "github.com/Rayato159/isekai-shop-api/pkg/admin/exception"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type adminRepositoryImpl struct {
	db     *gorm.DB
	logger echo.Logger
}

func NewAdminRepositoryImpl(db *gorm.DB, logger echo.Logger) AdminRepository {
	return &adminRepositoryImpl{
		db:     db,
		logger: logger,
	}
}

func (r *adminRepositoryImpl) Creating(adminEntity *entities.Admin) (string, error) {
	tx := r.db.Create(adminEntity)

	if tx.Error != nil {
		r.logger.Errorf("Error inserting player: %s", tx.Error.Error())
		return "", &_adminExpception.AdminCreating{AdminID: adminEntity.ID}
	}

	return adminEntity.ID, nil
}

func (r *adminRepositoryImpl) FindByID(adminID string) (*entities.Admin, error) {
	admin := new(entities.Admin)
	tx := r.db.Where("id = ?", adminID).First(admin)

	if tx.Error != nil {
		r.logger.Errorf("Error finding player: %s", tx.Error.Error())
		return nil, &_adminExpception.AdminNotFound{AdminID: adminID}
	}

	return admin, nil
}