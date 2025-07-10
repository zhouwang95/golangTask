package middleware

import (
	"PersonBlog/logger"
	"PersonBlog/response"
	"PersonBlog/utils"
	"github.com/gin-gonic/gin"
)

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		authHeader := context.GetHeader("Authorization")
		if authHeader == "" {
			logger.AddLog(response.Unauthorized, "token不存在", "token 不存在")
			response.Fail(context, response.Unauthorized, "token 不存在")
			context.Abort()
			return
		}

		claim, err := utils.ParseToken(authHeader)
		if err != nil {
			response.Fail(context, 10008, "token 无效或已过期")
			context.Abort()
			return
		}
		context.Set("userName", claim.Issuer)
		context.Next()
	}
}
