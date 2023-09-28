package utils

import (
	"errors"
	"fmt"
	"github.com/gorilla/websocket"
	"sync"
)

type Connection struct {
	wsConn    *websocket.Conn
	inChan    chan []byte
	outChan   chan []byte
	closeChan chan []byte
	mutex     sync.Mutex
	isClosed  bool
}

func InitConection(wsConn *websocket.Conn) (conn *Connection, err error) {
	conn = &Connection{
		wsConn:    wsConn,
		inChan:    make(chan []byte, 1000),
		outChan:   make(chan []byte, 1000),
		closeChan: make(chan []byte, 1), //用来判断连接是否被关闭，防止阻塞
	}

	// 启用读协程
	go conn.readLoop()
	// 启用写协程
	go conn.writeLoop()
	return
}

func (conn *Connection) ReadMessage() (data []byte, err error) {
	select {
	case data = <-conn.inChan:
	case <-conn.closeChan: //防止阻塞
		err = errors.New("connect is close!")
	}
	return
}

func (conn *Connection) WriteMessage(data []byte) (err error) {
	select {
	case conn.outChan <- data:
	case <-conn.closeChan: // 防止阻塞
		err = errors.New("connect is close!")
	}
	return
}

func (conn *Connection) Close() {
	conn.wsConn.Close() //线程安全的
	conn.mutex.Lock()
	if !conn.isClosed {
		//一个chan只能关闭一次，保证此代码只执行一次
		close(conn.closeChan)
		conn.isClosed = true
	}
	conn.mutex.Unlock()
}

func (conn *Connection) readLoop() {
	var (
		data []byte
		err  error
	)
	for {
		if _, data, err = conn.wsConn.ReadMessage(); err != nil {
			goto ERR
		}
		fmt.Println("data", data)
		select {
		case conn.inChan <- data: //连接关闭，inChan一直未消费，会产生阻塞
		case <-conn.closeChan: //closeChan被关闭，跳出阻塞
			goto ERR
		}
	}

ERR:
	conn.Close()
}

func (conn *Connection) writeLoop() {
	var (
		data []byte
		err  error
	)
	for {
		fmt.Println("write")
		data = <-conn.outChan
		select {
		case data = <-conn.outChan: //连接关闭，没有发送的消息，会产生阻塞
		case <-conn.closeChan: //closeChan被关闭，跳出阻塞
			goto ERR
		}
		if err = conn.wsConn.WriteMessage(websocket.TextMessage, data); err != nil {
			goto ERR
		}
	}

ERR:
	conn.Close()
}
