package http

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

const BASE_URL = ""

type MacApi struct {
	BaseUrl       string
	Authorization string
}

func (m *MacApi) PostJson(requestPath string, data interface{}, headers map[string]string) string {
	client := &http.Client{}
	requestBytes, _ := json.Marshal(data)
	requestUrl := BASE_URL + requestPath
	if m.BaseUrl != "" {
		requestUrl = m.BaseUrl + requestPath
	}
	fmt.Println(requestUrl)
	req, err := http.NewRequest("POST", requestUrl, strings.NewReader(string(requestBytes)))
	if err != nil {
		log.Fatal("http.newRequest", err)
		return err.Error()
	}
	req.Header.Set("Content-Type", "application/json;charset=UTF-8")
	if m.Authorization != "" {
		req.Header.Set("Authorization", m.Authorization)
	}
	if headers != nil {
		for key, val := range headers {
			req.Header.Set(key, val)
		}
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("client.do", err)
		return err.Error()
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("ioutil.ReadAll", err)
		return err.Error()
	}
	respHtml := string(body)
	fmt.Println(respHtml)
	return respHtml
}

func (m *MacApi) GetJson(requestPath string, headers map[string]string) string {
	clientGet := http.Client{}
	requestUrl := BASE_URL + requestPath
	if m.BaseUrl != "" {
		requestUrl = m.BaseUrl + requestPath
	}
	req, err := http.NewRequest(http.MethodGet, requestUrl, nil)
	if err != nil {
		log.Fatal("http.NewRequest", err)
	}
	if m.Authorization != "" {
		req.Header.Set("Authorization", m.Authorization)
	}
	req.Header.Set("Authorization", m.Authorization)
	if headers != nil {
		for key, val := range headers {
			req.Header.Set(key, val)
		}
	}
	resp, err := clientGet.Do(req)
	if err != nil {
		log.Fatal("clientGet.Do", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("ioutil.ReadALL", err)
	}

	respGetHtml := string(body)

	return respGetHtml
}
