package main

import "fmt"

type person struct {
	name    string
	age     int
	address string
	pincode int
	mobile  int64
}
type Car struct {
	brand, model, color string
	year, price         int
}

func main() {

	p1 := person{"zobayer", 23, "dhaka", 1234, 3456789}

	fmt.Println(p1)
	fmt.Println(p1.address)
	p1.age = 33         //updating age field
	fmt.Println(p1.age) //printing updating age field

	var p2 person = person{name: "jamal", age: 43, address: "ctg", pincode: 5678, mobile: 9876543210}
	p2.name = "speed"
	p2.age = 21
	p2.address = "usa"
	p2.pincode = 5678
	p2.mobile = 9876543210

	fmt.Println(p2)

	c := &Car{"toyota", "corolla", "red", 2020, 30000}

	fmt.Println(*c)
	fmt.Println(c.brand)//dereferencing done automatically
	

}
