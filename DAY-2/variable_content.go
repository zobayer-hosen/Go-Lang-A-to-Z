package main

import (
	"fmt"
)

func main() {
	a := 10
	fmt.Println(a)
	// Reassigning variable 'a' to a new value
	// a:= 20 // This will cause a compile-time error
	a = 50
	fmt.Println(a)
	name, address, city := "zobayer", "mirpur", "dhaka"
	fmt.Printf("the name is %s\n", name)
	fmt.Printf("the type of %T\n", name)
	_ = address
	_ = city //blank identifier to avoid unused variable error
	age, _, _ := 25, "address", "city"
	fmt.Printf("the person age is %d\n", age)

	// Using short variable declaration again in the same scope will cause an error
	// name := "newname" // This will cause a compile-time error
	b := 30
	age = 40
	_ = age //blank identifier to avoid unused variable error
	fmt.Println(b)

	//build-in vs package variable
	c := 100
	println(c) // using build-in print function without fmt package
	d := 200
	fmt.Printf("the value is %d\n", d)

}
