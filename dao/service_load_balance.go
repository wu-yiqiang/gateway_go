package dao

import "strings"

type LoadBalance struct {
	ID                     int64  `json:"id" grom:"id"`
	ServiceId              int64  `json:"service_id" grom:"service_id"`
	CheckMethod            int64  `json:"check_method" grom:"check_method"`
	CheckTimeout           int64  `json:"check_timeout" grom:"check_timeout"`
	CheckInterval          int64  `json:"check_interval" grom:"check_interval"`
	RoundType              int64  `json:"round_type" grom:"round_type"`
	IpList                 string `json:"ip_list" grom:"ip_list"`
	WeightList             int64  `json:"weight_list" grom:"weight_list"`
	ForbidList             string `json:"forbid_list" grom:"forbid_list"`
	UpstreamConnectTimeout int64  `json:"upstream_connect_timeout" grom:"upstream_connect_timeout"`
	UpstreamHeaderTimeout  int64  `json:"upstream_header_timeout" grom:"upstream_header_timeout"`
	UpstreamIdleTimeout    int64  `json:"upstream_idle_timeout" grom:"upstream_idle_timeout"`
	UpstreamMaxIdle        int64  `json:"upstream_max_idle" grom:"upstream_max_idle"`
}

func (l *LoadBalance) GetIpListByModel() []string {
	return strings.Split(l.IpList, ",")
}
