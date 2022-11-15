package main

import (
	"onlineDisk/utils/display"
	"onlineDisk/utils/download"
	"onlineDisk/utils/login"
	"onlineDisk/utils/upload"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	// 处理multipart forms提交文件时默认的内存限制是32 MiB
	// 可以通过下面的方式修改
	// router.MaxMultipartMemory = 8 << 20  // 8 MiB
	router.LoadHTMLFiles("./source/index.html", "./source/login.html")
	// 登录界面
	router.GET("/login", login.LoginHtml)
	// 表单上传账号密码信息
	router.POST("/login/content", login.SubmitContent)
	// 表单操作上传文件
	router.POST("/upload/content", upload.UploadFile)
	// 展示用户的文件
	router.GET("/showfiles", display.ShowFiles)
	// 下载指定名字的文件
	router.GET("/download", download.DownloadFile)
	// 上传文件分块存储
	router.POST("/upload/block", upload.UploadBlock)
	// 查看图片
	router.GET("/showpng/block", display.ShowPng)
	router.Run(":8080")
}
