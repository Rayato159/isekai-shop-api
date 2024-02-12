package entity

import (
	"time"

	"github.com/google/uuid"
)

type (
	Passport struct {
		ID           uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
		PlayerID     string    `gorm:"type:varchar(64);not null;"`
		AccessToken  string    `gorm:"type:varchar(512);unique;not null;"`
		RefreshToken string    `gorm:"type:varchar(512);unique;not null;"`
		CreatedAt    time.Time `gorm:"not null;autoCreateTime;"`
		UpdatedAt    time.Time `gorm:"not null;autoUpdateTime;"`
	}

	UpdateAccessTokenDto struct {
		AccessToken  string
		RefreshToken string
	}
)
