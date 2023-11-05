package controllers

import (
	"fmt"
	"gateway_go/common"
	"gateway_go/dao"
	"gateway_go/dto"
	"gateway_go/global"
	"gateway_go/request"
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
// @Param page_no query int true "页码" default(1)
// @Param page_size query int true "页数" default(10)
// @Success 200 {object} response.Response{data=dto.ServicesListOutput} "success"
// @Router /service/service_list [get]
func (ser *servicesController) ServicesList(c *gin.Context) {
	var form dto.ServicesListInput
	if err := c.ShouldBindQuery(&form); err != nil {
		response.ValidateFail(c, request.GetErrorMsg(form, err))
		return
	}
	list, total, err := services.ServicesService.FindList(&form)
	if err != nil {
		response.BusinessFail(c, "查询失败")
		return
	}
	outList := []dto.ServicesListItemOutput{}
	for _, listItem := range list {
		item, err := services.ServicesService.ServiceDetail(&listItem)
		if err != nil {

		}
		serviceAddr := ""
		clusterIp := global.App.Config.Cluster.ClusterIp
		clusterPort := global.App.Config.Cluster.ClusterPort
		clusterSslPort := global.App.Config.Cluster.ClusterSslPort
		if listItem.LoadType == common.LoadTypeHttp && item.Http.RuleType == common.HttpRuleTypePrefixURL && item.Http.NeedHttps == 1 {
			serviceAddr = clusterIp + clusterSslPort + item.Http.Rule
		}
		if listItem.LoadType == common.LoadTypeHttp && item.Http.RuleType == common.HttpRuleTypePrefixURL && item.Http.NeedHttps == 0 {
			serviceAddr = clusterIp + clusterPort + item.Http.Rule
		}
		if listItem.LoadType == common.LoadTypeHttp && item.Http.RuleType == common.HttpRuleTypeDomain {
			serviceAddr = item.Http.Rule
		}
		if listItem.LoadType == common.LoadTypeTcp {
			serviceAddr = fmt.Sprintf("%s:%d", clusterIp, item.Tcp.Port)
		}
		if listItem.LoadType == common.LoadTypeGrpc {
			serviceAddr = fmt.Sprintf("%s:%d", clusterIp, item.Grpc.Port)
		}
		ipList := dao.LoadBalance{ID: item.LoadBalance.ID, ServiceId: item.LoadBalance.ServiceId, CheckMethod: item.LoadBalance.CheckMethod,
			CheckTimeout: item.LoadBalance.CheckTimeout, CheckInterval: item.LoadBalance.CheckInterval, RoundType: item.LoadBalance.RoundType,
			IpList: item.LoadBalance.IpList, WeightList: item.LoadBalance.WeightList, ForbidList: item.LoadBalance.ForbidList,
			UpstreamConnectTimeout: item.LoadBalance.UpstreamConnectTimeout, UpstreamHeaderTimeout: item.LoadBalance.UpstreamHeaderTimeout,
			UpstreamIdleTimeout: item.LoadBalance.UpstreamIdleTimeout, UpstreamMaxIdle: item.LoadBalance.UpstreamMaxIdle,
		}
		iplist := ipList.GetIpListByModel()
		outItem := dto.ServicesListItemOutput{
			ID:          listItem.ID,
			ServiceName: listItem.ServiceName,
			ServiceDesc: listItem.ServiceDesc,
			LoadType:    listItem.LoadType,
			ServiceAddr: serviceAddr,
			Qps:         0,
			Qpd:         0,
			TotalNode:   len(iplist),
		}
		outList = append(outList, outItem)
	}
	out := &dto.ServicesListOutput{
		Total: total,
		List:  outList,
		Info:  form.Info,
	}
	response.Success(c, out)
}

// ListPage godoc
// @Summary 服务删除
// @Description 服务删除
// @Tags 服务管理
// @ID /service/service_delete
// @Accept  json
// @Produce  json
// @Param id query string false "服务ID"
// @Success 200 {object} response.Response{} "success"
// @Router /service/service_delete [get]
func (ser *servicesController) ServicesDelete(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		response.BusinessFail(c, "服务ID不能为空")
		return
	}
	Id, _ := strconv.Atoi(id)
	err := services.ServicesService.ServiceDelete(Id)
	if err != nil {
		response.BusinessFail(c, err.Error())
		return
	}
	response.Success(c, "id为"+id+"的服务删除成功")
	return
}

// ListPage godoc
// @Summary grpc服务新增
// @Description grpc服务新增
// @Tags 服务管理
// @ID /service/service_add_grpc
// @Accept  json
// @Produce  json
// @Param polygon body dto.GrpcServiceInput true "body"
// @Success 200 {object} response.Response{} "success"
// @Router /service/service_add_grpc [post]
func (ser *servicesController) ServicesAddGrpc(c *gin.Context) {
	var form dto.GrpcServiceInput
	if err := c.ShouldBindJSON(&form); err != nil {
		response.ValidateFail(c, request.GetErrorMsg(form, err))
		return
	}
	fmt.Println("asda", form)
}
