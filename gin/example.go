package gin

import "github.com/gin-gonic/gin"

func main01() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong111",
		})
	})
	r.Run() // 监听并在 0.0.0.0:8080 上启动服务
}
