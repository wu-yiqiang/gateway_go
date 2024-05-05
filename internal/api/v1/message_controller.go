package v1

import (
	"net/http"

	"gateway_go/chat-room/common/request"
	"gateway_go/chat-room/common/response"
	"gateway_go/internal/service"

	"github.com/gin-gonic/gin"
)

// 获取消息列表
func GetMessage(c *gin.Context) {
	// log.Logger.Info(c.Query("uuid"))
	var messageRequest request.MessageRequest
	err := c.BindQuery(&messageRequest)
	if nil != err {
		// log.Logger.Error("bindQueryError", log.Any("bindQueryError", err))
	}
	// log.Logger.Info("messageRequest params: ", log.Any("messageRequest", messageRequest))

	messages, err := service.MessageService.GetMessages(messageRequest)
	if err != nil {
		c.JSON(http.StatusOK, response.FailMsg(err.Error()))
		return
	}

	c.JSON(http.StatusOK, response.SuccessMsg(messages))
}
