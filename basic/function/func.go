package main

import (
	"fmt"
	"strings"
)

// ----------------------------------------
// ----------------------------------------
// 返回值命名、可变参数
func add(list ...int) (c int) {
	for _, v := range list {
		c += v // 直接赋值，不得重新声明
	}
	return // 直接 return
}

// ----------------------------------------

// ----------------------------------------
// ----------------------------------------
// 函数是一等公民，可以作为参数传递。
func stringProccess(list []string, chain []func(string) string) {
	for index, str := range list { // 遍历每一个字符串
		result := str                // 第一个需要处理的字符串
		for _, proc := range chain { // 遍历每一个处理链
			result = proc(result) // 依次调用处理函数。
		}
		list[index] = result // 将结果放回切片
	}
}

// 自定义的移除前缀的处理函数
func removePrefix(str string) string {
	return strings.TrimPrefix(str, "go")
}

// ----------------------------------------

// ----------------------------------------
// ----------------------------------------
// 闭包
func acc(value int) func() int {
	return func() int { // 返回一个闭包
		value++
		return value
	}
}

/*
	函数的参数都是值拷贝
	指针、切片和 map 等引用型对象指向的内容在参数传递中不会发生复制，而是将指针进行复制，类似于创建一次引用。
*/
func main() {
	// --------------------------------
	// --------------------------------
	fmt.Println(add(1, 2, 3))
	fmt.Println(add(1, 2, 3, 4, 5))

	// --------------------------------
	// --------------------------------
	// 匿名函数
	visit := func(list []int, f func(int)) {
		for _, v := range list {
			f(v)
		}
	}
	visit([]int{4, 5, 6}, func(v int) {
		fmt.Println(v)
	})
	// --------------------------------

	// --------------------------------
	// --------------------------------
	// 函数作为参数传递
	list := []string{
		"go scanner",
		"go parser",
		"go compiler",
		"go printer",
		"go formater",
	}
	chain := []func(string) string{ // 处理函数链
		removePrefix,
		strings.TrimSpace,
		strings.ToUpper,
	}
	stringProccess(list, chain) // 处理字符串
	for _, str := range list {  // 输出处理好的字符串
		fmt.Println(str)
	}
	// --------------------------------

	// --------------------------------
	// --------------------------------
	// 闭包
	acc1 := acc(1) // 初始化，记录了状态
	fmt.Println(acc1(), &acc1)
	fmt.Println(acc1(), &acc1)

	acc2 := acc(10)
	fmt.Println(acc2(), &acc2)
	fmt.Println(acc2(), &acc2)
	// --------------------------------

	// --------------------------------
	// --------------------------------
	// defer panic recover
	func() {
		defer func() {
			if p := recover(); p != nil {
				fmt.Println(p)
			}
		}()
		panic("throw error")
		fmt.Println("do not continue") // 同级不再继续执行而是直接返回
	}()
	fmt.Println("continue running") // 正常执行
	// --------------------------------

	// --------------------------------
	// --------------------------------
	// 方法
	var b myint = 5
	fmt.Println(b.addBy(50))
	m := b.addBy
	fmt.Println(m(550))

	cc := ms{
		x: 1,
		y: 2,
	}
	cc.change()
	fmt.Println(cc)
}

// --------------------------------
// 方法：可以为任意类型添加方法
type myint int

func (v myint) addBy(a int) int {
	return int(v) + a
}

type ms struct {
	x int
	y int
}

func (m *ms) change() { // 指针类型
	m.x, m.y = m.y, m.x
}
