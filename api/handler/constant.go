package handler

// ContextKeys are keys set in gin.Context
var ContextKeys = struct {
	LoginUser string
	Password  string
	TraceID   string
}{
	LoginUser: "LoginUser",
	Password:  "Password",
	TraceID:   "TraceID",
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
