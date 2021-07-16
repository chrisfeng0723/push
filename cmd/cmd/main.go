/**
 * @author fengxinlei
 * @date 2021/7/16 16:56
 */
package main

import (
	"github.com/gin-gonic/gin"
	"net/http"

	"push/internal/app/routers"
	"push/servers/websocket"
)

func main() {
	router := gin.Default()
	routers.Init(router)

	go websocket.StartWebSocket()
	http.ListenAndServe(":"+"7777", router)
	//router.Run(":"+"7777")
}
