package error

import (
	"gateway_go/types/error"
)

// 公共部分以1010001开头
var (
	IdIsNotExist = error.ServiceError{1010000, "ID不存在"}
)