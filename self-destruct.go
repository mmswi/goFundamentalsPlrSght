package main

import (
	"fmt"
	"time"
)

func main() {
	for timer := 10; timer >= 0; timer -- {
		//*
		if timer == 0 {
			fmt.Println("BOOM!!!")
			break // this gets us out of the loop and blocks the code following it
		}
		//*/
		//* - comment below code to see difference
		if timer % 2 == 0{
			continue // continue stops the normal loop flow, skipping the code below and jumps back to the top of the loop
		}
		//*/
		fmt.Println(timer)
		time.Sleep(1 * time.Second)
	}
}
