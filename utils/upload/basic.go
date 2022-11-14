package upload

import (
	"fmt"
	"log"
	"net/http"
	"onlineDisk/module"

	"github.com/gin-gonic/gin"
)

var UploadFile = func(c *gin.Context) {
	// 单个文件
	file, err := c.FormFile("f1")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	log.Println(file.Filename)
	dst := fmt.Sprintf(module.SaveDir+"%s", file.Filename)
	// 上传文件到指定的目录
	c.SaveUploadedFile(file, dst)
	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("'%s' uploaded!", file.Filename),
	})
}