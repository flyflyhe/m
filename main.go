package main

import "fmt"

func main() {
	s := []int{1,2,3}
	fmt.Printf("%v\n", s)
	ChangeSlice(s)
	fmt.Printf("%v\n", s)
}

func ChangeSlice(s []int) {
	s = append(s, 4)
	s[0] = 2
}