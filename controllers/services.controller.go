package controllers

import (
	"gateway_go/dto"
	"gateway_go/response"
	"gateway_go/services"
	"github.com/gin-gonic/gin"
	"strconv"
)

type servicesController struct {
}

var ServicesController = new(servicesController)

// ListPage godoc
// @Summary 服务查询
// @Description 服务查询
// @Tags 服务管理
// @ID /service/service_list
// @Accept  json
// @Produce  json
// @Param info query string false "服务名"
// @Param page_no query int false "页码"
// @Param page_size query int true "页数"
// @Success 200 {object} response.Response{data=dto.ServicesListOutput} "success"
// @Router /service/service_list [get]
func (ser *servicesController) ServicesList(c *gin.Context) {
	info := c.Query("info")
	no := c.Query("page_no")
	size := c.Query("page_size")
	if no == "" || size == "" {
		response.BusinessFail(c, "分页参数不能为空")
		return
	}
	sizeNum, _ := strconv.Atoi(size)
	noNum, _ := strconv.Atoi(no)
	parmas := &dto.ServicesListInput{
		Info:     info,
		PageSize: sizeNum,
		PageNo:   noNum,
	}
	list, total, err := services.ServicesService.FindList(c, parmas)
	if err != nil {
		response.BusinessFail(c, "查询失败")
		return
	}
	out := &dto.ServicesListOutput{
		Total: total,
		List:  list,
		Info:  info,
	}
	response.Success(c, out)
}
