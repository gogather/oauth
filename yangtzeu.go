package oauth

import (
	"errors"
	"fmt"
	"github.com/guotie/gogb2312"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
	"strings"
)

type YangtzeuAuth struct {
	BaseOAuth
	studentNum string
	password   string
}

// 获取预参数
func (this *YangtzeuAuth) getParms(content string) (string, string) {
	regexp1 := `([\d\D]+id="__VIEWSTATE" value=")([\d\D][^"]+)("[\d\D]+id="__EVENTVALIDATION" value=")([\d\D][^"]+)("[\d\D]+)`
	reg := regexp.MustCompile(regexp1)
	viewstate := reg.ReplaceAllString(content, "$2")
	eventvalidation := reg.ReplaceAllString(content, "$4")

	return viewstate, eventvalidation
}

// 获取数据，入口函数
// studentNumber 学号
// password 密码
func (this *YangtzeuAuth) GetData(studentNumber string, password string) (interface{}, error) {
	reqUrl := "http://jwc.yangtzeu.edu.cn:8080/login.aspx"

	this.studentNum = studentNumber
	this.password = password

	content, err := this.get(reqUrl)
	if nil != err {
		fmt.Println("Get Request Error: ", reqUrl)
		return nil, err
	} else {
		viewstate, eventvalidation := this.getParms(content)
		if len(viewstate) > 0 && len(eventvalidation) > 0 {

			resp, err := http.Post(reqUrl,
				"application/x-www-form-urlencoded",
				strings.NewReader(fmt.Sprintf(
					"__VIEWSTATE=%s&__EVENTVALIDATION=%s&txtUid=%s&txtPwd=%s&selKind=1&selKind=1&btLogin=%B5%C7%C2%BD",
					url.QueryEscape(viewstate),
					url.QueryEscape(eventvalidation),
					studentNumber,
					password,
				)),
			)
			if err != nil {
				fmt.Println(err)
			}

			defer resp.Body.Close()
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				fmt.Println(err)
			}

			result := string(body)

			resultUtf8, err, _, _ := gogb2312.ConvertGB2312String(result)
			if err != nil {
				return nil, errors.New("gb2312 convert into utf8 error.")
			}

			name, class := this.parseStuInfo(resultUtf8)

			return map[string]interface{}{
				"name":  name,
				"class": class,
			}, err

		} else {
			fmt.Println("Invalid Parms")
			return nil, errors.New("Invalid Parms")
		}
	}
}

// 提取学生信息
func (this *YangtzeuAuth) parseStuInfo(body string) (string, string) {
	regexp1 := `([\d\D]+FlowLayout">)([\d]+)( <br> )([\d\D][^<]+)(<br> )([\d\D][^<]+)(</DIV><BR>[\d\D]+)`
	reg := regexp.MustCompile(regexp1)

	name := reg.ReplaceAllString(body, "$4")
	class := reg.ReplaceAllString(body, "$6")

	return name, class
}
