package controllers

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"gateway_go/dao"
	"gateway_go/dto"
	"gateway_go/global"
	"gateway_go/request"
	"gateway_go/response"
	"gateway_go/utils"
	"github.com/disintegration/imaging"
	"github.com/gin-gonic/gin"
	ffmpeg "github.com/u2takey/ffmpeg-go"
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

const weatherkey = "4477f6632ae5bc3c2e8f5c47c9772e30"
const weatherUrl = "http://apis.juhe.cn/simpleWeather/query"

const newsKey = "933372cc5c55ce8582e2c8db5e3a44a7"
const newsUrl = "http://v.juhe.cn/toutiao/index"

type fletController struct {
}

var FletController = new(fletController)

// ListPage godoc
// @Summary flutter管理
// @Description 查询天气信息
// @Tags flutter管理
// @ID /flet/weather
// @Accept  json
// @Produce  json
// @Security Auth
// @Param token query string true "token" （时间来不及，该swagger文档入参有问题，需要更新该swagger文档）
// @Router /flet/weather [post]
func (f *fletController) GetWeatherInfo(c *gin.Context) {
	var form dto.FletWeatherInput
	if err := c.ShouldBindQuery(&form); err != nil {
		response.ValidateFail(c, request.GetErrorMsg(form, err))
		return
	}
	// 初始化参数
	param := url.Values{}
	// 接口请求参数
	param.Set("city", form.Location) // 要查询的城市名称/id，城市名称如：温州、上海、北京
	param.Set("key", weatherkey)     // 接口请求Key

	// 发送请求
	datas, err := Get(weatherUrl, param)
	if err != nil {
		// 请求异常，根据自身业务逻辑进行调整修改
		fmt.Errorf("请求异常:\r\n%v", err)
		response.ValidateFail(c, err.Error())
		return
	}
	var netReturn map[string]interface{}
	jsonerr := json.Unmarshal(datas, &netReturn)
	if jsonerr != nil {
		// 解析JSON异常，根据自身业务逻辑进行调整修改
		fmt.Errorf("请求异常:%v", jsonerr)
		response.ValidateFail(c, jsonerr.Error())
		return
	}
	errorCode := netReturn["error_code"]
	reason := netReturn["reason"]
	data := netReturn["result"]
	// 当前天气信息
	//realtime := data.(map[string]interface{})["realtime"]
	if errorCode.(float64) != 0 {
		// 查询失败，根据自身业务逻辑进行调整修改
		fmt.Printf("请求失败:%v_%v", errorCode.(float64), reason)
		response.ValidateFail(c, reason.(string))
		return
	}
	response.Success(c, data)
}

// get 方式发起网络请求
func Get(apiURL string, params url.Values) (rs []byte, err error) {
	var Url *url.URL
	Url, err = url.Parse(apiURL)
	if err != nil {
		fmt.Printf("解析url错误:\r\n%v", err)
		return nil, err
	}
	//如果参数中有中文参数,这个方法会进行URLEncode
	Url.RawQuery = params.Encode()
	resp, err := http.Get(Url.String())
	if err != nil {
		fmt.Println("err:", err)
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

// ListPage godoc
// @Summary flutter管理
// @Description 查询资讯
// @Tags flutter管理
// @ID /flet/newsList
// @Accept  json
// @Produce  json
// @Security Auth
// @Param token query string true "token" （时间来不及，该swagger文档入参有问题，需要更新该swagger文档）
// @Router /flet/newsList [post]
func (f *fletController) GetNewsLists(c *gin.Context) {
	var form dto.FletNewsInput
	if err := c.ShouldBindQuery(&form); err != nil {
		response.ValidateFail(c, request.GetErrorMsg(form, err))
		return
	}
	// 接口请求参数
	params := map[string]string{
		"key":  newsKey,   // 在个人中心->我的数据,接口名称上方查看
		"type": form.Type, // 类型：top(推荐,默认)；更多看请求参数说明

	}
	// 请求头设置
	headers := map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
	}
	data, err := HttpRequest("GET", newsUrl, params, headers, 15)
	if err != nil {
		fmt.Println("请求异常：", err.Error())
		response.ValidateFail(c, request.GetErrorMsg(form, err))
		return
	}
	response.Success(c, data)
}

// http请求发送
func HttpRequest(method, rawUrl string, bodyMaps, headers map[string]string, timeout time.Duration) (result string, err error) {
	var (
		request  *http.Request
		response *http.Response
		res      []byte
	)
	if timeout <= 0 {
		timeout = 5
	}
	client := &http.Client{
		Timeout: timeout * time.Second,
	}

	// 请求的 body 内容
	data := url.Values{}
	for key, value := range bodyMaps {
		data.Set(key, value)
	}

	jsons := data.Encode()

	if request, err = http.NewRequest(method, rawUrl, strings.NewReader(jsons)); err != nil {
		return
	}

	if method == "GET" {
		request.URL.RawQuery = jsons
	}

	// 增加header头信息
	for key, val := range headers {
		request.Header.Set(key, val)
	}

	// 处理返回结果
	if response, err = client.Do(request); err != nil {
		return "", err
	}

	defer response.Body.Close()

	if res, err = io.ReadAll(response.Body); err != nil {
		return "", err
	}
	return string(res), nil
}

// 获取当前目录下的视频
func (f *fletController) GetVideo(c *gin.Context) {
	// 获取当前文件下的视频
	path := global.App.Config.Storage.Disks.LocalStorage.RootVideoDir
	var list = make([]dao.Video, 0)
	taverFile(path, &list)
	data := dto.VideoListOutput{
		List: &list,
	}
	response.Success(c, data)
	return
}

func taverFile(path string, list *[]dao.Video) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
		return
	}

	for _, file := range files {
		if file.IsDir() {
			taverFile(path+file.Name()+"/", list)
		} else {
			if mimeType(path+file.Name()) == "video/mp4" {
				videotype := strings.Split(path, "/")
				url, err := GetSnapshot(path+file.Name(), global.App.Config.Storage.Disks.LocalStorage.RootImageDir+filepath.Base(file.Name()), file.Name(), 1)
				if err != nil {
					log.Fatal("获取封面失败")
				}
				addr, _ := netAddr()
				videoUrl := "http://" + addr + "/videos/play/" + file.Name()
				item := &dao.Video{
					Name:     file.Name(),
					Path:     path + file.Name(),
					Types:    videotype[len(videotype)-2],
					ImgUrl:   url,
					VideoUrl: videoUrl,
				}
				*list = append(*list, *item)
			}
		}
	}
}

func mimeType(path string) string {
	file, err := os.Open(path)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	// Get the file content
	contentType, err := GetFileContentType(file)

	if err != nil {
		panic(err)
	}
	return contentType
}

func GetFileContentType(ouput *os.File) (string, error) {

	// to sniff the content type only the first
	// 512 bytes are used.

	buf := make([]byte, 512)

	_, err := ouput.Read(buf)

	if err != nil {
		return "", err
	}
	contentType := http.DetectContentType(buf)
	return contentType, nil
}

func GetSnapshot(videoPath, snapshotPath string, filename string, frameNum int) (string, error) {
	isExist, err := utils.PathExists(snapshotPath + filename)

	if err != nil {
		return "", err
	}
	if isExist != true {
		buf := bytes.NewBuffer(nil)
		err := ffmpeg.Input(videoPath).
			Filter("select", ffmpeg.Args{fmt.Sprintf("gte(n,%d)", frameNum)}).
			Output("pipe:", ffmpeg.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).
			WithOutput(buf, os.Stdout).
			Run()
		if err != nil {
			log.Fatal("输入缩略图失败：", err)
			return "", err
		}

		img, err := imaging.Decode(buf)
		if err != nil {
			log.Fatal("缩略图编码失败：", err)
			return "", err
		}

		err = imaging.Save(img, snapshotPath+".png")
		if err != nil {
			log.Fatal("保存缩略图失败：", err)
			return "", err
		}
	}
	// 本地网络
	addr, err := netAddr()
	if err != nil {
		return "", err
	}
	url := "http://" + addr + ":" + global.App.Config.App.Port + "/assets/" + filename + ".png"
	return url, nil
}

func netAddr() (string, error) {
	// 思路来自于Python版本的内网IP获取，其他版本不准确
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return "", errors.New("internal IP fetch failed, detail:" + err.Error())
	}
	defer conn.Close()
	// udp 面向无连接，所以这些东西只在你本地捣鼓
	res := conn.LocalAddr().String()
	res = strings.Split(res, ":")[0]
	return res, nil

}

func (f *fletController) PlayVideo(c *gin.Context) {
	//通过动态路由方式获取文件名，以实现下载不同文件的功能
	name := c.Query("name")
	types := c.Query("types")
	if name == "" {
		response.ValidateFail(c, "请输入视频名字")
		return
	}
	if types == "" {
		response.ValidateFail(c, "请输入视频类型")
		return
	}
	//拼接路径,如果没有这一步，则默认在当前路径下寻找
	addr, err := netAddr()
	if err != nil {
		response.ValidateFail(c, err.Error())
		return
	}
	filename := "http://" + addr + ":" + global.App.Config.App.StaticPort + "/" + types + "/" + name
	//响应一个文件
	var data = make(map[string]string)
	data["addr"] = filename
	response.Success(c, data)
	return
}

func (f *fletController) GetBanner(c *gin.Context) {
	lists := make([]string, 0)
	addr, err := netAddr()
	if err != nil {
		response.BusinessFail(c, err.Error())
		return
	}
	for i := 1; i < 4; i++ {
		// str := "http://" + addr + global.App.Config.Storage.Disks.LocalStorage.RootImageDir + "hot" + string(i) + ".avif"
		str := "http://" + addr + ":" + global.App.Config.App.Port + "/assets/" + "hot" + strconv.Itoa(i) + ".avif"
		lists = append(lists, str)
	}
	response.Success(c, lists)
	return
}
