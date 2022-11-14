package display

import (
	"os"
	"onlineDisk/module"
	"net/http"
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