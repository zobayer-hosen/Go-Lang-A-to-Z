package main

import "fmt"

func add(x, y int) int {
	add := x + y

	return add
}

//call by reference
func division(a, b *int) float64 {
	div := float64(*a) / float64(*b)
	return div
}

func main() {

	a := add(10, 20)
	fmt.Println(a)

	b := 2
	fmt.Println(division(&a, &b))

}
