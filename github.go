package oauth

import (
	"encoding/json"
	"errors"
	"fmt"
)

type GithubOAuth struct {
	BaseOAuth
	parms string
}

func (this *GithubOAuth) newGithubOAuth(ClientId string, ClientSecret string, Code string) error {
	var err error
	this.ClientId = ClientId
	this.ClientSecret = ClientSecret
	this.Code = Code
	this.parms, err = this.getParms()

	return err
}

func (this *GithubOAuth) GetData(ClientId string, ClientSecret string, Code string) (interface{}, error) {
	err := this.newGithubOAuth(ClientId, ClientSecret, Code)
	if err != nil {
		fmt.Println("[oauth error]: ", err)
		return nil, err
	}

	url := "https://api.github.com/user?" + this.parms

	jsonStr, err := this.get(url)
	if nil != err {
		fmt.Println("Request Failed: " + url)
		return nil, err
	}

	json, err := this.jsonDecode(jsonStr)
	if check, ok := json.(map[string]interface{}); ok {
		if nil != check["message"] {
			fmt.Println("[msg] ", check["message"])
			return nil, errors.New(check["message"].(string))
		}
	}
	return json, err
}

// 获取token等参数
func (this *GithubOAuth) getParms() (string, error) {
	url := `https://github.com/login/oauth/access_token?client_id=` + this.ClientId + `&client_secret=` + this.ClientSecret + `&code=` + this.Code

	return this.get(url)
}

// json解码
func (this *GithubOAuth) jsonDecode(data string) (interface{}, error) {
	dataByte := []byte(data)
	var dat interface{}

	err := json.Unmarshal(dataByte, &dat)
	if nil != err {
		return nil, err
	}
	return dat, nil
}
