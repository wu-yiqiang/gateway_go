package serviceSuccess

import (
	"gateway_go/types"
)

// 用户部分以2020000开头
var (
	UserNotFound = types.ServiceError{2020000, "用户注册成功"}
	UserNotPermission = types.ServiceError{2020001, "用户登陆成功"}
)