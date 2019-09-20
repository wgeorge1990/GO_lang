package main

import (
	"fmt"
	"time"
)

func main() {
	// c := make(chan string, 3)
	// c <- "hello"
	// c <- "world"
	// c <- "goodbye"

	// msg := <- c
	// fmt.Println(msg)

	// msg = <- c
	// fmt.Println(msg)

	// msg = <- c
	// fmt.Println(msg)

	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		for {
			c1 <- "every 500ms"
			time.Sleep(time.Millisecond * 500)
		}
	}()

	go func() {
		for {
			c2 <- "every two seconds"
			time.Sleep(time.Second * 2)
		}
	}()

		for {
			select {
			case msg1 := <- c1:
				fmt.Println(msg1)
			case msg2 := <- c2:
				fmt.Println(msg2)
			}
		}
	}