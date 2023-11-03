package dao

import "time"

type ServicesInfo struct {
	ID          int64     `json:"id" grom:"id"`
	ServiceName string    `json:"service_name" grom:"service_name"`
	ServiceDesc string    `json:"service_desc" grom:"service_desc"`
	LoadType    int       `json:"load_type" grom:"load_type"`
	UpdatedAt   time.Time `json:"updated_at" grom:"updated_at"`
	CreatedAt   time.Time `json:"created_at" grom:"created_at"`
	IsDelete    int8      `json:"total_node" grom:"total_node"`
}
