package route

import (
	"fmt"
	"gateway_go/internal/server"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	_ "go.uber.org/zap"
	_ "log"
	"net/http"
)

var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func RunSocekt(c *gin.Context) {
	user := c.Query("user")
	if user == "" {
		return
	}
	ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}

	client := &server.Client{
		Name: user,
		Conn: ws,
		Send: make(chan []byte),
	}
	fmt.Println("client", client)
	server.MyServer.Register <- client
	go client.Read()
	go client.Write()
}
