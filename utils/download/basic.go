package download

import (
	"os"
	"net/http"
	"github.com/gin-gonic/gin"
	"onlineDisk/module"
)

var DownloadFile = func(ctx *gin.Context) {
	filename := ctx.Query("file")
	_, e := os.Stat(module.SaveDir + filename)
	if e != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": e.Error(),
		})
		return
	} else {
		ctx.File(module.SaveDir + filename)
	}
}