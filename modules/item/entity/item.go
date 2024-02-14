package entity

type Item struct {
	ID          string `gorm:"primaryKey;autoIncrement"`
	AdminID     string `gorm:"type:varchar(64);"`
	Name        string `gorm:"type:varchar(64);unique;not null;"`
	Description string `gorm:"type:varchar(256);not null;"`
	Picture     string `gorm:"type:varchar(256);not null;"`
	Price       int    `gorm:"not null;"`
	IsArchive   bool   `gorm:"not null;default:false;"`
	CreatedAt   string `gorm:"not null;autoCreateTime;"`
	UpdatedAt   string `gorm:"not null;autoUpdateTime;"`
}
