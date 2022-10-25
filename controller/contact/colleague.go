package contact

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	pkg_api "wecome-ipad/util/http"
)

type paramsColleagueSeq struct {
	Seq string `json:"seq"`
}

type ColleagueGenerated struct {
	Code int `json:"code"`
	Data struct {
		Version    string `json:"version"`
		Colleagues []struct {
			ID                string        `json:"id"`
			Weixin            string        `json:"weixin"`
			NickName          string        `json:"nick_name"`
			Sex               int           `json:"sex"`
			HeadSmallImageURL string        `json:"head_small_image_url"`
			IsAdmin           bool          `json:"is_admin"`
			Friend            bool          `json:"friend"`
			Flag              int           `json:"flag"`
			CorpID            string        `json:"corp_id"`
			Corporation       string        `json:"corporation"`
			Title             string        `json:"title"`
			Phones            []interface{} `json:"phones"`
		} `json:"colleagues"`
	} `json:"data"`
}

func Colleague(c *gin.Context) {

	var paramsColleagueSeq paramsColleagueSeq

	if err := c.BindJSON(&paramsColleagueSeq); err != nil {
		c.JSON(http.StatusOK, "参数不正确")
		return
	}

	apiHandler := pkg_api.MacApi{Authorization: c.GetHeader("Authorization")}
	ret := apiHandler.GetJson("/contact/sync/colleague?version="+paramsColleagueSeq.Seq, map[string]string{})

	respColleague := &ColleagueGenerated{}
	_ = json.Unmarshal([]byte(ret), respColleague)
	c.JSON(http.StatusOK, respColleague)
}
