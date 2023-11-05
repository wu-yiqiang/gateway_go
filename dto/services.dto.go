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

type ServicesGrpcInput struct {
	ServiceName       string `json:"service_name" form:"service_name" comment:"服务名称" validate:"required, valid_service_name"`
	ServiceDesc       string `json:"service_desc" form:"service_desc" comment:"服务描述" validate:"required"`
	Port              int    `json:"port" form:"port" comment:"端口，8001-8999" validate:"required,min=8001,max=8999"`
	HeaderTransfor    string `json:"header_transfor" form:"header_transfor" comment:"metadata转换" validate:"valid_header_transfor"`
	OpenAuth          int    `json:"open_auth" form:"open_auth" comment:"是否开启权限" validate:""`
	BlackList         string `json:"black_list" form:"black_list" comment:"黑名单主机，以逗号隔开，白名单优先级高于黑名单" validate:"valid_iplist"`
	WhiteList         string `json:"white_list" form:"white_list" comment:"白名单主机，以逗号隔开，白名单优先级高于黑名单" validate:"valid_iplist"`
	WhiteHostName     string `json:"white_host_name" form:"white_host_name" comment:"白名单主机，以逗号隔开" validate:"valid_iplist"`
	ClientIPFlowLimit int    `json:"clientip_flow_limit" form:"clientip_flow_limit" comment:"客户端限流" validate:""`
	ServiceFlowLimit  int    `json:"service_flow_limit" form:"service_flow_limit" comment:"服务端限流" validate:""`
	RoundType         int    `json:"round_type" form:"round_type" comment:"轮询策略" validate:""`
	IpList            string `json:"ip_list" form:"ip_list" comment:"ip列表" validate:"required,valid_ipportlist"`
	WeightList        string `json:"weight_list" form:"weight_list" comment:"权重列表" validate:"required,valid_iplist"`
	ForbidList        string `json:"forbid_list" form:"forbid_list" comment:"禁用IP列表" validate:"valid_iplist"`
}

func (grpcService ServicesGrpcInput) GetMessages() request.ValidatorMessages {
	return request.ValidatorMessages{
		"username.required":    "用户名不能为空",
		"password.required":    "旧密码不能为空",
		"password.password":    "旧密码必须8-16位，必须包含有一个大写字母，一个小写字母，一个数字",
		"newpassword.required": "新密码不能为空",
		"newpassword.password": "新密码必须8-16位，必须包含有一个大写字母，一个小写字母，一个数字",
	}
}
