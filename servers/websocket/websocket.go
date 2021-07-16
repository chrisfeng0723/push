/**
 * @author fengxinlei
 * @date 2021/7/16 17:33
 */
package websocket

import (
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
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
	//初始化发送一个hello
	err = conn.WriteMessage(websocket.TextMessage, []byte("hello"))
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		_, data, err := conn.ReadMessage()
		if err != nil {
			fmt.Println(err)
			return
		}
		err = conn.WriteMessage(websocket.TextMessage, data)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}

/**
// 升级协议
	conn, err := (&websocket.Upgrader{CheckOrigin: func(r *http.Request) bool {
		fmt.Println("升级协议", "ua:", r.Header["User-Agent"], "referer:", r.Header["Referer"])

		return true
	}}).Upgrade(w, req, nil)
	if err != nil {
		http.NotFound(w, req)

		return
	}

	fmt.Println("webSocket 建立连接:", conn.RemoteAddr().String())

	currentTime := uint64(time.Now().Unix())
	client := NewClient(conn.RemoteAddr().String(), conn, currentTime)

	go client.read()
	go client.write()
*/
