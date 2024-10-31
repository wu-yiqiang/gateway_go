package response

import (
	"gateway_go/global"
	"github.com/gin-gonic/gin"
	"net/http"
	"gateway_go/types"
)

// 响应结构体
type Response struct {
	Code    int         `json:"code"`    // 自定义错误码
	Data    interface{} `json:"data"`    // 数据
	Message string      `json:"message"` // 信息
}

// Success 响应成功  Code为200表示成功
func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Response{
		200,
		data,
		"success",
	})
}

// Fail 响应失败 Code不为200表示失败
func Fail(c *gin.Context, errorCode int, msg string) {
	c.JSON(http.StatusOK, Response{
		errorCode,
		nil,
		msg,
	})
}

// FailByError 失败响应 返回自定义错误的错误码、错误信息
func FailByError(c *gin.Context, error global.CustomError) {
	Fail(c, error.ErrorCode, error.ErrorMsg)
}

// ValidateFail 请求参数验证失败
func ValidateFail(c *gin.Context, msg string) {
	Fail(c, global.Errors.ValidateError.ErrorCode, msg)
}

// BusinessFail 业务逻辑失败
func BusinessFail(c *gin.Context, msg string) {
	Fail(c, global.Errors.BusinessError.ErrorCode, msg)
}

// BusinessFail 业务逻辑失败
func ServiceFail(c *gin.Context, err types.ServiceError) {
	Fail(c, err.ErrorCode, err.ErrorMsg)
}

// 鉴权失败
func TokenFail(c *gin.Context) {
	FailByError(c, global.Errors.TokenError)
}
