package download

import (
	"os"
	"net/http"
	"github.com/gin-gonic/gin"
	"onlineDisk/model"
)

func DownloadFile(ctx *gin.Context) {
	filename := ctx.Query("file")
	_, e := os.Stat(model.SaveDir + filename)
	if e != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": e.Error(),
		})
		return
	} else {
		ctx.File(model.SaveDir + filename)
	}
}