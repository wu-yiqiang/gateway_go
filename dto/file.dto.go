package dto

import "gateway_go/request"

type FileUploadInput struct {
	File     []byte `json:"file" form:"file" gorm:"file" comment:"文件" binding:"required"`
	Hash     string `json:"hash" form:"hash" gorm:"hash" comment:"哈希值" default:"sds1223ssd" binding:"required"`
	FileHash string `json:"filehash" form:"filehash" gorm:"filehash" comment:"文件哈希" default:"asda121" binding:"required"`
	FileName string `json:"filename" form:"filename" gorm:"filename" comment:"文件名" default:"文件名" binding:"required"`
}

func (fileUploadInput FileUploadInput) GetMessages() request.ValidatorMessages {
	return request.ValidatorMessages{
		"filename.required": "文件名不能为空",
		"hash.required":     "哈希值不能为空",
		"filehash.required": "文件哈希值不能为空",
		"file.required":     "文件不能为空",
	}
}

type FileMergeInput struct {
	FileName string `json:"filename" form:"filename" gorm:"filename" comment:"文件名" default:"文件名" binding:"required"`
	Size     int64  `json:"size" form:"size" gorm:"size" comment:"文件大小" default:"1234" binding:"required"`
}

func (fileMergeInput FileMergeInput) GetMessages() request.ValidatorMessages {
	return request.ValidatorMessages{
		"filename.required": "文件名不能为空",
		"size.required":     "文件大小不能为空",
	}
}
