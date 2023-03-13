package test

import (
	"encoding/json"
	"fmt"
	"iot-platform/define"
	"iot-platform/helper"

	"testing"
)

var userServiceAddr = "http://127.0.0.1:14000"

type loginReply struct {
	Token string
}

func TestUserLogin(t *testing.T) {

	m := define.M{
		"name":     "test",
		"password": "123456",
	}
	data, _ := json.Marshal(m)
	rep, err := helper.HttpPost(userServiceAddr+"/user/login", data)
	if err != nil {
		t.Fatal(err)
	}
	lr := new(loginReply)
	err = json.Unmarshal(rep, &lr)
	fmt.Println("11111", lr.Token)
	if err != nil {
		return
	}
	fmt.Println(string(rep))
}
