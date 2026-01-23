package main

import "fmt"

const pi = 3.14

func main() {
	const name = "zoayer" //untyped constant
	fmt.Println("hello", name)

	fmt.Println("Happy", pi, "Day")

	const Correct = true
	fmt.Println("GO rules?", Correct)

	var B = "Hosen"
	var fullname = name + " " + B
	fullname += "!"
	fmt.Println(fullname)

	fmt.Println(pi > 3)

}
