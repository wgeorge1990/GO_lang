package main

import (
	"fmt"
)

func main() {

	//looping over slice returning index and value using range
	x := []int{1, 2, 3, 4, 5, 6}

	for i, v := range x {
		fmt.Println(i, v)
	}

	
}