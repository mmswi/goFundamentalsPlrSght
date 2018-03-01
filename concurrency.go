package main

import (
	"time"
	"fmt"
	"sync"
)

func main() {

	var waitGrp sync.WaitGroup
	waitGrp.Add(2) //telling to wait for 2 goroutines

	go func() {
		//adding defer so the main doesn't close before executing this function
		defer waitGrp.Done()
		// if you don't have the "go" keyword in front of the func
		// it will put the entire program on sleep for 2 seconds
		// but now it's a goroutine
		time.Sleep(2 * time.Second)
		fmt.Println("Hello")
	}()

	go func() {
		defer waitGrp.Done()
		fmt.Println("Pluralsight")
	}()

	waitGrp.Wait()

	// goroutines can communicate vie Channels
	// channels are buffeder or unbuffered
	// myChannel := make(chan int) // a channel that can't hold any data
	// a goroutine sends data and waits for another goroutine to come and pick the data
	// forcing a synchronicity

	// myBuffChannel : make(chan int, 5) - the number indicates how many data items the channel can hold
	// the goroutine doesn't wait after sending the data on this channel

	//Getting data off the channel, blocks the goroutine, if the channel doesn't have any data
}
