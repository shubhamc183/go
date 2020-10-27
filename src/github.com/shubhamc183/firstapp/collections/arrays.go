package main

import (
	"fmt"
)

func main() {

	// array with fixed size
	var students = [3]string{"shubham", "jesse"}
	fmt.Printf("%v, %T\n", students, students)

	// array to store any number of data
	var students1 = [...]string{"shubham", "jesse"}
	fmt.Printf("%v, %T\n", students1, students1)

	// empty array
	var students2 [3]string
	fmt.Printf("%v, %T\n", students2, students2)

	a := [...]int{10, 20, 30}
	// assign by value
	b := a
	b[0] = -10
	fmt.Printf("%v, %v\n", a, b)

	// assign by refrence
	c := &a
	c[0] = 111
	fmt.Printf("%v, %v\n", a, c)
	fmt.Printf("%v, %v\n", a[1:2], c)
}
