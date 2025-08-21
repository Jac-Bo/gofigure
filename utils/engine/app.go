package engine

import "github.com/gin-gonic/gin"

func StartGin() {
	gin := gin.Default()
	gin.GET("/figures")
	gin.POST("/figures")
	gin.Run("localhost")
}
