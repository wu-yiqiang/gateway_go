package dao

type Tenement struct {
	Id       int    `json:"id" gorm:"primary_key" description:"自增主键"`
	AppId    string `json:"app_id" gorm:"column:app_id" description:"租户名称"`
	Name     string `json:"name" gorm:"column:name" description:"租户名称"`
	Secret   string `json:"secret" gorm:"column:secret" description:"密钥"`
	WhiteIps string `json:"white_ips" gorm:"column:white_ips" description:"白名单"`
	Qpd      int64  `json:"qpd" gorm:"column:qpd" description:"qpd"`
	Qps      int64  `json:"qps" gorm:"column:qps" description:"qps"`
}

func (t *Tenement) TableName() string {
	return "gateway_app"
}
