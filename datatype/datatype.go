package main

import (
	"encoding/json"
	"fmt"
)

/*
	数据类型：
		基础类型：数字、字符串、布尔型
		复合类型：数组、结构体
		引用类型：指针、slice、map、函数、channel
		接口类型
*/
func main() {
	// -----------------------------------------
	// -----------------------------------------
	// 指针
	x := 1
	p := &x         // &x 表达式（取 x 变量的内存地址）将产生一个指向该变量的指针，指针对应的数据类型是 *int
	fmt.Println(p)  // 内存地址
	fmt.Println(*p) // *p 表达式对应 p 指针指向的变量的值
	fmt.Println(*&x)
	*p = 2 // *p 可以出现在赋值语句的左边
	fmt.Println(x)
	// -----------------------------------------

	// -----------------------------------------
	// 数组长度不可变，初始化可以通过 {索引: 值} 的方式
	r := [...]int{99: -1}
	fmt.Println(r)

	a := [2]int{1, 2}
	b := [...]int{1, 2}
	c := [2]int{1, 3}
	fmt.Println(a == b, a == c, b == c) // 数组相等：数组里的所有元素相等（元素类型可以互相比较的前提下）
	// -----------------------------------------

	// -----------------------------------------
	// -----------------------------------------
	// slice 由三部分构成：指针、长度、容量。 指针指向第一个元素对应的底层数组元素的地址。
	// 函数的形参都是值拷贝，但是复制一个 slice 只是对底层数组创建了一个新的 slice 别名
	// slice 值包含指向第一个 slice 元素的指针，因此可以在函数内部修改底层数组的元素
	// slice 无法比较
	// slice 是对底层数组元素间接访问
	reverse := func(s []int) {
		for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
			s[i], s[j] = s[j], s[i]
		}
	}
	arr := [...]int{0, 1, 2, 3, 4, 5}
	reverse(arr[:])
	fmt.Println(arr)

	// ---
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	addr1 := &numbers[0]
	fmt.Println("numbers: ", numbers, " addr: ", addr1)

	num2 := numbers[:2]
	fmt.Println("num2: ", num2, "   numbers: ", numbers)

	num2 = append(num2, 1) // 由于引用了同一个底层数组，此处的 append 导致 numbers 的值改变
	fmt.Println("num2: ", num2, " numbers: ", numbers)

	addr2 := &num2[0]
	fmt.Println("numbers: ", numbers, " addr: ", addr2)
	// -----------------------------------------

	// -----------------------------------------
	// -----------------------------------------
	// map[K]V: key/value 无序集合，key 必须是支持 == 比较的数据类型
	// 若 key 没有则将返回零值，无法对 map 中的元素取址
	ages := map[string]int{
		"alice":   31,
		"charlie": 34,
	}
	for name, age := range ages {
		fmt.Printf("%s\t%d\n", name, age)
	}
	if age, ok := ages["bob"]; !ok { // 检测 key 是否存在，区分零值
		fmt.Println("key not exist", age)
	}
	mps := map[string]map[string]int{ // map 嵌套
		"a": {
			"b": 123,
		},
	}
	fmt.Println(mps, mps["a"], mps["a"]["b"], mps["c"], mps["c"]["d"])
	// -----------------------------------------

	// -----------------------------------------
	// -----------------------------------------
	// struct
	type Point struct {
		X, Y int
	}
	type Circle struct { // 匿名嵌套
		Point
		Radius int
	}
	type Wheel struct { // 匿名嵌套
		Circle
		Spokes int
	}
	var w Wheel
	w.X = 8      // 匿名嵌套
	w.Y = 8      // 匿名嵌套
	w.Radius = 5 // 匿名嵌套
	w.Spokes = 20

	ww := Wheel{
		Circle: Circle{
			Point:  Point{X: 8, Y: 8},
			Radius: 5,
		},
		Spokes: 20,
	}
	ww2 := Wheel{
		Circle{
			Point{X: 8, Y: 8},
			5,
		},
		20,
	}
	fmt.Println(ww, ww2)
	// -----------------------------------------

	// -----------------------------------------
	// -----------------------------------------
	// JSON
	type Movie struct {
		Title  string
		Year   int  `json:"released"`        // Tag 中的 json 控制 encoding 编解码行为，第一部分指定 json 对象的名字
		Color  bool `json:"color,omitempty"` // omitempty 表示为空或零值时不生成 json 对象
		Actors []string
	}

	var movies = []Movie{
		{
			Title:  "Casablanca",
			Year:   1942,
			Color:  false,
			Actors: []string{"Humphrey Bogary", "Ingrid Bergman"},
		},
		{
			Title:  "Cool Hand Luke",
			Year:   1967,
			Color:  true,
			Actors: []string{"Paul Newman"},
		},
	}
	// marshal 将结构体slice转为 json
	// marshalIndent 多出来的两个参数分别表示每行输出前缀和每一个层级的缩进
	// data, err := json.Marshal(movies)
	data, err := json.MarshalIndent(movies, "", "    ")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%s\n", data)

	var titles []struct{ Title string }
	if err := json.Unmarshal(data, &titles); err != nil { // unmarshal 将 json 数据解码
		fmt.Println(err)
	}
	fmt.Println(titles)
}
