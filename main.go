package main
import (
    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()
    r.GET("/ping", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "ycl24",
        })
    })
    r.Run(":3000") // 在 0.0.0.0:8080 上监听并服务
}
