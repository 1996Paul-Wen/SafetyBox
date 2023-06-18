package safetydatarepo

import "github.com/1996Paul-Wen/SafetyBox/model"

type Filter struct {
	ID          uint   `json:"id"`
	ArchiveKey  string `json:"archive_key"` // like
	Description string `json:"description"` // like
	UserID      uint   `json:"user_id"`
}

type SafetyDataRepo interface {
	List(filter Filter) ([]model.SafetyData, error)
	InsertOne(safetyData model.SafetyData) (model.SafetyData, error)
	Update(safetyData model.SafetyData, filter Filter) error
}
