package dto

import "gateway_go/dao"

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
	Info  string             `json:"info" form:"info" description:"关键词" example:"admin"`
	Total int64              `json:"total" form:"total" description:"总数" example:"400"`
	List  []dao.ServicesInfo `json:"list" form:"list" description:"数据"`
}
