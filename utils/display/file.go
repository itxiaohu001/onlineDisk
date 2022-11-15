package display

import (
	"fmt"
	"io"
	"net/http"
	"onlineDisk/module"
	"os"

	"github.com/gin-gonic/gin"
)

var ShowFiles = func(c *gin.Context) {
	_, e := os.Stat(module.SaveDir)
	if e != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": e.Error(),
		})
		return
	} else {
		fes, _ := os.ReadDir(module.SaveDir)
		list := []string{}
		for _, fe := range fes {
			list = append(list, fe.Name())
		}
		c.JSON(http.StatusOK, gin.H{
			"message": list,
		})
	}
}

var ShowPng = func(c *gin.Context) {
	c.Writer.Header().Set("Transfer-encoding", "chunked")
	c.Writer.Header().Set("Content-type", "image/png")
	for i := 0; i <= module.BlockNum; i++ {
		f, err := os.Open(module.SaveDir + fmt.Sprintf("file_%d", i))
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
