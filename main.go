package main

import (
	_ "embed"
	"fmt"
)

type Handler interface {
	Do(k, v interface{})
}

// HandlerFunc 接口型函数
type HandlerFunc func(k, v interface{})

func (f HandlerFunc) Do(k, v interface{}) {
	f(k, v)
}


func Each(m map[interface{}]interface{}, h Handler) {
	if m != nil && len(m) > 0 {
		for k, v := range m {
			h.Do(k, v)
		}
	}
}

func EachFunc(m map[interface{}]interface{}, f func(k, v interface{})) {
	Each(m, HandlerFunc(f))
}


type welcome string

// Do 如果不用接口型函数 函数名就不能表达确切的含义
func (w welcome) Do(k, v interface{}) {
	fmt.Printf("%s,我叫%s,今年%d岁\n", w,k, v)
}

func selfInfo(k, v interface{}) {
	fmt.Printf("大家好,我叫%s,今年%d岁\n", k, v)
}

func InterfaceFunc() {
	persons := make(map[interface{}]interface{})
	persons["张三"] = 20
	persons["李四"] = 23
	persons["王五"] = 26

	var w welcome = "大家好"

	Each(persons, w)
}

//go:embed main.go
var src string

func main() {
	//decimal.DivisionPrecision = 2
	//d1 := decimal.NewFromFloat(0.1).Add(decimal.NewFromFloat(0.2))
	//fmt.Println(d1.String()) // output: "0.6666666666666667"\
	//var a float64
	//var b float64
	//a = 0.1
	//b = 0.2
	//fmt.Println(a + b)
}

func testAbs()  {
	var a int8
	a = -128
	fmt.Println(absInt(a))
}

func absInt(x int8) int8 {
	return absDiffInt(x, 0)
}

func absDiffInt(x, y int8) int8 {
	if x < y {
		return y - x
	}
	return x - y
}