package main

import (
	"fmt"
)

/*
type CallBack interface {
	CallMe(int, int)
}*/

type Handler func(int, int) int

func (f Handler) Call(a int, b int) int {
	return f(a, b)
}

func Add(a int, b int) int {
	return a + b
}

func main() {
	result := Handler(Add).Call(1, 2)
	fmt.Println(result)
}
