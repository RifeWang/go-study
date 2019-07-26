package main

import (
	// "encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	// 构造客户端
	client := &http.Client{}

	// 构造请求数据
	body := url.Values{}
    body.Set("name", string("++=="))
	body.Set("password", string("123456"))
	fmt.Println("req body: ", body, body.Encode())

	// 构造 request 请求，注意 body 编码
	req, _ := http.NewRequest("PATCH", "http://localhost:4001/test", strings.NewReader(body.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// 客户端发起请求
	res, err := client.Do(req)
	fmt.Println("response: ", res, err)

}
