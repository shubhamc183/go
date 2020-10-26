package main

import "fmt"

func main() {
	var userLoggedIn bool = false
	fmt.Printf("%v, %T\n", userLoggedIn, userLoggedIn)

	var i int = 10
	var j int8 = 20

	// invalid operation: i + j (mismatched types int16 and int8)
	// fmt.Printf(i + j)

	fmt.Printf("%v\n", i+int(j))

	var s string = "I am learning go lang!"
	fmt.Printf("%v, %T\n", s, s)
	// String in Go is aliases for bytes
	fmt.Printf("%v, %T\n", s[0], s[0])
	fmt.Printf("%v, %T\n", string(s[0]), s[0])

	// can't do assingment for string just like in Python
	// s[0] = "u"

	// String as slice of bytes
	s = "this is string"
	b := []byte(s)
	fmt.Printf("%v, %T\n", b, b)

	// String Represent UTF-8 character
	// Rune represent UTF-32 character

	// Single Quote for rune
	// Rune are just a type alias for int32
	var r rune = 'a'
	fmt.Printf("%v, %T\n", r, r)

}
