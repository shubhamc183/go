package main

import (
	"fmt"
	"strconv"
)

// vars can be grouped which are closer
var (
	firstName  string = "shubham"
	secondName string = "choudhary"
	age        int    = 23
)

// global variable
// var I int = 100

func main() {

	// 3 ways to declare variables
	var i int
	i = 10
	fmt.Println(i)

	var j int = 11
	fmt.Println(j)

	k := 12
	fmt.Println(k)

	// fmt.Println(I)

	x := 99.9
	fmt.Println(x)

	// Implicit type conversion so that data loss is known to user
	var y int = int(x)
	fmt.Println(y)

	fmt.Printf("%v, %T\n", age, age)
	ageAsString := strconv.Itoa(age)
	fmt.Printf("%s, %T\n", ageAsString, ageAsString)

}
