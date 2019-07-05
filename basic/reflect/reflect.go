package main

import (
	"fmt"
	"reflect"
)

/*
	Type : reflect.Type
	Value : reflect.Value
*/
func main() {
	// Reflection goes from interface value to reflection object.
	var x float64
	x = 3.4
	fmt.Println("type:", reflect.TypeOf(x))
	fmt.Println("value:", reflect.ValueOf(x).String()) // String 方法不会深挖到具体的值

	/*
		Value 的方法:
			func (v Value) Kind() Kind : 底层类型
			func (v Value) Type() Type : 值的类型，用户自定义类型
			func (v Value) Float() float64 : 底层的值
	*/
	v := reflect.ValueOf(x)
	fmt.Println("type: ", v.Type())
	fmt.Println("kind is fload64: ", v.Kind() == reflect.Float64)
	fmt.Println("value: ", v.Float())

	type MyInt int
	var a MyInt = 7
	av := reflect.ValueOf(a)
	fmt.Println("type: ", av.Type())
	fmt.Println("kind is fload64: ", av.Kind() == reflect.Int)
	fmt.Println("value: ", av.Int())

	// ------------------------------------------
	// Reflection goes from reflection object to interface value.
	/*
		func (v Value) Interface() (i interface{})
		var i interface{} = (v's underlying value)
	*/
	y := v.Interface().(float64) // y will have type float64.
	fmt.Println(v.Interface(), y)

	// ------------------------------------------
	// To modify a reflection object, the value must be settable.
	/*
		func (v Value) CanSet() bool : 是否可设置
	*/
	fmt.Println("settability of v:", v.CanSet()) // false

	p := reflect.ValueOf(&x) // Note: take the address of x.
	fmt.Println("type of p:", p.Type())
	fmt.Println("settability of p:", p.CanSet()) // false

	/*
		func (v Value) Elem() Value : returns the value that the interface v contains or that the pointer v points to.
	*/
	vv := reflect.ValueOf(&x).Elem()
	fmt.Println("settability of vv:", vv.CanSet())
	vv.SetFloat(7.1) // 改变源值
	fmt.Println(vv.Interface())
	fmt.Println(x)

	// ------------------------------------------
	/*
		func (v Value) NumField() int : struct v 的 fields 数量
		func (v Value) Field(i int) Value : struct v 的 field( Value 类型 )
	*/
	type T struct {
		A int
		B string
	}
	t := T{23, "skidoo"}
	s := reflect.ValueOf(&t).Elem() // 指向源值
	typeOfT := s.Type()             // 源值的类型
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		fmt.Printf("%d: %s %s = %v\n", i, typeOfT.Field(i).Name, f.Type(), f.Interface())
	}
	s.Field(0).SetInt(77)
	s.Field(1).SetString("Sunset Strip")
	fmt.Println("t is now", t)
}
