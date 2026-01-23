package main

import "fmt"

func multiply(a, b int) int {
	return a * b
}

//call by reference
func division(a, b *int) int {

	*a = *a / *b //modify the value at its memory address
	return *a

}
func main() {
	result := multiply(10, 5)
	fmt.Println(result)
	a, b := 20, 4
	res := division(&a, &b)
	fmt.Println(res)
}
