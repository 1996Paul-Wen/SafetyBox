package handler

import "github.com/gin-gonic/gin"

// 所有BusinessHandler需要注册到此
var totalBusinessHandlers = []BusinessHandler{}

type BusinessHandler interface {
	RegisterUserRequiredRoutersTo(rg *gin.RouterGroup)
	RegisterNoUserRequiredRoutersTo(rg *gin.RouterGroup)
}

func RegisterBusinessHandler(handler ...BusinessHandler) {
	totalBusinessHandlers = append(totalBusinessHandlers, handler...)
}

func TotalBusinessHandlers() []BusinessHandler {
	return totalBusinessHandlers
}
