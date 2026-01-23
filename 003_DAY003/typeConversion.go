package main

import (
	"fmt"
	"strconv"
)

func main() {
	a := 10

	c := float64(a)
	fmt.Println(c)
	fmt.Printf("the value type is %T\n", c)
	d, _ := fmt.Println(strconv.Itoa(a))
	fmt.Println(d)
	fmt.Printf("the a value type is %T\n", d)

	myint := 42
	myfloat := float64(myint)
	//println//just use for quick print and debug
	fmt.Println(myfloat)
	fmt.Printf("the myfloat value type is %T\n", myfloat)
	f := 42.99
	e := int(f)
	fmt.Println(e)

	mystring := "123"
	intvalue, err := strconv.Atoi(mystring)
	if err != nil {
		fmt.Println("Error converting string to int ", err)
	} else {
		fmt.Println("string to int ", intvalue)
	}

	number := 456
	g := strconv.Itoa(number)
	fmt.Printf("the type of this value is %T\n", g)

}
