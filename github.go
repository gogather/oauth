package oauth

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type GithubOAuth struct {
	BaseOAuth
	parms string
}

func (this *GithubOAuth) NewGithubOAuth(ClientId string, ClientSecret string, Code string) error {
	var err error
	this.ClientId = ClientId
	this.ClientSecret = ClientSecret
	this.Code = Code
	this.parms, err = this.getParms()
	return err
}

func (this *GithubOAuth) GetData() (interface{}, error) {
	url := "https://api.github.com/user?access_token=" + this.parms
	json, err := this.get(url)
	if nil != err {
		fmt.Println("Request Failed: " + url)
		return nil, err
	}
	return this.jsonDecode(json)
}

func (this *GithubOAuth) getParms() (string, error) {
	url := `https://github.com/login/oauth/access_token?client_id=` + this.ClientId + `&client_secret=` + this.ClientSecret + `&code=` + this.Code
	return this.get(url)
}

// json解码
func (this *GithubOAuth) jsonDecode(data string) (interface{}, error) {
	dataByte := []byte(data)
	var dat interface{}

	err := json.Unmarshal(dataByte, &dat)
	return dat, err
}

func (this *GithubOAuth) get(url string) (string, error) {
	response, _ := http.Get(url)
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	return string(body), err
}
