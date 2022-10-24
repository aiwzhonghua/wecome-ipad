package router

import (
	"github.com/gin-gonic/gin"
	"wecome-ipad/controller/auto"
	"wecome-ipad/controller/check"
	"wecome-ipad/controller/contact"
	"wecome-ipad/controller/qrcode"
)

func Router() *gin.Engine {

	r := gin.New()

	wecome := r.Group("/v1")
	{
		//获取二维码
		wecome.GET("/qrcode", qrcode.QrcodeController)

		//检查用户是否登陆
		wecome.GET("/login/check", check.Check)

		//自动登陆
		wecome.GET("/login/auto", auto.Auto)

		//获取外部联系人
		wecome.GET("/contact/sync/customer", contact.SyncCustomer)
	}

	return r
}
