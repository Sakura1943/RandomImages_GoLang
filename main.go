package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	app "random_images/app/random_images"
	"random_images/configuration"
)

func main() {
	Config := configuration.Configuarion()
	r := gin.Default()
	r.Static("/image", Config.Base.ImagePath)
	r.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "aaa")
	})
	app.RandomImagesRouters(r)
	err := r.Run(fmt.Sprintf(":%d", Config.Base.Port))
	if err != nil {
		return
	}
}
