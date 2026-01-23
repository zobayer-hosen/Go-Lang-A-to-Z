package main

import "fmt"

func main() {

	//procedural way
	arr := []string{"my", "name", "is", "zobayer"}

	for i, j := range arr {
		fmt.Println(i, j)
	}

	mapping := map[string]string{
		"name":  "zobayer",
		"age":   "22",
		"class": "hey",
	}

	for i, j := range mapping {
		fmt.Println(i, j)
	}

	i := 0
	for i < 4 {

		i += 2

	}
	fmt.Println(i)
	//indfinity loops
	for {
		fmt.Println("the is my journey")
	}

}
