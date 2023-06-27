package handler

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/1996Paul-Wen/SafetyBox/api/proto"
	"github.com/1996Paul-Wen/SafetyBox/infrastructure/constant"
	"github.com/1996Paul-Wen/SafetyBox/util/ratelimit"
	log "github.com/InVisionApp/go-logger"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	jsoniter "github.com/json-iterator/go"
	uuid "github.com/satori/go.uuid"
)

var myjson jsoniter.API
var reqValidator *validator.Validate

func init() {
	myjson = jsoniter.Config{
		EscapeHTML:    true,
		CaseSensitive: true, // 配置大小写敏感
	}.Froze()
	reqValidator = validator.New()
}

// BaseHandler 负责基本的通用能力，与业务逻辑无关。如请求体解析、返回response消息体、链路标记、限流等能力
type BaseHandler struct {
	APIVersion string
	log.Logger
	reqLimit float64
	reqBurst float64
}

func NewBaseHandler(apiVersion string, logger log.Logger, opt ...BaseHandlerOption) *BaseHandler {
	bh := &BaseHandler{
		APIVersion: apiVersion,
		Logger:     logger,
		reqLimit:   10,
		reqBurst:   10,
	}
	for _, o := range opt {
		o(bh)
	}
	ratelimit.RefreshLimiters(bh.reqLimit, bh.reqBurst)
	return bh
}

type BaseHandlerOption func(*BaseHandler)

func WithLimitSettings(reqLimit, reqBurst float64) BaseHandlerOption {
	return func(bh *BaseHandler) {
		if reqLimit > 0 {
			bh.reqLimit = reqLimit
		}
		if reqBurst > 0 {
			bh.reqBurst = reqBurst
		}
	}
}

// GetVersion returns version
func (h *BaseHandler) GetAPIVersion() string {
	return h.APIVersion
}

// UnmarshalPost unmarshal struct from Post Content to data v
func (h *BaseHandler) UnmarshalPost(c *gin.Context, v interface{}) error {
	var err error
	// 支持读c.Request.Body多次
	bodyBytes, _ := ioutil.ReadAll(c.Request.Body)
	err = myjson.NewDecoder(bytes.NewReader(bodyBytes)).Decode(v)
	if err != nil {
		return fmt.Errorf("Post Data Should Be JSON Map Format : %w", err)
	}
	// 回写body
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

	return reqValidator.Struct(v)
}

// HandleSuccessResponse 发送成功数据
func (h *BaseHandler) HandleSuccessResponse(c *gin.Context, data interface{}) {

	rsp := proto.GeneralResponse{
		Code:    CodeSuccess,
		Data:    data,
		TraceID: c.Request.Context().Value(constant.BasicContextKeys.TraceID).(string),
	}
	c.AbortWithStatusJSON(http.StatusOK, rsp)
}

// HandleFailedResponse 发送失败数据
func (h *BaseHandler) HandleFailedResponse(c *gin.Context, errCode int, err error) {

	rsp := proto.GeneralResponse{
		Code:    errCode,
		Message: err.Error(),
		TraceID: c.Request.Context().Value(constant.BasicContextKeys.TraceID).(string),
	}
	c.AbortWithStatusJSON(http.StatusOK, rsp)
}

func (h *BaseHandler) GlobalMiddlewares() []gin.HandlerFunc {
	return []gin.HandlerFunc{
		h.SetTraceID, h.SetLoginUser, h.ReqRateLimit,
	}
}

// SetTraceID 添加 TraceID 到 Context
func (h *BaseHandler) SetTraceID(c *gin.Context) {
	startTime := time.Now()
	u4 := uuid.NewV4()
	traceID := u4.String()

	// generate a new context with traceID
	ctx := context.WithValue(c.Request.Context(), constant.BasicContextKeys.TraceID, traceID)
	c.Request = c.Request.WithContext(ctx)

	msg := fmt.Sprintf("new request coming, set traceID : %s, request path : %s , method : %s ",
		traceID, c.Request.URL.Path, c.Request.Method)
	h.Info(msg)

	c.Next()

	finishTime := time.Now()
	duration := finishTime.Sub(startTime)
	msg = fmt.Sprintf("finish request, traceID : %s, duration : %s ", traceID, duration.String())
	h.Info(msg)
}

func (h *BaseHandler) SetLoginUser(c *gin.Context) {
	header := c.Request.Header
	// unauthorizedMsg := "用户名或密码不完整"

	userName, nameOK := header["Username"]
	passwd, passwdOK := header["Password"]
	if !nameOK || len(userName) == 0 {
		c.Set(GinContextKeys.LoginUser, AnonymousUser.Name)
	} else {
		c.Set(GinContextKeys.LoginUser, userName[0])
	}
	if !passwdOK || len(passwd) == 0 {
		c.Set(GinContextKeys.Password, AnonymousUser.Password)
	} else {
		c.Set(GinContextKeys.Password, passwd[0])
	}
}

// ReqRateLimit 用户请求限流中间件
func (h *BaseHandler) ReqRateLimit(c *gin.Context) {
	if loginUser, ok := c.MustGet(GinContextKeys.LoginUser).(string); ok {
		// savedCtx := c.Request.Context()
		// defer func() {
		// 	c.Request = c.Request.WithContext(savedCtx)
		// }()

		key := loginUser
		limiter := ratelimit.LoadLimiter(key)
		if !limiter.Allow() {
			err := fmt.Errorf("请求太频繁，请稍后再试")
			h.HandleFailedResponse(c, CodeTooManyRequests, err)
			return
		}
		h.Info(fmt.Sprintf("user [%s] passed watchdog", loginUser))
	}
}

// Pong is default response for request ping
func (h *BaseHandler) Pong(c *gin.Context) {
	h.HandleSuccessResponse(c, "pong!")
}
