package middleware

import (
	"blog/controllers"
	"blog/model"
	"time"

	err_code "blog/errcode"

	"github.com/gin-gonic/gin"
)

func ValidateToken(ctx *gin.Context) {
	tokenStr := ctx.GetHeader("token")
	if tokenStr == "" {
		controllers.ReturnError(ctx, err_code.ErrInvalidRequest, "user not login")
		ctx.Abort()
		return
	}

	token, err := model.GetTokenInfo(tokenStr)
	if err != nil {
		controllers.ReturnError(ctx, err_code.ErrInvalidRequest, err.Error())
		ctx.Abort()
		return
	}

	if token.Expire < time.Now().Unix() {
		controllers.ReturnError(ctx, err_code.ErrInvalidToken, "token expired")
		ctx.Abort()
		return
	}

	if token.State == model.Invalid {
		controllers.ReturnError(ctx, err_code.ErrInvalidToken, "token expired")
		ctx.Abort()
		return
	}

	ctx.Next()
}
