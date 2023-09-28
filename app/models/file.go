package models

type File struct {
	ID
	FileName   string `json:"file_name" gorm:"size:400;not null;comment:文件名"`
	FileHash   string `json:"file_hash" gorm:"size:400;not null;comment:文件哈希值"`
	UpdateTime int64  `json:"update_time;comment:更新时间"`
	IsDelete
}
