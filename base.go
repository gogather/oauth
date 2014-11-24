package oauth

import (
	// "fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	// "strings"
)

type BaseOAuth struct {
	AccessToken  string
	ClientId     string
	ClientSecret string
	Code         string
}

func (this *BaseOAuth) get(reqUrl string) (string, error) {
	response, err := http.Get(reqUrl)
	if nil != err {
		response.Body.Close()
		return "", err
	}
	body, err := ioutil.ReadAll(response.Body)
	if nil != err {
		response.Body.Close()
		return "", err
	}
	return string(body), nil
}

func (this *BaseOAuth) post(reqUrl string, parms url.Values) (string, error) {
	resp, err := http.PostForm(reqUrl, parms)

	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	return string(body), err
}
