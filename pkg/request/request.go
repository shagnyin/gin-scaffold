package request

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Resp struct {
	ErrCode int         `json:"err_code" binding:"required"` // 状态码 0 success 1 failed
	ErrMsg  string      `json:"err_msg" binding:"required"`  // 错误信息描述
	Data    interface{} `json:"data,omitempty"`              // 正常返回数据存放位置
}

func Success(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, Resp{
		ErrCode: 0,
		ErrMsg:  "success",
		Data:    data,
	})
}
func Fail(ctx *gin.Context, httpCode, code int, errMsg string) {
	ctx.JSON(httpCode, Resp{
		ErrCode: code,
		ErrMsg:  errMsg,
	})
}
