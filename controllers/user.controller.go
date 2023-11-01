package controllers

import (
	"gateway_go/common"
	"gateway_go/services"
	"gateway_go/validator"
	"github.com/gin-gonic/gin"
)

type AdminController struct {
}

// ListPage godoc
// @Summary 用户管理
// @Description 用户管理
// @Tags 用户
// @ID /user/register
// @Accept  json
// @Produce  json
// @Param polygon body dto.RegisterInput true "body"
// @Success 200 {object} common.Response{} "success"
// @Router /user/register [post]
func AdminRegister(c *gin.Context) {
	var form validator.Register
	if err := c.ShouldBindJSON(&form); err != nil {
		common.ValidateFail(c, common.GetErrorMsg(form, err))
		return
	}
	if err, _ := services.UserService.Register(form); err != nil {
		common.BusinessFail(c, err.Error())
	} else {
		common.Success(c, "新用户注册成功")
	}
}

func AdminLogin(c *gin.Context) {

}

func AdminChangePassword(c *gin.Context) {

}

func AdminLogout(c *gin.Context) {

}
