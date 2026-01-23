package main

import "fmt"

func main() {
	for i := 0; i < 5; fmt.Println("after") {

		fmt.Println("before", i)
		i++
		continue
	}
}
