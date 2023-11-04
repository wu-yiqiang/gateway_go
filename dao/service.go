package dao

type ServicesDetail struct {
	Info          *ServicesInfo  `json:"info"`
	Http          *HttpRule      `json:"http_rule"`
	Tcp           *TcpRule       `json:"tcp_rule"`
	Grpc          *GrpcRule      `json:"grpc_rule"`
	LoadBalance   *LoadBalance   `json:"load_balance"`
	AccessControl *AccessControl `json:"access_control"`
}
