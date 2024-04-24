package hello

import "github.com/gin-gonic/gin"

func HelloHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.String(200, "Hello, World!")
	}
}
