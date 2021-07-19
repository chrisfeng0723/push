/**
 * @author fengxinlei
 * @date 2021/7/16 17:33
 */
package websocket

import (
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
	"time"
)

func StartWebSocket() {
	http.HandleFunc("/push", WsPush)
	fmt.Println("WebSocket 启动程序成功", "ip", "8888")

	err := http.ListenAndServe(":"+"8888", nil)
	if err != nil {
		fmt.Println("WebSocket 启动程序失败", "ip", "8888")
	}
}

func WsPush(w http.ResponseWriter, req *http.Request) {
	conn, err := (&websocket.Upgrader{CheckOrigin: func(r *http.Request) bool {
		return true
	}}).Upgrade(w, req, nil)
	if err != nil {
		http.NotFound(w, req)
		return
	}

	currentTime := uint64(time.Now().Unix())
	client := NewClient(conn.RemoteAddr().String(),conn,currentTime)
	//初始化发送一个hello
	err = client.Socket.WriteMessage(websocket.TextMessage, []byte("hello"))
	if err != nil {
		fmt.Println(err)
		return
	}
	go client.Write()
	go client.Read()

}

