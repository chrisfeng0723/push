/**
 * @author fengxinlei
 * @date 2021/7/16 17:10
 */
package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Test(c *gin.Context){
	c.JSON(http.StatusOK,"hello,push")
	return
}
