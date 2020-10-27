package main

import "fmt"

// Student struct is exported i.e visible to other packages
type Student struct {
	// but name, age, marks is private as they don't start with Capital letter
	name  string
	age   int
	marks []int
}

type Animal struct {
	Name   string
	Origin string
}

// Bird is not Animal but has Animal
type Bird struct {
	// Composition to achieve OOPs
	// Embedded
	Animal
	CanFly bool
	KMPH   float64
}

func main() {

	student := Student{
		name:  "Shubham",
		age:   24,
		marks: []int{90, 80, 30},
	}
	fmt.Printf("%v, %T\n", student, student)
	fmt.Printf("%v, %T\n", student.name, student.name)

	// Anonymous Struct just like Anonymous function/class in java
	aDoctor := struct{ name string }{name: "jessePinkman"}
	fmt.Printf("%v, %T\n", aDoctor.name, aDoctor.name)

	// pass by value
	// a copy is made, so more memory is used
	bDoctor := aDoctor
	bDoctor.name = "hesinberg"
	fmt.Printf("%v, %T\n", aDoctor.name, aDoctor.name)
	fmt.Printf("%v, %T\n", bDoctor.name, bDoctor.name)

	// pass by refremce
	cDoctor := &aDoctor
	cDoctor.name = "Gustavo"
	fmt.Printf("%v, %T\n", aDoctor.name, aDoctor.name)
	fmt.Printf("%v, %T\n", cDoctor.name, cDoctor.name)

	b := Bird{}
	b.Name = "Emu"
	b.Origin = "US"
	b.CanFly = true
	fmt.Printf("%v, %T\n", b, b)
	fmt.Printf("%v, %T\n", b.Animal, b.Animal)
	fmt.Printf("%v, %T\n", b.Animal.Name, b.Animal.Name)

	bird2 := Bird{
		Animal: Animal{
			Name:   "Duck",
			Origin: "IN",
		},
		CanFly: false,
		KMPH:   3,
	}
	fmt.Printf("%v, %T\n", bird2, bird2)

}
