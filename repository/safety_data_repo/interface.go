package safetydatarepo

import (
	"context"

	"github.com/1996Paul-Wen/SafetyBox/model"
)

type Filter struct {
	ID          uint   `json:"id"`
	ArchiveKey  string `json:"archive_key"` // like
	Description string `json:"description"` // like
	UserID      uint   `json:"user_id"`
}

type SafetyDataRepo interface {
	List(ctx context.Context, filter Filter) ([]model.SafetyData, error)
	InsertOne(ctx context.Context, safetyData model.SafetyData) (model.SafetyData, error)
	Update(ctx context.Context, safetyData model.SafetyData, filter Filter) error
}
