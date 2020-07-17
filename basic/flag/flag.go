package main

import (
	"flag"
	"fmt"
)

func main() {
	// 命令行参数 -f
	f := flag.String("f", "default value", "usage commend")

	// 命令行参数 -port
	var port string
	flag.StringVar(&port, "port", "8080", "http listen port")

	flag.Parse()

	fmt.Println(*f, port)

	// 执行 go run flag.go -f myflag -port 1234
	// 使用未定义的命令行 flag 会导致 panic
}
