package app

import (
	"fmt"
	"gateway_go/app/common/response"
	"gateway_go/global"
	"gateway_go/utils"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"os"
	"os/exec"
)

func WebSocketHandler(c *gin.Context) {
	// 获取WebSocket连接

	ws, err := websocket.Upgrade(c.Writer, c.Request, nil, 1024, 1024)
	if err != nil {
		return
	}
	utils.InitConection(ws)
}

// 文件上传
func UploadHandler(c *gin.Context) {
	// 文件校验

	// 新建文件夹
	// 上传文件块
	//

}

// 文件合并
func FileMerge(foldName string) {
	// 获取文件夹名
	// 合并该文件夹下的所有文件

}

// 文件下载
func DownloadHandler(c *gin.Context) {
	// 文件校验
	var fileName string
	fileName = c.Query("fileName")
	// 读取文件
	var filePath = "./storage/files/" + fileName
	isExit := FileIsExit(fileName)
	if isExit == false {
		// 文件不存在
		response.ValidateFail(c, "该文件不存在")
		return
	}

	fileTmp, err := os.Open(filePath)
	if err != nil {
		const error = "获取文件失败"
		global.App.Log.Info(error)
		// 文件不存在
		response.ValidateFail(c, error)
		return
	}
	defer fileTmp.Close()
	//获取文件的名称
	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Disposition", "attachment; filename="+fileName)
	c.Header("Content-Transfer-Encoding", "binary")
	c.Header("Cache-Control", "no-cache")
	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Disposition", "attachment; filename="+fileName)
	c.Header("Content-Transfer-Encoding", "binary")
	c.File(filePath)
	return
}

func FileIsExit(fileName string) bool {
	var dir = "./storage/files/"
	var path = dir + fileName
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	fmt.Println("is exit", exec.Command("pwd"))
	if os.IsExist(err) {
		return true
	}
	return false
}

// 构建定时任务扫描目录下的文件信息构建文件表
func time() {

}
