package test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestUserGet(t *testing.T) {

	client := &http.Client{}

	url := "http://172.16.15.98:8001/lowcode/bug/instState/1070674822611329024"
	//提交请求
	reqest, err := http.NewRequest("GET", url, nil)

	if err != nil {
		fmt.Println(err)
	}
	//增加header选项
	reqest.Header.Add("Authorization", "Bearer "+"eyJhbGciOiJIUzUxMiJ9.eyJ1c2VyX2lkIjoxLCJ1c2VyX2tleSI6IjVjMGNkNzY4LTI1NTgtNDZmZS04YWY5LTk2OGQ3MWNmZDI0ZCIsInVzZXJuYW1lIjoiYWRtaW4ifQ.bgeYetz3x9P6d2kcPLgXbHIsOJfNFq4ihWNj5hsFx68jaf1ihzSoPbdPamlNbGRU6VhADMNDRZZSkSxOdHamww")
	//

	//处理返回结果
	response, err := client.Do(reqest)
	if err != nil {
		panic(err)
	}
	if response.Body != nil {
		defer response.Body.Close()
	}
	getRes, gerErr := ioutil.ReadAll(response.Body)
	if gerErr != nil {
		fmt.Println(gerErr)
	} else {

		fmt.Println(string(getRes))
	}

}
