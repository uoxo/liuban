package main

//导入gin框架
import (
	"github.com/gin-gonic/gin"
)

//函数主入口
func main() {

	//初始化gin web框架
	r := gin.Default()

	//get接口
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "测试成功",
		})
	})

	//端口
	r.Run("0.0.0.0:80")
}
