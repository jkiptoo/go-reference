// Structs are typed collections of fields
package main

import "fmt"

//Create person struct
type person struct {
	name string
	age int
}

//Create newPerson struct that creates new person using name given in the person struct

func newPerson(name string) *person {

	
	p := person{name: name}
	p.age = 52

	//Safely return a pointer to a local variable
	return &p
}

func main() {
	
	fmt.Println(person{"Mike", 30}) //Creates new struct
	
	fmt.Println(person{name: "Alex", age: 22}) //Name field when initializing struct
	
	fmt.Println(person{name: "Rick"}) //Undeclared fields will be zero valued in case of numbers or emptied in case of strings

	fmt.Println(&person{name: "Aline", age: 35}) //Gives a pointer to the struct

	fmt.Println(newPerson("Jon"))//Surround new structs using a constructor functions

	//Use dot to access struct fields
	n := person{name: "Erick", age: 24}
	fmt.Println(n.name)

	//Use dot with struct pointers
	sp := &n
	fmt.Println(sp.age)

	//Mutability in structs
	sp.age = 29
	fmt.Println(sp.age)

	//For table driven tests, use anonymous structs where only a single value is used
	animal := struct {
		name string
		isGood bool }{
			"Simba",
			true,
		}
		fmt.Println(animal)
}