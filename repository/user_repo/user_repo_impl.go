package userrepo

import (
	"errors"
	"fmt"

	"github.com/1996Paul-Wen/SafetyBox/config"
	"github.com/1996Paul-Wen/SafetyBox/model"
	"github.com/1996Paul-Wen/SafetyBox/util/security"
	structvalidator "github.com/1996Paul-Wen/SafetyBox/util/struct_validator"
	"gorm.io/gorm"
)

type UserRepoImpl struct {
	session *gorm.DB
}

func NewUserRepoImpl(session *gorm.DB) *UserRepoImpl {
	return &UserRepoImpl{
		session: session, // db handle
	}
}

func (ur *UserRepoImpl) DescribeUser(u UserIDCard) (user model.User, err error) {
	user = model.User{
		Name: &u.Name,
		ID:   u.ID,
	}
	result := ur.session.First(&user)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return model.User{}, ErrUserNotFound
		}
		return model.User{}, result.Error
	}

	return user, nil
}

func (ur *UserRepoImpl) CreateUser(p ParamCreateUser) (user model.User, err error) {
	err = structvalidator.Validator.Struct(p)
	if err != nil {
		return
	}

	// 生产盐
	salt := security.RandStr(6)
	var aesSalt string
	// aes加密 盐
	aesSalt, err = security.AESEncrypt(salt, config.GlobalConfig().AESKeyForUserLoginPWSalt)
	if err != nil {
		return
	}

	// 密码加盐计算md5
	pwSaltMD5 := security.MD5_SALT(p.PassWord, salt)

	user = model.User{
		Name:             &p.Name,
		PublicKey:        &p.PublicKey,
		EncryptedSalt:    &aesSalt,
		SaltPasswordHash: &pwSaltMD5,
	}
	fmt.Printf("%+v\n", user)
	result := ur.session.Create(&user)
	if result.Error != nil {
		err = result.Error
	}
	return
}

// VerifyUser 检查用户名密码是否正确
func (ur *UserRepoImpl) VerifyUser(name, passwd string) (dbUser model.User, err error) {
	dbUser, err = ur.DescribeUser(UserIDCard{
		Name: name,
	})
	if err != nil {
		return
	}

	// 解密盐
	var salt string
	salt, err = security.AESDecrypt(*dbUser.EncryptedSalt, config.GlobalConfig().AESKeyForUserLoginPWSalt)
	if err != nil {
		return
	}
	// fmt.Printf("salt is %+v\n", salt)

	// 计算用户输入的加盐哈西值
	inputPswSaltMD5 := security.MD5_SALT(passwd, salt)
	if inputPswSaltMD5 != *dbUser.SaltPasswordHash {
		fmt.Printf("input md5: %+v, db md5: %+v\n", inputPswSaltMD5, *dbUser.SaltPasswordHash)
		err = ErrUserNameOrPassWordInvalid
		return
	}

	return
}
