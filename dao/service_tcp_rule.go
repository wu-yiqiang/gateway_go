package dao

type TcpRule struct {
	ID        int64 `json:"id" grom:"id"`
	ServiceId int64 `json:"service_id" grom:"service_id"`
	Port      int64 `json:"port" grom:"port"`
}
