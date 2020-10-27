package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Printf("%v, %T\n", a, a)
	fmt.Printf("%v\n", a[1:])
	fmt.Printf("%v\n", a[2:4])
	fmt.Printf("%v\n", a[:])

	b := make([]int, 3)
	fmt.Printf("%v, len=%v, capacity=%v, %T\n", b, len(b), cap(b), b)
	a = append(a, 10)
	fmt.Printf("%v\n", a)
	a = append(a, 20, 30)
	fmt.Printf("%v\n", a)
	a = append(a, b...)
	fmt.Printf("%v\n", a)

	// pop first element
	a = a[1:]

	// pop last element
	a = a[:len(a)-1]

	// This will change a!
	// danagerous
	b = append(a[:2], a[4:]...)

}
