package error

import (
	"gateway_go/types"
)

// 文件部分以1030000开头
var (
	FileNotFound = types.ServiceError{1030000, "文件不存在"}
	FileCreateFail = types.ServiceError{1030001, "创建文件失"}
	FileOpenFail = types.ServiceError{1030002, "文件打开失败"}
	FileWriteFail = types.ServiceError{1030003, "文件写入失败"}
	FileUploadFail = types.ServiceError{1030004, "文件上传失败"}
	FileMergeFail = types.ServiceError{1030005, "文件合并失败"}
	FileIsNotEmpty = types.ServiceError{1030006, "文件不能为空"}
	FileNameIsNotEmpty = types.ServiceError{1030007, "文件名不能为空"}
	FileTypeIsNotEmpty = types.ServiceError{1030008, "文件类型不能为空"}
	FileAlreadyExists = types.ServiceError{1030009, "文件已存在"}
	FileHashIsNotEmpty = types.ServiceError{1030010, "文件Hash不能为空"}
)