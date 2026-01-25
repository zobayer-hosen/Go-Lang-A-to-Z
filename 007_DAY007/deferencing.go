package main

import "fmt"

func increment(n *int) {
	*n++
}

func main() {

	x := 5
	increment(&x)
	fmt.Println(x)

}
