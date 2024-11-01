package controllers

import (
	"gateway_go/dto"
	"gateway_go/request"
	"gateway_go/response"
	"gateway_go/services"
	"github.com/gin-gonic/gin"
)

type contactController struct {
}

var ContactController = new(contactController)

// ListPage godoc
// @Summary 通讯模块
// @Description 添加好友
// @Tags 通讯模块
// @ID /contact/add
// @Accept  json
// @Produce  json
// @Security Auth
// @Success 200 {object} response.Response{} "success"
// @Router /contact/add [post]
func (con *contactController) AddFriend(c *gin.Context) {
	userId, idIsExist := c.Get("userId")
	if idIsExist == false {
		response.BusinessFail(c, "主用户不存在")
		return
	}
	var form dto.AddFriend
	if err := c.ShouldBindJSON(&form); err != nil {
		response.ValidateFail(c, request.GetErrorMsg(form, err))
		return
	}
	err, item := services.UserFriendsService.AddFriend(userId.(string), form.UserId)
	if err != nil {
		response.BusinessFail(c, err.Error())
		return
	}
	response.Success(c, item)
	return
}

// ListPage godoc
// @Summary 通讯模块
// @Description 查询好友
// @Tags 通讯模块
// @ID /contact/queryUserFriends
// @Accept  json
// @Produce  json
// @Security Auth
// @Success 200 {object} response.Response{data=dto.AdminInfoOutput} "success"
// @Router /contact/queryUserFriends [post]
func (admin *contactController) QueryUserFriends(c *gin.Context) {
	userId, idIsExist := c.Get("userId")
	if idIsExist == false {
		response.BusinessFail(c, "主用户不存在")
		return
	}
	err, items := services.UserFriendsService.QueryFriends(userId.(string))
	if err != nil {
		response.BusinessFail(c, err.Error())
		return
	}
	// 用过friendId 查询frinends
	var lists []dto.AdminInfoOutput
	for _, i2 := range items {
		err, list := services.UserService.UserInfo(i2)
		if err == nil {
			lists = append(lists, list)
		}
	}
	response.Success(c, lists)
	return
}
