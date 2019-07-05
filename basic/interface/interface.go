package main

import (
	"fmt"
	"math"
	"time"
)

// I ...
type I interface {
	M()
}

// T ...
type T struct {
	S string
}

// M ...
func (t *T) M() {
	if t == nil { // 方法的 receiver 可以是 nil
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

// F ...
type F float64

// M ...
func (f F) M() {
	fmt.Println(f)
}

// Person ...
type Person struct {
	Name string
	Age  int
}

/* 最普遍的是 fmt 包的 Stringer 接口
type Stringer interface {
	String() string
}
*/
// 实现 String() 方法，控制打印行为
func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

// MyError ...
type MyError struct {
	When time.Time
	What string
}

// 自定义 Error 方法
func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s",
		e.When, e.What)
}
func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

/*
	接口是一种特殊类型，只要实现了接口的所有函数就实现了该接口
	接口具有：值、值类型、方法
*/
func main() {
	var i I
	// describe(i)  // nil 接口既不包含值也不包含具体类型
	// i.M()  // 调用 nil 接口上的方法将导致运行时错误

	var t *T
	i = t
	describe(i)
	i.M() // receiver 是 nil

	i = &T{"Hello"} // 接口类型赋值
	describe(i)
	i.M()

	// ------------------------------
	// 类型断言   value, ok := interface.(T)
	ss, ok := i.(*T)
	fmt.Println(ss, ok)

	i = F(math.Pi) // 接口类型赋值
	describe(i)
	i.M()

	// ------------------------------
	// 空接口
	var ei interface{}
	describer2(ei)

	ei = 42
	describer2(ei)

	ei = "hello"
	describer2(ei)

	// 类型断言
	s, ok := ei.(string)
	fmt.Println(s, ok)

	// type switch
	do(21)
	do("hello")
	do(true)

	// ------------------------------
	// 实现自定义 String() 方法，控制打印行为
	a := Person{"Arthur Dent", 42}
	z := Person{"Zaphod Beeblebrox", 9001}
	fmt.Println(a, z)

	// ------------------------------
	// 实现自定义 Error() 方法
	if err := run(); err != nil {
		fmt.Println(err)
	}
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func describer2(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

// ---------------------------------
// 类型判定 type switch 特定写法
func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}
