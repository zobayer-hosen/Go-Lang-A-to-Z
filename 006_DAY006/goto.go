package main

import "fmt"

func main() {

	fmt.Println("start")

	goto skip

	fmt.Println("this line will be skiped")

skip:
	fmt.Println("okey this program is end ")
	goto see //the important this where it is start and where it is stop this is way it is going
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i == 1 && j == 1 {
				goto Endloop
			}
			fmt.Printf("i:%d,j:%d\n", i, j)
		}
	}
Endloop:
	fmt.Println("Exited the loop")
see:
	fmt.Println("this is done")

}
