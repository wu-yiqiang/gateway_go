package controllers

import (
	"fmt"
	"gateway_go/dto"
	"gateway_go/global"
	"gateway_go/request"
	"gateway_go/response"
	"github.com/gin-gonic/gin"
	"io"
	"io/ioutil"
	"os"
	"gateway_go/utils"
	"path/filepath"
	"strings"
)

type fileController struct {
}

var FileController = new(fileController)

// ListPage godoc
// @Summary 文件上传
// @Description 文件上传
// @Tags 文件管理
// @ID /file/upload
// @Accept  json
// @Produce  json
// @Security Auth
// @Accept multipart/form-data
// @Param file formData file true "文件"
// @Success 200 {object} response.Response{} "success"
// @Router /file/upload [post]
func (f *fileController) Upload(c *gin.Context) {
	file, _, err := c.Request.FormFile("file")
	if err != nil {
		response.BusinessFail(c, "文件不能为空")
		return
	}
	filename := c.Request.FormValue("file_name")
	if err != nil {
		response.BusinessFail(c, "filename不能为空")
		return
	}
	filehash := c.Request.FormValue("file_hash")
	fmt.Println("filehash", filehash)
	if err != nil {
		response.BusinessFail(c, "filehash不能为空")
		return
	}
	hash := c.Request.FormValue("hash")
	if err != nil {
		response.BusinessFail(c, "hash不能为空")
		return
	}
	index := strings.Split(hash, "-")[1]
	saveDir := global.App.Config.Storage.Disks.LocalStorage.RootFileDir
	filenameDir := filename + ".dir"
	if index == "0" {
		isexist := isExist(saveDir + filenameDir)
		if isexist == false {
			// 创建文件夹
			err := os.Mkdir(saveDir+filenameDir, os.FileMode(0777))
			if err != nil {
				response.ValidateFail(c, err.Error())
				return
			}
		}
	}

	exist := isExist(saveDir + filenameDir + "/" + hash)
	if exist == true {
		response.Success(c, "文件已存在, 上传成功")
		return
	}
	go func() {
		//写入文件
		out, err := os.Create(saveDir + filenameDir + "/" + hash)
		if err != nil {
			response.BusinessFail(c, "服务错误")
			return
		}
		defer out.Close()
		_, err = io.Copy(out, file)
		if err != nil {
			response.BusinessFail(c, "上传失败")
			return
		}
	}()
	response.Success(c, "上传成功")
}

func (f *fileController) MergeChunks(c *gin.Context) {
	var form dto.FileMergeInput
	if err := c.ShouldBindJSON(&form); err != nil {
		response.ValidateFail(c, request.GetErrorMsg(form, err))
		return
	}
	filename := form.FileName
	saveDir := global.App.Config.Storage.Disks.LocalStorage.RootFileDir
	// 打开之前上传文件
	_, err := os.Create(saveDir + filename)

	if err != nil {
		fmt.Println("创建文件失败", saveDir+filename, err)
		response.BusinessFail(c, "创建文件失败")
		return
	}
	file, err := os.OpenFile(saveDir+filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
	defer file.Close()
	if err != nil {
		response.BusinessFail(c, "打开之前上传文件不存在")
		return
	}
	filenameDir := filename + ".dir"
	part_list, err := filepath.Glob(saveDir + filenameDir + "/*")
	if err != nil {
		response.BusinessFail(c, "需要合并的文件夹出现错误")
		return
	}
	i := 0
	for _, v := range part_list {
		f, err := os.OpenFile(v, os.O_RDONLY, os.ModePerm)
		if err != nil {
			fmt.Println(err)
			return
		}
		b, err := ioutil.ReadAll(f)
		if err != nil {
			fmt.Println(err)
			return
		}
		file.Write(b)
		f.Close()
		i++
	}
	// 删除文件夹
	error := os.RemoveAll(saveDir + filenameDir)
	if err != nil {
		fmt.Println("error", error, filenameDir)
	}
	response.Success(c, "合并成功")
	return
}

// ListPage godoc
// @Summary 文件上传
// @Description 文件上传(无分片上传功能)
// @Tags 文件管理
// @ID /file/uploadFile
// @Accept  json
// @Produce  json
// @Security Auth
// @Accept multipart/form-data
// @Param file formData file true "file"
// @Param filename formData string true "文件名"
// @Param filetype formData string true "文件类型"
// @Success 200 {object} response.Response{} "success"
// @Router /file/uploadFile [post]
func (f *fileController) UploadFile(c *gin.Context) {
	file, _, err := c.Request.FormFile("file")
	filename, fileNameErr := c.GetPostForm("filename")
	filetype, fileTypeErr := c.GetPostForm("filetype")
	if err != nil {
		response.BusinessFail(c, "文件不能为空")
		return
	}
	if fileNameErr == false {
		response.BusinessFail(c, "文件名不能为空")
		return
	}
	if fileTypeErr == false {
		response.BusinessFail(c, "文件类型不能为空")
		return
	}
	var saveDir string
	var url string
	addr, err := utils.NetAddr()
	if filetype == "2" {
		saveDir = global.App.Config.Storage.Disks.LocalStorage.RootImageDir
		url = "http://" + addr + ":" + global.App.Config.App.Port + "/assets/" + filename
	}
	if filetype == "3" {
		saveDir = global.App.Config.Storage.Disks.LocalStorage.RootVideoDir + "upload/"
		url = "http://" + addr + ":" + global.App.Config.App.StaticPort + "/upload/" + filename
	}
	if filetype == "5" {
		saveDir = global.App.Config.Storage.Disks.LocalStorage.RootFileDir
		url = "http://" + addr + "/file/" + filename
	}
	path := make(map[string]string)
	//写入文件
	out, err := os.Create(saveDir + filename)
	if err != nil {
		response.BusinessFail(c, "文件写入失败")
		return
	}
	defer out.Close()
	_, err = io.Copy(out, file)
	if err != nil {
		response.BusinessFail(c, "上传失败")
		return
	}
	if filetype == "3" {
		imgUrl, err := GetSnapshot(global.App.Config.Storage.Disks.LocalStorage.RootVideoDir+"upload/"+filename, global.App.Config.Storage.Disks.LocalStorage.RootImageDir+filepath.Base(filename), filename, 1)
		fmt.Println("sdsd", imgUrl)
		if err != nil {
			response.BusinessFail(c, "获取视频封面失败")
			return
		}
		path["poster"] = imgUrl
	}
	path["path"] = url
	response.Success(c, path)
	return
}

func isExist(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}
