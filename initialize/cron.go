package initialize

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"gateway_go/global"
	models2 "gateway_go/models"
	"github.com/robfig/cron"
	"io"
	"io/ioutil"
	"os"
	"time"
)

// 多个定时任务设置channel 从channel里面获取

func InitCron() {
	c := cron.New()
	if FileCronError := c.AddFunc("*/* * 24 * * ?", FileCron); FileCronError != nil {
		global.App.Log.Info(FileCronError.Error())
		fmt.Println("err", FileCronError)
		return
	}
	//if LogCronError := c.AddFunc("*/2 * * * * ?", LogCron); LogCronError != nil {
	//	global.App.Log.Info(LogCronError.Error())
	//	// 错误记录日志
	//	return
	//}
	c.Start()
	defer c.Stop()
	select {}
}

func FileCron() {
	// 将文件表数据全部设置为1表示已删除
	err := global.App.DB.Exec("update files set is_delete = 1").Error
	if err != nil {
		global.App.Log.Info(err.Error())
	}
	// 递归查询目录文件
	fileArray := GetFileArray()
	// 计算md5码
	for i := 0; i < len(fileArray); i++ {
		fileHash, err := GetFileHash(fileArray[i])
		if err != nil {
			global.App.Log.Info(err.Error())
			break
		}
		// 更新文件信息
		InsertFileInfo(fileArray[i], fileHash)
	}
}

func GetFileArray() []string {
	const dir string = "./storage/files/"
	var fileArray = make([]string, 0)
	FileInfo, err := ioutil.ReadDir(dir)
	if err != nil {
		global.App.Log.Info(err.Error())
	}
	for _, v := range FileInfo {
		if v.IsDir() {

		} else {
			fileArray = append(fileArray, v.Name())
		}

	}
	return fileArray
}

func GetFileHash(fileName string) (string, error) {
	var dir string = "./storage/files/" + fileName
	file, err := os.Open(dir)
	defer file.Close()
	if err != nil {
		global.App.Log.Info(err.Error())
		return "", err
	}
	hash := sha256.New()
	if _, err := io.Copy(hash, file); err != nil {
		global.App.Log.Info(err.Error())
		return "", err
	}
	sum := hash.Sum(nil)
	return hex.EncodeToString(sum), nil
}

func InsertFileInfo(fileName string, fileHash string) {
	//  首先查询是否存在
	var result = global.App.DB.Where("file_hash = ?", fileHash).Select("id").First(&models2.File{})
	if result.RowsAffected != 0 {
		updateError := global.App.DB.Where("file_hash = ?", fileHash).First(&models2.File{}).Update("is_delete", 0).Error
		if updateError != nil {
			global.App.Log.Info(updateError.Error())
		}
		fmt.Println("err", updateError)
		return
	}
	file := models2.File{FileName: fileName, FileHash: fileHash, UpdateTime: time.Now().Unix(), IsDelete: models2.IsDelete{0}}
	err := global.App.DB.Create(&file).Error
	if err != nil {
		global.App.Log.Info(err.Error())
	}
}

func LogCron() {
	// 清除一个月前的日志内容

}
