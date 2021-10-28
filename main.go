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

	s := []int{1,2,3, 4,5}
	fmt.Println(s)
	s1 := make([]int, len(s))
	copy(s1, s)
	s1 = append(s1[0:2], s1[3:]...)

	fmt.Println(s1)
	fmt.Println(s)
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