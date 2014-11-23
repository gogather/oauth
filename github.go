package oauth

import (
	"fmt"
)

type GithubOAuth struct {
	BaseOAuth
}

func (this *GithubOAuth) NewGithubOAuth(accessToken string) {
	this.AccessToken = accessToken
}

func (this *GithubOAuth) GetData() (interface{}, error) {
	url := "https://api.github.com/user?access_token=" + this.AccessToken
	json, err := this.get(url)
	if nil != err {
		fmt.Println("Request Failed: " + url)
		return nil, err
	}
	return this.jsonDecode(json)
}

func (this *GithubOAuth) get(url string) (string, error) {
	response, _ := http.Get(url)
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	return string(body), err
}

// json解码
func (this *GithubOAuth) jsonDecode(data string) (interface{}, error) {
	dataByte := []byte(data)
	var dat interface{}

	err := json.Unmarshal(dataByte, &dat)
	return dat, err
}
