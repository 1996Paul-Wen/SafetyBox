package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	// gorm.Model
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	Name             *string `json:"user_name" gorm:"unique;not null;default:''"`
	SaltPasswordHash *string `gorm:"type:varchar(255);not null;default:''"`
	EncryptedSalt    *string `gorm:"type:varchar(255);not null;default:''"`
	PublicKey        *string `gorm:"type:varchar(1024);not null;default:''"` // PublicKey is used to encrypt user archive, which saved in tb `safety_data`

	// CreatedAt        time.Time `gorm:"column:created_at;type:TIMESTAMP;default:CURRENT_TIMESTAMP;<-:create"`
	// UpdatedAt        time.Time `gorm:"column:updated_at;type:TIMESTAMP;default:CURRENT_TIMESTAMP  on update current_timestamp"`
}
