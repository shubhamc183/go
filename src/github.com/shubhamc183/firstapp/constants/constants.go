package main

import (
	"fmt"
)

const (
	a = iota // 0
	b = iota // 0
	c = iota // 0
)

const (
	d = iota // 0
	e        // 1
	f        // 2
)

const (
	_ = iota // 0
	g        // 1
	h        // 2
	i        // 3
)

const (
	_ = 5
	j // 6
	k // 7
	l // 8
)

const (
	_  = iota
	KB = 1 << (10 * iota)
	MB
	BG
	TB
	PB
)

func main() {

	// IF made in CAPS first letter then it would be exported
	// cont MYCONTS

	// will fail as constant are assigned at compile time
	// const myConts int = math.Sin(1)

	const myConts int = 20
	fmt.Printf("%v, %T\n", myConts, myConts)

	// Can't change its value
	// myConts = 213

	const myConts2 = 10
	var b int16 = 10
	fmt.Printf("%v, %T\n", myConts2+b, myConts2+b)
	// fmt.Printf("%v, %T\n", 10+b, 10+b)
	// implicit type conversion works with Const but not with Vars

}
