package upload

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"onlineDisk/model"
	"os"

	"github.com/gin-gonic/gin"
)

func UploadFile(c *gin.Context) {
	// 单个文件
	file, err := c.FormFile("f1")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	log.Println(file.Filename)
	dst := fmt.Sprintf(model.SaveDir+"%s", file.Filename)
	// 上传文件到指定的目录
	c.SaveUploadedFile(file, dst)
	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("'%s' uploaded!", file.Filename),
	})
}

// 分成块存储
func UploadBlock(c *gin.Context) {
	file, head, err := c.Request.FormFile("f1")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	// 每个块的大小
	block := head.Size / model.BlockNum
	buf := make([]byte, block)
	i := 0
	for {
		n, err := file.Read(buf)
		if n == 0 {
			break
		}
		if err != nil && err != io.EOF {
			panic(err)
		}
		dst := model.SaveDir + fmt.Sprintf("file_%d", i)
		if err := saveToDst(dst, buf); err != nil {
			panic(err)
		}
		i++
	}
	c.JSON(200, "ok")
}

func saveToDst(name string, buf []byte) (err error) {
	f, err := os.OpenFile(name, os.O_CREATE|os.O_RDWR, 0666)
	_, err = f.Write(buf)
	defer f.Close()
	return
}
