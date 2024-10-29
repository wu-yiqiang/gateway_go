package error

import (
	"gateway_go/types/error"
)

// 用户部分以1020000开头
var (
	UserNotFound = error.ServiceError{1020000, "用户不存在"}
	UserNotPermission = error.ServiceError{1020001, "该用户没有权限"}
	UserIdNotExist = error.ServiceError{1020002, "用户ID不存在"}
	UserIsLocked = error.ServiceError{1020003, "该账号已被锁定"}
)