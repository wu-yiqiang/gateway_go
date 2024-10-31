package error

import (
	"gateway_go/types"
)

// 用户部分以1020000开头
var (
	UserNotFound = types.ServiceError{1020000, "用户不存在"}
	UserNotPermission = types.ServiceError{1020001, "该用户没有权限"}
	UserIdNotExist = types.ServiceError{1020002, "用户ID不存在"}
	UserInfoNotExist = types.ServiceError{1020003, "用户信息不存在"}
	UserIsLocked = types.ServiceError{1020004, "该账号已被锁定"}
	UserLogoutFail = types.ServiceError{1020005, "用户注销失败"}
)