package main

import (
	"fmt"
	"strconv"
)

func main() {
	arr := [...]int{10, 20, 30, 40, 50}
	fmt.Println(arr)
	fmt.Println(len(arr))

	arr1 := [...]string{"apple", "banana", "cherry", " date", " elderberry"}
	fmt.Println(arr1)
	arr1[2] = "citrus"
	fmt.Println(arr1)
	arr1[4] = strconv.Itoa(arr[0]) + " pie"
	fmt.Println(arr1[4])
	fmt.Println(arr1)
	a, _ := strconv.Atoi(arr1[0])
	arr[1] = 25 + a
	fmt.Println(arr)

	var c [5]int //this c is defined interger types array which has 5 elements
	var b [5]int
	b = [...]int{10, 20, 30, 40, 50} // this b is defined inner the b array another 5 interger types elements which is interger types elements
	fmt.Println(c == b)

	//array pointer
	p := [...]int{1, 2, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(&p[0])
	d := &p[len(p)-1]
	fmt.Println(d)
	fmt.Println(*d)

	//MULTI DIMENSIONAL ARRAY

	arr3 := [3][4]int{{4, 5, 6, 7}, {3, 4, 6, 7}, {6, 7, 8, 9}}
	fmt.Println(arr3)

	o := make([]int, 3) //capacity  and length 3
	fmt.Println(o)

	j := make([]int, 3, 5) //length 3 and capacity 5
	j[0] = 10
	j[1] = 20
	j[2] = 30
	// j = append(j, 40, 50)
	fmt.Println(j)

	fmt.Println(len(j))
	fmt.Println(cap(j))

	

}
