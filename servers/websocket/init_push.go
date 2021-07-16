/**
 * @author fengxinlei
 * @date 2021/7/16 19:52
 */
package websocket

import (
	"fmt"
	"github.com/gorilla/websocket"
	"time"
)

type Client struct {
	Socket *websocket.Conn
}

func NewClient(socket *websocket.Conn) *Client {
	return &Client{
		Socket: socket,
	}
}
func (c *Client) Read() {
	for {
		_, data, err := c.Socket.ReadMessage()
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(data)
	}

}
func (c *Client) Write(message string) {
	if message == "" {
		message = "hello" + time.Now().Format("2006-01-02 15:04:05")
	}
	for {
		data := []byte(message)
		err := c.Socket.WriteMessage(websocket.TextMessage, data)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}
