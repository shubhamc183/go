package main

import "fmt"

func main() {
	statePopulations := map[string]int{
		"california": 23,
		"nj":         12,
		"ny":         22,
	}
	fmt.Printf("%v, %T\n", statePopulations, statePopulations)
	statePopulations["ohio"] = 2
	fmt.Printf("%v, %T\n", statePopulations, statePopulations)
	delete(statePopulations, "nj")
	fmt.Printf("%v, %T\n", statePopulations, statePopulations)

	value, ok := statePopulations["asdasd"]
	fmt.Printf("%v, %v\n", value, ok)

	fmt.Printf("%v", len(statePopulations))

}
