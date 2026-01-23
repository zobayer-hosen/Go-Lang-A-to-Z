package main

import "fmt"

const (
	firstname = "zobayer"
	lastname  = "hosen"
	age       = 23
)

func main() {
	const trueConst = true
	type myBool bool
	var defaultBool = trueConst
	var customBool myBool = trueConst

	fmt.Println(defaultBool)
	fmt.Println(customBool)

	fmt.Println(firstname, lastname, age)

}
