package main

import "fmt"

func GFG() func(i, j string) string {
	value := func(i, j string) string {
		return i + j + "hello world"
	}

	return value
}

func add(a, b int) int {
	add := a + b
	return add
}

func main() {

	value := GFG() //ai return tai hobe akta function

	c := value("welcome", "to")
	fmt.Println(c)

	a := add(10, 20) //ami je function nia kaj korsi oita aikan thake call hobe and  add(10,20) mane hoce oi add function ar return value dewa

	fmt.Println(a)
	e := func() {
		fmt.Println("Welcome! to GeeksforGeeks")

	}
	e()

	func(ele string) {
		fmt.Println(ele)
	}("hello,world")

}
