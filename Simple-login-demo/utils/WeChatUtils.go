package wechat

import (
	"Simple-login-demo/setting"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

var (
	GET_ACCESS_TOKEN = "https://qyapi.weixin.qq.com/cgi-bin/gettoken?corpid=ID&corpsecret=SECRET"
	GET_USER_ID      = "https://qyapi.weixin.qq.com/cgi-bin/user/getuserinfo?access_token=ACCESS_TOKEN&code=CODE"
	GET_USER         = "https://qyapi.weixin.qq.com/cgi-bin/user/get?access_token=ACCESS_TOKEN&userid=USERID"
)

type WeChatResponse struct {
	Errcode     int    `json:"errcode"`
	Errmsg      string `json:"errmsg"`
	AccessToken string `json:"access_token"`
	ExpiresIn   string `json:"expires_in"`
	UserId      string `json:"UserId"`
}

// 获取access_token
func GetAccessToken() WeChatResponse {
	url := strings.Replace(GET_ACCESS_TOKEN, "ID", setting.Conf.WeChatConfig.Corpid, -1)
	url = strings.Replace(url, "SECRET", setting.Conf.WeChatConfig.Corpsecret, -1)
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("获取AccessToken失败ERRO：", err)
	}
	defer response.Body.Close() // 这步是必要的，防止以后的内存泄漏，切记

	// 读取响应 body, 返回为 []byte
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("转化为结构体：", err)
	}
	//转变为结构体对象
	var weChatResponse WeChatResponse
	json.Unmarshal(body, &weChatResponse)

	return weChatResponse
}

// 使用code获取用户ID
func GetUserIdByCode(code string) (WeChatResponse, error) {
	weChatResponse := GetAccessToken()
	if 0 == weChatResponse.Errcode {
		// 获取access_token成功
		url := strings.Replace(GET_USER_ID, "ACCESS_TOKEN", weChatResponse.AccessToken, -1)
		url = strings.Replace(url, "CODE", code, -1)
		response, err := http.Get(url)
		if err != nil {
			fmt.Println("获取UserId失败ERRO：", err)
			return weChatResponse, errors.New("获取UserId失败")
		}
		body, _ := ioutil.ReadAll(response.Body)

		json.Unmarshal(body, &weChatResponse)
		defer response.Body.Close() // 这步是必要的，防止以后的内存泄漏，切记
		if 0 == weChatResponse.Errcode {
			// 获取userId成功
			return weChatResponse, nil
		} else {
			return weChatResponse, errors.New("获取userId失败")
		}
		return weChatResponse, nil
	} else {
		// TODO 获取accessToken失败处理
		return weChatResponse, errors.New("获取accessToken失败")
	}

}

// 使用code获取非企业微信用户ID
