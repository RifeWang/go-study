package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("hello")

	// ----------------------------------
	// 并发执行，不会等待全部输出完成
	// for i := 0; i < 10; i++ {
	// 	go func(x, y int) {
	// 		z := x + y
	// 		fmt.Println(x, z)
	// 	}(i, i)
	// }
	// ----------------------------------

	// ----------------------------------
	// 通过 channel 阻塞
	// chs := make([]chan int, 10)
	// for i := 0; i < 10; i++ {
	// 	chs[i] = make(chan int)
	// 	go func(x, y int, ch chan int) {
	// 		z := x + y
	// 		fmt.Println("====", x, z)
	// 		ch <- 1
	// 	}(i, i, chs[i])
	// }
	// for _, ch := range chs {
	// 	<-ch // 只有对应 channel 写入之后才能读取，此操作是阻塞的
	// }
	// ----------------------------------

	// ----------------------------------
	// 通过 sync 处理
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(x, y int) {
			defer wg.Done()
			z := x + y
			fmt.Println("----", x, z)
		}(i, i)
	}
	wg.Wait()
	// ----------------------------------
}
