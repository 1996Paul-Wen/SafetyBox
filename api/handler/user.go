package handler

import (
	"github.com/1996Paul-Wen/SafetyBox/infrastructure/db"
	userrepo "github.com/1996Paul-Wen/SafetyBox/repository/user_repo"
	"github.com/gin-gonic/gin"
)

var _ BusinessHandler = (*UserHandler)(nil)

type UserHandler struct {
	*BaseHandler
}

func NewUserHandler(bh *BaseHandler) *UserHandler {
	return &UserHandler{
		BaseHandler: bh,
	}
}

func (u *UserHandler) RegisterUserRequiredRoutersTo(rg *gin.RouterGroup) {
	userWithVerification := rg.Group("/User")
	{
		userWithVerification.POST("Describe", u.DescribeUser)
	}
}

func (u *UserHandler) RegisterNoUserRequiredRoutersTo(rg *gin.RouterGroup) {
	userWithoutVerification := rg.Group("/User")
	{
		userWithoutVerification.POST("Create", u.CreateUser)
	}
}

func (u *UserHandler) VerifyUser(c *gin.Context) {
	name := c.MustGet(ContextKeys.LoginUser).(string)
	password := c.MustGet(ContextKeys.Password).(string)

	tx := db.DefaultDBManager().BeginTransaction()
	defer tx.Rollback()
	userRepo := userrepo.NewUserRepoImpl(tx)
	userDetail, err := userRepo.VerifyUser(name, password)
	if err != nil {
		u.HandleFailedResponse(c, CodeUnauthorized, err)
		return
	}
	tx.Commit()

	c.Set(ContextKeys.UserModel, *&userDetail)
}

func (u *UserHandler) DescribeUser(c *gin.Context) {
	var err error

	var req = userrepo.UserIDCard{}
	err = u.UnmarshalPost(c, &req)
	if err != nil {
		u.HandleFailedResponse(c, CodeInvalidQueryParameter, err)
		return
	}

	tx := db.DefaultDBManager().BeginTransaction()
	defer tx.Rollback()
	userRepo := userrepo.NewUserRepoImpl(tx)
	user, err := userRepo.DescribeUser(req)
	if err != nil {
		u.HandleFailedResponse(c, CodeProcessDataFailed, err)
		return
	}
	tx.Commit()
	u.HandleSuccessResponse(c, user)
}

func (u *UserHandler) CreateUser(c *gin.Context) {
	var err error
	var req = userrepo.ParamCreateUser{}
	err = u.UnmarshalPost(c, &req)
	if err != nil {
		u.HandleFailedResponse(c, CodeInvalidQueryParameter, err)
		return
	}

	tx := db.DefaultDBManager().BeginTransaction()
	defer tx.Rollback()
	userRepo := userrepo.NewUserRepoImpl(tx)
	user, err := userRepo.CreateUser(req)
	if err != nil {
		u.HandleFailedResponse(c, CodeProcessDataFailed, err)
		return
	}
	tx.Commit()
	u.HandleSuccessResponse(c, user)
}
