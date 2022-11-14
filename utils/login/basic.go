package login

import (
	"net/http"
	"onlineDisk/module"
	"github.com/gin-gonic/gin"
	"fmt"
)

// 待完善，用数据库存储账号密码信息
var pass = module.User{Email: "222@qq.com",Passwd: "000000"}

var LoginHtml = func(ctx *gin.Context) {
	ctx.HTML(http.StatusOK,"login.html",nil)
}

var SubmitContent = func(ctx *gin.Context) {
	var user module.User
	if err := ctx.ShouldBind(&user);err == nil {
		fmt.Printf("login info:%#v\n",user)
		if user.Email != pass.Email || user.Passwd != pass.Passwd {
			ctx.JSON(401,"password err")
		} else {
			ctx.HTML(http.StatusOK,"index.html",nil)
		}
	}
}