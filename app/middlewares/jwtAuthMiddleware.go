package middlewares

import (
	"github.com/gin-gonic/gin"
	"rbac/util"
	"strings"
)

// JWTAuthMiddleware 基于JWT的认证中间件
func JWTAuth() func(c *gin.Context) {
	return func(c *gin.Context) {
		utilGin := util.Gin{c}
		// 客户端携带Token有三种方式 1.放在请求头 2.放在请求体 3.放在URI
		// 这里假设Token放在Header的Authorization中，并使用Bearer开头
		// 这里的具体实现方式要依据你的实际业务情况决定
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			utilGin.ParamsError(util.ParseTokenError.Error())
			utilGin.Ctx.Abort()
			return
		}
		// 按空格分割
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			utilGin.ParamsError(util.ParseTokenError.Error())
			utilGin.Ctx.Abort()
			return
		}
		mc, err := util.ParseToken(parts[1])
		if err != nil {
			utilGin.ParamsError(util.ParseTokenError.Error())
			utilGin.Ctx.Abort()
			return
		}
		c.Set("username", mc.Username)
		c.Next()
	}
}
