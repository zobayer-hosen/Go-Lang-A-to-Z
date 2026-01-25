package main

import "fmt"

var a = 1
//a :=1 you never write this way outside of the main function

func main() {
	{
		v := 1
		{
			fmt.Println(v)
		}
		fmt.Println(v)
	}
	fmt.Println(a)
	//fmt.Println(v) is going to be error cause outside of the block sport
	
}
