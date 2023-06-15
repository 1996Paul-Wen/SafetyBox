package model

import (
	"gorm.io/gorm"
)

type SafetyData struct {
	gorm.Model

	ArchiveKey          *string `gorm:"type:varchar(255);index;not null;default:''"` // such as username for xxx website
	EncryptArchiveValue *string `gorm:"type:text;not null"`                          // such password for xxx website 注意：text类型字段不允许有默认值
	Description         *string `gorm:"type:varchar(255);not null;default:''"`
	UserID              uint64
	User                User `json:"user" gorm:"foreignKey:UserID"`

	// ID               uint64    `gorm:"primaryKey"`
	// CreatedAt        time.Time `gorm:"column:created_at;type:TIMESTAMP;default:CURRENT_TIMESTAMP;<-:create"`
	// UpdatedAt        time.Time `gorm:"column:updated_at;type:TIMESTAMP;default:CURRENT_TIMESTAMP  on update current_timestamp"`
}
