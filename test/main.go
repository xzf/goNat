package main

import "fmt"

func main() {
	fmt.Println("aaa")
}

type A struct {
	F1 string
	F2 int
}

type B struct {
	A
	F3 float64
	F4 bool
}
