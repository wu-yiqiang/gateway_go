package controllers

import (
	"encoding/json"
	"fmt"
	"gateway_go/dto"
	"gateway_go/request"
	"gateway_go/response"
	"github.com/gin-gonic/gin"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
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
