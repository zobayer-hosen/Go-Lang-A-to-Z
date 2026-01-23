package main

import "fmt"

//multiple return
func number(a, b int) (int, int) {
	add := a + b
	sub := a - b

	return add, sub
}

//nacked return
func fun(a, b int) (add int, div int) {
	add = a + b
	div = a / b
	return
}

func main() {

	c, d := number(20, 10)
	fmt.Println(c, d)

	b, d := fun(20, 2)
	fmt.Println(b, d)

}
