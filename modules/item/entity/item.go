package entity

import (
	"time"

	_itemModel "github.com/Rayato159/isekai-shop-api/modules/item/model"
)

type (
	Item struct {
		ID          uint64    `gorm:"primaryKey;autoIncrement"`
		AdminID     *string   `gorm:"type:varchar(64);"`
		Name        string    `gorm:"type:varchar(64);unique;not null;"`
		Description string    `gorm:"type:varchar(256);not null;"`
		Picture     string    `gorm:"type:varchar(256);not null;"`
		Price       uint      `gorm:"not null;"`
		IsArchive   bool      `gorm:"not null;default:false;"`
		CreatedAt   time.Time `gorm:"not null;autoCreateTime;"`
		UpdatedAt   time.Time `gorm:"not null;autoUpdateTime;"`
	}

	ItemFilterDto struct {
		Name        string
		Description string
		PaginateDto
	}

	PaginateDto struct {
		Page int64
		Size int64
	}

	UpdateItemDto struct {
		AdminID     *string
		Name        string
		Description string
		Picture     string
		Price       uint
	}
)

func (i *Item) ToItemModel() *_itemModel.Item {
	return &_itemModel.Item{
		ID:          i.ID,
		AdminID:     i.AdminID,
		Name:        i.Name,
		Description: i.Description,
		Picture:     i.Picture,
		Price:       i.Price,
	}
}
