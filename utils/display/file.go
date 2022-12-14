package display

import (
	"fmt"
	"io"
	"net/http"
	"onlineDisk/model"
	"os"

	"github.com/gin-gonic/gin"
)

func ShowFiles(c *gin.Context) {
	_, e := os.Stat(model.SaveDir)
	if e != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": e.Error(),
		})
		return
	} else {
		fes, _ := os.ReadDir(model.SaveDir)
		list := []string{}
		for _, fe := range fes {
			list = append(list, fe.Name())
		}
		c.JSON(http.StatusOK, gin.H{
			"message": list,
		})
	}
}

func ShowPng(c *gin.Context) {
	// 分块编码
	c.Writer.Header().Set("Transfer-encoding", "chunked")
	// 图片类型
	c.Writer.Header().Set("Content-type", "image/png")
	// 这里是模拟将分块数据写入http连接
	for i := 0; i <= model.BlockNum; i++ {
		f, err := os.Open(model.SaveDir + fmt.Sprintf("file_%d", i))
		if err != nil {
			panic(err)
		}
		defer f.Close()
		b, err := io.ReadAll(f)
		if err != nil {
			panic(err)
		}
		c.Writer.Write(b)
		c.Writer.(http.Flusher).Flush()
	}
}
