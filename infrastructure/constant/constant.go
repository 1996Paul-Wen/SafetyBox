package constant

// BasicContextKeys are keys set in gin.Context.Request.Context().
// 方便handler向下透传context.Context以携带一些基础信息如TraceID、认证信息、超时设置
var BasicContextKeys = struct {
	TraceID   string
	UserModel string
}{
	TraceID:   "TraceID",
	UserModel: "UserModel",
}
