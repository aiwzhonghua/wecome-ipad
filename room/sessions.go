package room

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	pkg_api "wecome-ipad/util/http"
)

type SessionsGenerated struct {
	Code int      `json:"code"`
	Data []string `json:"data"`
}

type SessionsMessageGenerated struct {
	Code int `json:"code"`
	Data []struct {
		ID        string   `json:"id"`
		WxID      string   `json:"wx_id"`
		NickName  string   `json:"nick_name"`
		Owner     string   `json:"owner"`
		MemberNum int      `json:"member_num"`
		Status    int      `json:"status"`
		Admins    []string `json:"admins"`
		Members   []string `json:"members"`
		Dismissed bool     `json:"dismissed"`
		Flag      string   `json:"flag"`
	} `json:"data"`
}

func Sessions(c *gin.Context) {

	apiHandler := pkg_api.MacApi{Authorization: c.GetHeader("Authorization")}
	ret := apiHandler.GetJson("/room/sessions", map[string]string{})

	//群聊ID
	respSessions := &SessionsGenerated{}
	_ = json.Unmarshal([]byte(ret), respSessions)

	sessionsMessageGenerated := &SessionsMessageGenerated{}
	//获取群信息
	if respSessions.Code == 0 {
		for _, respSessionsIds := range respSessions.Data {
			retSessionsMessage := apiHandler.GetJson("/room/"+respSessionsIds, map[string]string{})
			_ = json.Unmarshal([]byte(retSessionsMessage), sessionsMessageGenerated)
			c.JSON(http.StatusOK, sessionsMessageGenerated)
		}

	} else {
		c.JSON(http.StatusOK, gin.H{
			"data": "暂无数据",
		})
	}

}
