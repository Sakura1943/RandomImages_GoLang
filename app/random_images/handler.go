package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"os"
	"random_images/configuration"
)

func RandomImagesHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		Config := configuration.Configuarion()
		_, err := os.Stat(Config.Base.ImagePath)
		if err != nil {
			err := os.Mkdir(Config.Base.ImagePath, os.ModePerm)
			if err != nil {
				panic(fmt.Sprintf("Folder %s create failed!", Config.Base.ImagePath))
			}
			fmt.Printf("Folder %s created!", Config.Base.ImagePath)
		}
		readAll, err := os.ReadDir(Config.Base.ImagePath)
		if err != nil {
			panic(fmt.Sprintf("read folder failed: %s", Config.Base.ImagePath))
		}
		var readList []string
		for _, file := range readAll {
			if !file.IsDir() {
				fullName := Config.Base.ImagePath + "/" + file.Name()
				readList = append(readList, fullName)
			}
		}
		if len(readList) == 0 {
			panic(fmt.Sprintf("Folder %s have not image.", Config.Base.ImagePath))
		}
		randomNumber := rand.Intn(len(readList))
		fmt.Println(fmt.Sprintf("/%s", readList[randomNumber]))
		ctx.Redirect(http.StatusFound, fmt.Sprintf("/%s", readList[randomNumber]))
	}
}
