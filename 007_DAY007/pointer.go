package main

import "fmt"

func main() {
	var1 := 10
	prt := &var1

	fmt.Println("var=", var1)
	fmt.Println("address of var =", &var1)
	fmt.Println(prt)
	fmt.Println(*prt) //dereferencing

	*prt = 20
	fmt.Println("var after change =", var1)
}
