package controllers

import (
	"gateway_go/dto"
	"gateway_go/request"
	"gateway_go/response"
	"github.com/gin-gonic/gin"
)

type tenementController struct {
}

var TenementController = new(tenementController)

// ListPage godoc
// @Summary 租户查询
// @Description 租户查询
// @Tags 租户管理
// @ID /app/app_list
// @Accept  json
// @Produce  json
// @Param info query string false "租户名"
// @Param page_no query int true "页码" default(1)
// @Param page_size query int true "页数" default(10)
// @Success 200 {object} response.Response{} "success"
// @Router /app/app_list [get]
func (ten *tenementController) TenementList(c *gin.Context) {
	var form dto.TenementListInput
	if err := c.ShouldBindQuery(&form); err != nil {
		response.ValidateFail(c, request.GetErrorMsg(form, err))
		return
	}

	//
}
