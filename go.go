/**
常见问题：
- err is shadowed during return
*/
package main

import (
	"fmt"
	"regexp"
)

/*
func main() {
	count(10)
}

func count(num int) (err error) { // 函数中定义了一个 err变量
	if num <= 0 {
		// if分支中又重新定义了一个err变量
		err := errors.New("num数不合法")
		log.Fatalln(err)
		// 直接返回，并没有返回参数，在子作用域中，并不能直接返回上一层定义作用域中返回变量
		// 也就是说命名参数返回，只能返回同一级别作用域下参数，不是同一级别作用域参数需要指明返回参数值
		return
		// 正确返回
		//return err
	}
	// 同一级别命名参数作用域，可以直接返回
	return
}
*/

/*
	[]byte 转换为 io.Reader
*/
// Convert byte slice to io.Reader
// r := bytes.NewReader(byteData)

/*
	中文的字符长度 rune
*/
// v := "Hello, 世界"
// fmt.Println(len([]rune(v)))

/*
	bytes string 相互转换：
		b := []byte("ABC€")
		s := string([]byte{65, 66, 67, 226, 130, 172})
*/

// --------------------------
/*
	初始化顺序：
		常量 -> 变量 -> init() -> main()
*/
/*
package main

import (
   "fmt"
)

var T int64 = a()

func init() {
   fmt.Println("init in main.go ")
}

func a() int64 {
   fmt.Println("calling a()")
   return 2
}

func main() {
   fmt.Println("calling main")
}

输出：calling a() -> init in main.go -> calling main
*/

/*
	测试：避免磁盘 io ，构造假的 io.Reader
*/

func main() {
	url := "http://cdn-log-user.obs.myhwclouds.com:443/5MinLog%2F20200509%2F03%2Fsns-video-up.xhscdn.com%2Fsns-video-up.xhscdn.com_202005090335.gz?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=J3NKIBVMJZVUIHREOSIX%2F20200512%2Fcn-north-1%2Fs3%2Faws4_request&X-Amz-Date=20200512T091510Z&X-Amz-Expires=86400&X-Amz-SignedHeaders=host&response-content-disposition=attachment%3Bfilename%3D%22sns-video-up.xhscdn.com_202005090335.gz%22&X-Amz-Signature=b5443b158f0601f93c1ed18860a9ef068acf230239d0e2203c5485507dde1c5c"
	fmt.Println(url)

	urlParse := regexp.MustCompile(`(:80/)|(:443/)`)
	url = urlParse.ReplaceAllString(url, "/")
	fmt.Println(url)

	// client := &http.Client{}
	// req, err := http.NewRequest("GET", url, nil)
	// req.Host = "cdn-log-user.obs.myhwclouds.com"
	// resp, err := client.Do(req)

	// resp, err := http.Get("http://cdn-log-user.obs.myhwclouds.com/5MinLog%2F20200509%2F03%2Fsns-video-up.xhscdn.com%2Fsns-video-up.xhscdn.com_202005090335.gz?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=J3NKIBVMJZVUIHREOSIX%2F20200512%2Fcn-north-1%2Fs3%2Faws4_request&X-Amz-Date=20200512T091510Z&X-Amz-Expires=86400&X-Amz-SignedHeaders=host&response-content-disposition=attachment%3Bfilename%3D%22sns-video-up.xhscdn.com_202005090335.gz%22&X-Amz-Signature=b5443b158f0601f93c1ed18860a9ef068acf230239d0e2203c5485507dde1c5c")
	// if err != nil {
	// 	log.Printf("[ERROR UPLOAD download] %s", err.Error())
	// 	return
	// }
	// defer resp.Body.Close()
	// fmt.Println(resp.StatusCode)

}
