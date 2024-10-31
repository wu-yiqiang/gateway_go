package error

import (
	"gateway_go/types"
)

// 公共部分以1010001开头
var (
	IdIsNotExist = types.ServiceError{1010000, "ID不存在"}
)