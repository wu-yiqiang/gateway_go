package v1

import (
	"fmt"
	_ "gateway_go/chat-room/common/response"
	_ "io/ioutil"
	_ "net/http"
	_ "strings"

	_ "gateway_go/config"
	_ "gateway_go/internal/service"
	//"gateway_go/common/response"
	//"gateway_go/global/log"

	"github.com/gin-gonic/gin"
	_ "github.com/google/uuid"
)

// 前端通过文件名称获取文件流，显示文件
func GetFile(c *gin.Context) {
	fmt.Println("fileeeee")
	//fileName := c.Param("fileName")
	//// log.Logger.Info(fileName)
	//data, _ := ioutil.ReadFile(config.GetConfig().StaticPath.FilePath + fileName)
	//c.Writer.Write(data)
}

// 上传头像等文件
func SaveFile(c *gin.Context) {
	//namePreffix := uuid.New().String()
	//userUuid := c.PostForm("uuid")
	//file, _ := c.FormFile("file")
	//fileName := file.Filename
	//index := strings.LastIndex(fileName, ".")
	//suffix := fileName[index:]
	//newFileName := namePreffix + suffix
	//c.SaveUploadedFile(file, global.App.Config.Storage.Disks.LocalStorage.RootFileDir)
	//err := service.UserService.ModifyUserAvatar(newFileName, userUuid)
	//if err != nil {
	//	c.JSON(http.StatusOK, response.FailMsg(err.Error()))
	//}
	//c.JSON(http.StatusOK, response.SuccessMsg(newFileName))
}
