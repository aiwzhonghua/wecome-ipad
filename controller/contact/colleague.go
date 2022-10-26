package contact

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/url"
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

type paramsUpdateColleague struct {
	Id    string `json:"id"`
	Alias string `json:"alias"`
}

type UpdateColleagueGenerated struct {
	Code int `json:"code"`
}

//更新同事备注
func UpdateColleague(c *gin.Context) {

	var paramsUpdateColleague paramsUpdateColleague

	if err := c.BindJSON(&paramsUpdateColleague); err != nil {
		c.JSON(http.StatusOK, " 参数不正确")
		return
	}
	alias := url.QueryEscape(paramsUpdateColleague.Alias)
	apiHandler := pkg_api.MacApi{Authorization: c.GetHeader("Authorization")}
	ret := apiHandler.PutJson("/contact/colleague/"+paramsUpdateColleague.Id+"?alias="+alias, map[string]string{})

	respUpdateColleague := &UpdateColleagueGenerated{}
	_ = json.Unmarshal([]byte(ret), respUpdateColleague)
	c.JSON(http.StatusOK, respUpdateColleague)

}

type paramsColleagueCustomer struct {
	Id      string `json:"id" binding:"required"`
	Alias   string `json:"alias"`
	phone   string `json:"phone"`
	company string `json:"company"`
}

//更新外部联系电话，名称
func UpdateColleagueCustomer(c *gin.Context) {

	var paramsColleagueCustomer paramsColleagueCustomer

	if err := c.BindJSON(&paramsColleagueCustomer); err != nil {
		c.JSON(http.StatusOK, "参数不正确")
		return
	}

	alias := url.QueryEscape(paramsColleagueCustomer.Alias)
	company := url.QueryEscape(paramsColleagueCustomer.company)
	apiHandler := pkg_api.MacApi{Authorization: c.GetHeader("Authorization")}
	ret := apiHandler.PutJson("/contact/customer/"+paramsColleagueCustomer.Id+"?alias="+alias+"&company="+company+"&phone="+paramsColleagueCustomer.phone, map[string]string{})

	respUpdateColleague := &UpdateColleagueGenerated{}
	_ = json.Unmarshal([]byte(ret), respUpdateColleague)
	c.JSON(http.StatusOK, respUpdateColleague)
}
