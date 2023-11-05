package dto

import "gateway_go/request"

type ServicesListInput struct {
	Info     string `form:"info" form:"info" description:"关键词" example:"admin"`
	PageNo   int    `form:"page_no" form:"page_no" description:"页码" example:"1"`
	PageSize int    `form:"page_size" form:"page_size" description:"页数" example:"10"`
}

type ServicesListItemOutput struct {
	ID          int64  `json:"id" form:"id"`
	ServiceName string `json:"service_name" form:"service_name"`
	ServiceDesc string `json:"service_desc" form:"service_desc"`
	LoadType    int    `json:"load_type" form:"load_type"`
	ServiceAddr string `json:"service_addr" form:"service_addr"`
	Qps         int64  `json:"qps" form:"qps"`
	Qpd         int64  `json:"qpd" form:"qpd"`
	TotalNode   int    `json:"total_node" form:"total_node"`
}

type ServicesListOutput struct {
	Info  string                   `json:"info" form:"info" description:"关键词" example:"admin"`
	Total int64                    `json:"total" form:"total" description:"总数" example:"400"`
	List  []ServicesListItemOutput `json:"list" form:"list" description:"数据"`
}

type ServicesTcpInput struct {
	ID          int64  `json:"id" form:"id"`
	ServiceName string `json:"service_name" form:"service_name"`
	ServiceDesc string `json:"service_desc" form:"service_desc"`
	LoadType    int    `json:"load_type" form:"load_type"`
	ServiceAddr string `json:"service_addr" form:"service_addr"`
	Qps         int64  `json:"qps" form:"qps"`
	Qpd         int64  `json:"qpd" form:"qpd"`
	TotalNode   int    `json:"total_node" form:"total_node"`
}

type GrpcServiceInput struct {
	ServiceName       string `json:"service_name" form:"service_name" gorm:"service_name" comment:"服务名称" binding:"required,valid_service_name"`
	ServiceDesc       string `json:"service_desc" form:"service_desc" gorm:"service_desc" comment:"服务描述" binding:"required"`
	Port              int    `json:"port" form:"port" gorm:"port" comment:"端口，8001-8999" binding:"required,min=8001,max=8999"`
	HeaderTransfor    string `json:"header_transfor" form:"header_transfor" gorm:"header_transfor" comment:"metadata转换" binding:"valid_header_transfor"`
	OpenAuth          int    `json:"open_auth" form:"open_auth" gorm:"open_auth" comment:"是否开启权限" binding:""`
	BlackList         string `json:"black_list" form:"black_list" gorm:"black_list" comment:"黑名单主机，以逗号隔开，白名单优先级高于黑名单" binding:"valid_iplist"`
	WhiteList         string `json:"white_list" form:"white_list" gorm:"white_list" comment:"白名单主机，以逗号隔开，白名单优先级高于黑名单" binding:"valid_iplist"`
	WhiteHostName     string `json:"white_host_name" form:"white_host_name" gorm:"white_host_name" comment:"白名单主机，以逗号隔开" binding:"valid_iplist"`
	ClientIPFlowLimit int    `json:"clientip_flow_limit" form:"clientip_flow_limit" gorm:"clientip_flow_limit" comment:"客户端限流" binding:""`
	ServiceFlowLimit  int    `json:"service_flow_limit" form:"service_flow_limit" gorm:"service_flow_limit" comment:"服务端限流" binding:""`
	RoundType         int    `json:"round_type" form:"round_type" gorm:"round_type" comment:"轮询策略" binding:""`
	IpList            string `json:"ip_list" form:"ip_list" gorm:"ip_list" comment:"ip列表" binding:"required,valid_ipportlist"`
	WeightList        string `json:"weight_list" form:"weight_list" gorm:"weight_list" comment:"权重列表" binding:"required,valid_iplist"`
	ForbidList        string `json:"forbid_list" form:"forbid_list" gorm:"forbid_list" comment:"禁用IP列表" binding:"valid_iplist"`
}

func (grpcService GrpcServiceInput) GetMessages() request.ValidatorMessages {
	return request.ValidatorMessages{
		"service_name.required":                 "服务名称不能为空",
		"service_name.valid_service_name":       "服务名称不能为空",
		"service_desc.required":                 "服务描述不能为空",
		"port.required":                         "端口不能为空",
		"port.min":                              "端口值不能小于8001",
		"port.max":                              "端口值不能大于8999",
		"header_transfor.valid_header_transfor": "metadata转换",
		"black_list.valid_iplist":               "黑名单",
		"white_list.valid_iplist":               "白名单",
		"white_host_name.valid_iplist":          "白名单主机，以逗号隔开",
		"ip_list.required":                      "ip列表",
		"valid_ipportlist.required":             "ip列表",
		"weight_list.required":                  "权重列表",
		"weight_list.valid_iplist":              "权重列表",
		"forbid_list.required":                  "禁用ip列表",
		"forbid_list.valid_iplist":              "禁用ip列表",
	}
}
