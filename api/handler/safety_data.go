package handler

import (
	"github.com/1996Paul-Wen/SafetyBox/api/proto"
	"github.com/1996Paul-Wen/SafetyBox/infrastructure/db"
	"github.com/1996Paul-Wen/SafetyBox/model"
	safetydatarepo "github.com/1996Paul-Wen/SafetyBox/repository/safety_data_repo"
	"github.com/gin-gonic/gin"
)

var _ BusinessHandler = (*SafetyDataHandler)(nil)

type SafetyDataHandler struct {
	*BaseHandler
}

func NewSafetyDataHandler(bh *BaseHandler) *SafetyDataHandler {
	return &SafetyDataHandler{
		BaseHandler: bh,
	}
}

func (s *SafetyDataHandler) RegisterUserRequiredRoutersTo(rg *gin.RouterGroup) {
	withUserVerification := rg.Group("/SafetyData")
	{
		withUserVerification.POST("List", s.List)
		withUserVerification.POST("Create", s.CreateAchive)
	}
}

func (s *SafetyDataHandler) RegisterNoUserRequiredRoutersTo(rg *gin.RouterGroup) {
}

func (s *SafetyDataHandler) List(c *gin.Context) {
	userDetail := c.MustGet(ContextKeys.UserModel).(model.User)

	var req = proto.ParamToListMyArchive{}
	err := s.UnmarshalPost(c, &req)
	if err != nil {
		s.HandleFailedResponse(c, CodeInvalidQueryParameter, err)
		return
	}

	// 只能查自己的密码
	tx := db.DefaultDBManager().BeginTransaction()
	defer tx.Rollback()
	filter := safetydatarepo.Filter{
		UserID:      userDetail.ID,
		ArchiveKey:  req.ArchiveKey,
		Description: req.Description,
	}
	repoImpl := safetydatarepo.NewSafetyDataRepoImpl(tx)
	safetyData, err := repoImpl.List(filter)
	if err != nil {
		s.HandleFailedResponse(c, CodeProcessDataFailed, err)
		return
	}
	tx.Commit()
	s.HandleSuccessResponse(c, safetyData)
}

func (s *SafetyDataHandler) CreateAchive(c *gin.Context) {
	userDetail := c.MustGet(ContextKeys.UserModel).(model.User)

	var req = proto.ParamToCreateMyNewAchive{}
	err := s.UnmarshalPost(c, &req)
	if err != nil {
		s.HandleFailedResponse(c, CodeInvalidQueryParameter, err)
		return
	}

	newArchive := model.SafetyData{
		UserID:              userDetail.ID,
		ArchiveKey:          req.ArchiveKey,
		EncryptArchiveValue: req.ArchiveValue,
		Description:         req.Description,
	}
	tx := db.DefaultDBManager().BeginTransaction()
	defer tx.Rollback()
	repoImpl := safetydatarepo.NewSafetyDataRepoImpl(tx)
	safetyData, err := repoImpl.InsertOne(newArchive)
	if err != nil {
		s.HandleFailedResponse(c, CodeProcessDataFailed, err)
		return
	}
	tx.Commit()
	s.HandleSuccessResponse(c, safetyData)
}
