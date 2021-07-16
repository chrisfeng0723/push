/**
 * @author fengxinlei
 * @date 2021/7/16 16:56
 */
package main

import (

	"github.com/gin-gonic/gin"
	"net/http"
	"push/internal/app/routers"
)
func main(){
	router := gin.Default()
	routers.Init(router)


	http.ListenAndServe(":"+"7777", router)
}
