/**
 * @author fengxinlei
 * @date 2021/7/16 17:06
 */
package routers

import (
	"github.com/gin-gonic/gin"
	"push/internal/app/controller"
)

func Init(router *gin.Engine){

	testRouter := router.Group("/test")
	{
		testRouter.GET("/hello",controller.Test)
	}
}
