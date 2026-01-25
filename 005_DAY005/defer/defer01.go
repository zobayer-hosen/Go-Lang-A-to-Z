package main

import "fmt"

func mul(a, b int) int {
	mul := a * b
	fmt.Println(mul)
	return 0
}
func show() {
	fmt.Println("I show speed")
}

func main() {

	mul(10, 10)

	defer mul(10, 20)

	show()

}
