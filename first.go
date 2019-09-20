package main

import "fmt"

type person struct {
	name string
	age int
}

func main() {
	p := person{name: "william", age: 29}
	fmt.Println(p)
	var x int
	y := "you are amazing"
	//go  can infer  type
	fmt.Println(x)
	fmt.Println(y)

	
	a := []int{1,4,4,5,2}
	a = append(a, 14)
	mymap := make(map[string]int)
	mymap ["william"] = 29
	fmt.Println(mymap["william"])

	for index, value := range a {
		fmt.Println("index", index, "value", value)
	}
}	