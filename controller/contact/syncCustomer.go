package contact

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	pkg_api "wecome-ipad/util/http"
)

type paramsSyncSeq struct {
	Seq string `json:"seq"`
}

type SyncCustomerGenerated struct {
	Code int `json:"code"`
	Data struct {
		Seq   string `json:"seq"`
		Items []struct {
			ID                string        `json:"id"`
			Weixin            string        `json:"weixin"`
			NickName          string        `json:"nick_name"`
			Sex               int           `json:"sex"`
			HeadSmallImageURL string        `json:"head_small_image_url"`
			IsAdmin           bool          `json:"is_admin"`
			Friend            bool          `json:"friend"`
			Flag              int           `json:"flag"`
			CorpID            string        `json:"corp_id"`
			Description       string        `json:"description,omitempty"`
			Phones            []string      `json:"phones"`
			TagIds            []interface{} `json:"tag_ids"`
			AliasName         string        `json:"alias_name,omitempty"`
			Corporation       string        `json:"corporation,omitempty"`
		} `json:"items"`
	} `json:"data"`
}

func SyncCustomer(c *gin.Context) {

	var paramsSyncSeq paramsSyncSeq

	if err := c.BindJSON(&paramsSyncSeq); err != nil {
		c.JSON(http.StatusOK, "参数不正确")
		return
	}

	apiHandler := pkg_api.MacApi{Authorization: c.GetHeader("Authorization")}
	ret := apiHandler.GetJson("/contact/sync/customer?seq="+paramsSyncSeq.Seq, map[string]string{})

	respSyncCustomer := &SyncCustomerGenerated{}
	_ = json.Unmarshal([]byte(ret), respSyncCustomer)
	c.JSON(http.StatusOK, respSyncCustomer)

}
