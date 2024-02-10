package databases

import "gorm.io/gorm"

type Database interface {
	GetDb() *gorm.DB
}
