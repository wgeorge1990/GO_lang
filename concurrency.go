package main

import (
	"fmt"
	"time"
)

func main() {
	// go count("sheep")
	// go count("fish")
	// time.Sleep(time.Microsecond * 00)
	// fmt.Scanln()

	// var wg sync.WaitGroup
	// wg.Add(1)
	// go func() {
	// 	count("sheep")
	// 	wg.Done()
	// }()

	// wg.Wait()
	
	//Channel reciever
	c := make(chan string)
	go count("sheep", c)

	// for {
	// 	msg, open  := <- c
	// 	//if no longer open then break out of for loop and finish running function.
	// 	if !open {
	// 		break
	// 	}

	for msg := range c {
		fmt.Println(msg)
	}

	var message string = "If you are seeing this message then go routing was completed and the channel was successfully closed and function finished"

	fmt.Println(message)
}

func count(thing string, c chan string) {
	for i := 1; i  <= 5; i++ {
		// fmt.Println(i, thing)
		c <- thing
		time.Sleep(time.Millisecond * 500)
	}

	//sender is always responsible for closing the channel when it is done in order to prevent go routine deadlock.
	close(c)
}
