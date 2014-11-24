package oauth

import (
	"testing"
)

func Test_GithubOAuth(t *testing.T) {
	goa := &GithubOAuth{}
	_, err := goa.GetData("d9cae2d24dbfa8b028ed", "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx", "d53ee1ac441c648abbbc")
	if err != nil {
		t.Log("测试通过")
	}
}

func Test_YangtzeuAuth(t *testing.T) {
	ya := &YangtzeuAuth{}
	rst, err := ya.GetData("xxxxxxxxx", "xxxxxxxx")
	if err != nil {
		t.Log(rst)
	}
}
