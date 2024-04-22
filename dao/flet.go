package dao

type Video struct {
	Name     string `json:"name" gorm:"name" description:"名字"`
	Path     string `json:"path" gorm:"path" description:"路径"`
	Types    string `json:"types" gorm:"types" description:"类型"`
	ImgUrl   string `json:"img_url" gorm:"img_url" description:"oss"`
	VideoUrl string `json:"video_url" gorm:"video_url" description:"video url"`
}
