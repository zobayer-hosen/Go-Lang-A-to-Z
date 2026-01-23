package main

import "fmt"

var name = "zobayer"

//lastname := "hosen"//give error because of using := outside function

func main() {

	lastname := "hosen"
	fmt.Println(name)
	fmt.Println(lastname)

	a, b := 10, 20
	//a,b := 30,40 //give error because of re-declaration of a and b

	println(a)
	println(b)
	c, b := 20, 30 //the left side variable are newly declared
	println(c)
	println(b)

	_, c = 50, 60 //ignore the first value using blank identifier
	_ = c
	println(c) //give the value of 60 .when we ignore a value using blank identifier the value is still assigned to the variable

}
