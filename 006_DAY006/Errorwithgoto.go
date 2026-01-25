package main

import (
	"errors"
	"fmt"
)

func process() error {
	fmt.Println("the message strating process")

	err := errors.New("an error occurred")
	if err != nil {
		goto ErrorHandler
	}

	fmt.Println("Process complete successfully")
	return nil

ErrorHandler:
	fmt.Println("Handling error and exiting")
	return err

}
func main(){
	err := process()
	if err != nil{
		fmt.Println("process failed",err)
	}
}
