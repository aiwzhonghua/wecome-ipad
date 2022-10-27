package router

import (
	"github.com/gin-gonic/gin"
	"wecome-ipad/controller/auto"
	"wecome-ipad/controller/check"
	"wecome-ipad/controller/contact"
	"wecome-ipad/controller/qrcode"
	"wecome-ipad/room"
)

func Router() *gin.Engine {

	r := gin.New()

	wecome := r.Group("")
	{
		//获取二维码
		wecome.GET("/qrcode", qrcode.QrcodeController)

		//检查用户是否登陆
		wecome.GET("/login/check", check.Check)

		//自动登陆
		wecome.GET("/login/auto", auto.Auto)

		//获取外部联系人
		wecome.GET("/contact/sync/customer", contact.SyncCustomer)

		//同步所有同事信息
		wecome.GET("/contact/sync/colleague", contact.Colleague)

		//更新同事备注
		wecome.PUT("/contact/sync/colleague", contact.UpdateColleague)

		//更新外部联系人公司电话
		wecome.PUT("/contact/customer", contact.UpdateColleagueCustomer)

		//获取正在群聊的群，此群可能未保存到通讯录
		wecome.GET("/room/sessions", room.Sessions)

	}

	return r
}
