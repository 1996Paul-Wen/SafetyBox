package userrepo

import (
	"context"

	"github.com/1996Paul-Wen/SafetyBox/model"
)

type UserIDCard struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type ParamCreateUser struct {
	Name      string `json:"name" validate:"required"`
	PassWord  string `json:"password" validate:"required"`
	PublicKey string `json:"public_key" validate:"required"`
}

type UserRepo interface {
	DescribeUser(context.Context, UserIDCard) (model.User, error)
	CreateUser(context.Context, ParamCreateUser) (uint, error)
	VerifyUser(ctx context.Context, name, password string) (model.User, error)
}
