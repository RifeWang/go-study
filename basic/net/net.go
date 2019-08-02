package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	// 构造客户端
	client := &http.Client{}

	// 构造请求数据，针对 application/x-www-form-urlencoded
	body := url.Values{}
	body.Set("name", string("++=="))
	body.Set("password", string("123456"))
	fmt.Println("req body: ", body, body.Encode())

	// 构造 request 请求，注意 body 编码
	req, _ := http.NewRequest("PATCH", "http://localhost:4001/test", strings.NewReader(body.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// 客户端发起请求
	res, err := client.Do(req)
	if err != nil {
		return
	}
	defer res.Body.Close() // 必须关闭，避免内存泄漏
	fmt.Println("response: ", res, err)

	// --------------------------------------
	// --------------------------------------
	// 针对 application/json
	reqBody := new(bytes.Buffer)
	json.NewEncoder(reqBody).Encode(struct {
		Name     string `json:"name"`
		Password string `json:"password"`
	}{"++==", "123456"})

	// 构造 request 请求，注意 body 编码
	req2, _ := http.NewRequest("PATCH", "http://localhost:4001/test", reqBody)
	req.Header.Add("Content-Type", "application/json")

	// 客户端发起请求
	res2, err := client.Do(req2)
	if err != nil {
		return
	}
	defer res2.Body.Close() // 必须关闭，避免内存泄漏
	fmt.Println("response2: ", res2, err)

	// --------------------------------------
	// --------------------------------------
	// response body 是 io.ReadCloser 接口类型，包含基本的 Read 和 Close 方法
	robots, err := ioutil.ReadAll(res2.Body) // []byte, error
	if err != nil {
		fmt.Println("read response body err: ", err)
		return
	}

	result := Result{}
	err = json.Unmarshal(robots, &result) // json 解码
	if err != nil {
		fmt.Println("json unmarshal response body err: ", err)
		return
	}
}

// Result ... response body struct
type Result struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Body    struct {
		Data []string `json:"data"`
	}
}
