package main

//Embedding of structs and interfaces helps in having a smooth composition of types

import "fmt"



type base struct {
	number int
}

func (b base) describe() string {
	return fmt.Sprintf("Base with num=%v", b.number)
}

//Use container embeds a base. Embedded structs take the form of a field without a name.
type container struct {
	base
	str string
}

func main() {
//Initialise embedding with the embedded type serving as a field name
co := container {
	base: base {
		number: 1,
		},
		str: "Some name",
}

//Access fields of the base
fmt.Printf("co={number: %v, str: %v}\n", co.number, co.str)

//Output full path using embedded type name
fmt.Println("Alternative number is:", co.base.number)

//When container embeds base, methods of the base apply to the container. Method from base is directly embedded on co
fmt.Println("Describe:", co.describe())

type describer interface {
	describe() string
}

//Container implementing describer interface. This is made possible because embedding structs using methods transfers interface structure into other structs
var d describer = co
fmt.Println("Describer is:", d.describe())
}