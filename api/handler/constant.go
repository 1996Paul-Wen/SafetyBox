package handler

// GinContextKeys are keys set in gin.Context.
// 注意，GinContextKeys负责在gin.Context中存储数据，仅供handler之间透传使用，
// 而不是在gin.Context.Request.Context()中存储数据
var GinContextKeys = struct {
	LoginUser string
	Password  string
}{
	LoginUser: "LoginUser",
	Password:  "Password",
}

const (
	CodeSuccess int = iota
	CodeInvalidQueryParameter
	CodeProcessDataFailed
	CodeUnauthorized
	CodeForbidden
	CodeTooManyRequests
	CodeEntityNotFound
	CodeInternalServerError
)

// AnonymousUser 给某些不需要用户登录的接口使用。如创建新用户
var AnonymousUser = struct {
	Name     string
	Password string
}{
	Name:     "Anonymous",
	Password: "Anonymous",
}
