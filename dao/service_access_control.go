package dao

type AccessControl struct {
	ID                int64  `json:"id" grom:"id"`
	ServiceId         int64  `json:"service_id" grom:"service_id"`
	OpenAuth          int64  `json:"open_auth" grom:"open_auth"`
	BlackList         string `json:"black_list" grom:"black_list"`
	WhiteList         string `json:"white_list" grom:"white_list"`
	WhiteHostName     string `json:"white_host_name" grom:"white_host_name"`
	ClientipFlowLimit int64  `json:"clientip_flow_limit" grom:"clientip_flow_limit"`
	ServiceFlowLimit  int64  `json:"service_flow_limit" grom:"service_flow_limit"`
}
