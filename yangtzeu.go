package oauth

import (
	"code.google.com/p/mahonia"
	"errors"
	"fmt"
	"net/url"
	"regexp"
)

type YangtzeuAuth struct {
	BaseOAuth
	studentNum      string
	password        string
	viewstate       string
	eventvalidation string
}

func (this *YangtzeuAuth) getParms(content string) (string, string) {
	regexp1 := `([\d\D]+id="__VIEWSTATE" value=")([\d\D][^"]+)("[\d\D]+id="__EVENTVALIDATION" value=")([\d\D][^"]+)("[\d\D]+)`
	reg := regexp.MustCompile(regexp1)
	viewstate := reg.ReplaceAllString(content, "$2")
	eventvalidation := reg.ReplaceAllString(content, "$4")
	return viewstate, eventvalidation
}

func (this *YangtzeuAuth) GetData(studentNumber string, password string) (interface{}, error) {
	reqUrl := "http://jwc.yangtzeu.edu.cn:8080/login.aspx"

	this.studentNum = studentNumber
	this.password = password

	content, err := this.get(reqUrl)
	if nil != err {
		fmt.Println("Get Request Error: ", reqUrl)
		return nil, err
	} else {
		this.viewstate, this.eventvalidation = this.getParms(content)
		if len(this.viewstate) > 0 && len(this.eventvalidation) > 0 {
			fmt.Println(this.viewstate)
			fmt.Println(this.eventvalidation)

			// post request for login
			result, err := this.post(
				reqUrl,
				url.Values{
					"__VIEWSTATE":       {this.viewstate},
					"__EVENTVALIDATION": {this.eventvalidation},
					"txtUid":            {this.studentNum},
					"txtPwd":            {this.password},
					"selKind":           {"1"},
				},
			)

			// cd, err := iconv.Open("gbk", "utf-8")
			// if err != nil {
			// 	fmt.Println("iconv.Open failed!")
			// 	return
			// }
			// defer cd.Close()
			// result := cd.ConvString(result)

			fmt.Println(result)
			return result, err
		} else {
			fmt.Println("Invalid Parms")
			return nil, errors.New("Invalid Parms")
		}
	}
}
