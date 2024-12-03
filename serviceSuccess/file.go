package serviceSuccess

import (
	"gateway_go/types"
)

// 文件部分以2030000开头
var (
	FileUploadSuccess = types.ServiceError{2030000, "文件上传成功"}
)