package api

import (
	"fmt"

	"github.com/1996Paul-Wen/SafetyBox/api/handler"
	"github.com/1996Paul-Wen/SafetyBox/config"
	logmanager "github.com/1996Paul-Wen/SafetyBox/infrastructure/log_manager"
	log "github.com/InVisionApp/go-logger"
	"github.com/gin-gonic/gin"
)

// apiLogger is the logger handle on api layer. log.Logger is an interface
var apiLogger log.Logger

// Start the web server
func StartWebServer() error {

	// 指定当前层次使用的日志
	apiLogger = logmanager.DefaultLogManager().GatewayLog()

	app := gin.Default()
	app.RemoveExtraSlash = true
	// 初始化api层路由
	apiVersion := "1.0"
	initRouter(app, apiVersion)

	// listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	globalConf := config.GlobalConfig()
	address := fmt.Sprintf(":%d", globalConf.WebSettings.Port)
	fmt.Printf("start on address:  %s \n", address)
	return app.Run(address)
}

func initRouter(app *gin.Engine, apiVersion string) {
	gConf := config.GlobalConfig()
	baseHandler := handler.NewBaseHandler(
		apiVersion,
		apiLogger,
		handler.WithLimitSettings(
			gConf.WebSettings.LimitSettings.Limit, gConf.WebSettings.LimitSettings.Burst,
		))
	// 使用baseHandler提供的全局中间件
	app.Use(baseHandler.GlobalMiddlewares()...)

	app.GET("/ping", baseHandler.Pong)

	// 注册用于处理业务逻辑的 web handler
	userHandler := handler.NewUserHandler(baseHandler)
	safetyDataHandler := handler.NewSafetyDataHandler(baseHandler)
	handler.RegisterBusinessHandler(userHandler, safetyDataHandler)

	// 需要用户身份验证的路由
	apiGroupWithUserVerification := app.Group("/api", userHandler.VerifyUser)
	// 无需用户身份验证的路由
	apiGroupWithoutUserVerification := app.Group("/api")
	{
		for _, bsHandler := range handler.TotalBusinessHandlers() {
			bsHandler.RegisterUserRequiredRoutersTo(apiGroupWithUserVerification)
			bsHandler.RegisterNoUserRequiredRoutersTo(apiGroupWithoutUserVerification)
		}
	}

}
