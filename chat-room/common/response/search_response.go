package response

import "gateway_go/internal/model"

type SearchResponse struct {
	User  model.User  `json:"user"`
	Group model.Group `json:"group"`
}
