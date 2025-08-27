package main

import "fmt"

//Methods are defined in struct types
type square struct {
	width, height int
}

//Create area method that has a receiver type named *square
func (s *square) area() int {
	return s.width * s.height
}

//Define methods for pointer or value receiver types
func (s square) perimeter() int {
	return 2*s.width + 2*s.height
}

func main() {
s := square{width: 50, height: 50}

//Calling the two methods defined in the struct
fmt.Println("Area is equal to:", s.area())
fmt.Println("Perimeter is equal to:", s.perimeter())

//Using pointer to avoid copying method calls
squarepointer := &s
fmt.Println("Area is equal to:", squarepointer.area())
fmt.Println("Perimeter is equal to:", squarepointer.perimeter())
}