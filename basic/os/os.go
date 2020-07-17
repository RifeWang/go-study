package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("系统环境变量:\n %v \n", os.Environ())
	fmt.Println(os.Executable()) // 系统可执行文件路径，go run 返回的是临时打包编译的路径

	// defer func() {
	// 	fmt.Println("defer func")
	// }()
	// os.Exit(-1)  // 程序立即退出，不会触发 defer 函数。通常状态码零代表成功，非零代表异常。

	// 设置环境变量，ExpandEnv 替换字符串中的 $var 或者 ${var} 为环境变量值。
	os.Setenv("NAME", "gopher")
	os.Setenv("BURROW", "/usr/gopher")
	fmt.Printf("获取指定键值的环境变量: %v\n", os.Getenv("BURROW"))
	fmt.Println(os.ExpandEnv("$NAME lives in ${BURROW}."))

	fmt.Println(os.Getwd()) // 当前目录对应的根目录
	fmt.Println(os.Hostname())
	fmt.Println(os.UserHomeDir())

	file, err := os.Open("/dev/null")
	fmt.Println(file, err)
}
