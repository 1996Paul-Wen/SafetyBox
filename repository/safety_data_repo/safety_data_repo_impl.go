package safetydatarepo

import (
	"fmt"

	"github.com/1996Paul-Wen/SafetyBox/model"
	"github.com/1996Paul-Wen/SafetyBox/util/security"
	"gorm.io/gorm"
)

var _ SafetyDataRepo = (*SafetyDataRepoImpl)(nil)

type SafetyDataRepoImpl struct {
	session *gorm.DB
}

func NewSafetyDataRepoImpl(session *gorm.DB) *SafetyDataRepoImpl {
	return &SafetyDataRepoImpl{
		session: session,
	}
}

func (s *SafetyDataRepoImpl) List(filter Filter) (resp []model.SafetyData, err error) {
	session := s.session
	// gorm `belongs to` 的关联查询 需要 Preload 关联模型
	session = session.Preload("User")
	if filter.ID > 0 {
		session = session.Where("id = ?", filter.ID)
	}
	if filter.Description != "" {
		session = session.Where("description LIKE ?", fmt.Sprintf("%%%s%%", filter.Description))
	}
	if filter.ArchiveKey != "" {
		session = session.Where("archive_key LIKE ?", fmt.Sprintf("%%%s%%", filter.ArchiveKey))
	}
	if filter.UserID > 0 {
		session = session.Where("user_id = ?", filter.UserID)
	}

	result := session.Find(&resp)
	return resp, result.Error
}

func (s *SafetyDataRepoImpl) InsertOne(safetyData model.SafetyData) (resp model.SafetyData, err error) {
	user := model.User{ID: safetyData.UserID}
	result := s.session.First(&user)
	if result.Error != nil {
		return resp, result.Error
	}
	encryptArchiveValue, err := security.RSAStringEncrypt(*user.PublicKey, *safetyData.EncryptArchiveValue)
	if err != nil {
		return resp, err
	}
	safetyData.EncryptArchiveValue = &encryptArchiveValue
	result = s.session.Create(&safetyData)
	return safetyData, result.Error
}

func (s *SafetyDataRepoImpl) Update(safetyData model.SafetyData, filter Filter) error {
	session := s.session.Model(&model.SafetyData{})
	if filter.ID > 0 {
		session = session.Where("id = ?", filter.ID)
	}
	if filter.Description != "" {
		session = session.Where("description LIKE ?", fmt.Sprintf("%%%s%%", filter.Description))
	}
	if filter.ArchiveKey != "" {
		session = session.Where("archive_key LIKE ?", fmt.Sprintf("%%%s%%", filter.ArchiveKey))
	}
	if filter.UserID > 0 {
		session = session.Where("user_id = ?", filter.UserID)
	}
	result := session.Updates(&safetyData)

	return result.Error
}
