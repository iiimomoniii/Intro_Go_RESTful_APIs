package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	//buckup working dir ไปที่ runningDir
	runningDir, _ := os.Getwd()
	//
	router.MaxMultipartMemory = 8 << 20 // 8mb
	router.POST("/upload", func(c *gin.Context) {
		file, _ := c.FormFile("file")
		//save image ให้อยู่ working dir เดียวกับ index.go
		//c.SaveUploadedFile(file, fmt.Sprintf("%s/%s", runningDir, file.Filename))
		//save image ให้อยู่ sub dir  index.go
		c.SaveUploadedFile(file, fmt.Sprintf("%s/images/%s", runningDir, file.Filename))
		//reply status กลับ
		c.String(http.StatusOK, fmt.Sprintf("'%s' upload!", file.Filename))
	})
	router.Run(":85")
}
