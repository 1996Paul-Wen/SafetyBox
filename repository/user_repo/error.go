package userrepo

import "fmt"

var (
	ErrUserNotFound              = fmt.Errorf("user not found")
	ErrUserExists                = fmt.Errorf("user already exists")
	ErrUserNameOrPassWordInvalid = fmt.Errorf("username or password invalid")
)
