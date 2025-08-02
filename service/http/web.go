package http

import (
	"github.com/gin-gonic/gin"
)

/*
*
路由组、路由
*/
func web() {

	act := WEB.Group("/test")
	act.GET("/first", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"李屁芮": "确实大臭屁!",
		})
	})
}
