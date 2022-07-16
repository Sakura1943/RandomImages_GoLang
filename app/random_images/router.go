package app

import (
	"github.com/gin-gonic/gin"
)

func RandomImagesRouters(e *gin.Engine) {
	e.GET("/random_images", RandomImagesHandler())
}
