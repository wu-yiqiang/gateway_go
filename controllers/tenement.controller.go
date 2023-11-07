package controllers

import (
	"gateway_go/dto"
	"gateway_go/request"
	"gateway_go/response"
	"gateway_go/services"
	"github.com/gin-gonic/gin"
	"strconv"
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
// @Security Auth
// @Param info query string false "租户名"
// @Param page_no query int true "页码" default(1)
// @Param page_size query int true "页数" default(10)
// @Success 200 {object} response.Response{data=dto.TenementListOutput} "success"
// @Router /app/app_list [get]
func (ten *tenementController) TenementList(c *gin.Context) {
	var form dto.TenementListInput
	if err := c.ShouldBindQuery(&form); err != nil {
		response.ValidateFail(c, request.GetErrorMsg(form, err))
		return
	}
	list, total, err := services.TenementService.FindTenementList(form)
	if err != nil {
		response.BusinessFail(c, error.Error(err))
		return
	}
	data := dto.TenementListOutput{
		Total: total,
		List:  &list,
		Info:  form.Info,
	}
	response.Success(c, data)
	return
}

// ListPage godoc
// @Summary 租户删除
// @Description 租户删除
// @Tags 租户管理
// @ID /app/app_delete
// @Accept  json
// @Produce  json
// @Security Auth
// @Param id query int false "租户名ID"
// @Success 200 {object} response.Response{} "success"
// @Router /app/app_delete [get]
func (ten *tenementController) TenementDelete(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		response.BusinessFail(c, "租户ID不能为空")
		return
	}
	Id, _ := strconv.Atoi(id)
	err := services.TenementService.TenementDelete(Id)
	if err != nil {
		response.BusinessFail(c, error.Error(err))
		return
	}
	response.Success(c, "id为"+id+"的租户删除成功")
	return
}
