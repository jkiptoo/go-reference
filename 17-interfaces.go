package main

import (
	"fmt"
	"math"
)

//Interfaces groups and names related sets of methods

//Basic interface for geometric shapes
type geometry interface {
	area() float64
	perimeter() float64
}

//Implement interface on rectangle and circle types
type rectangle struct {
	width, height float64
	}

type circle struct {
	radius float64
}

//Implement all methods in the interface
func (r rectangle) area() float64 {
	return r.width * r.height
}
func (r rectangle) perimeter() float64 {
	return 2*r.width + 2*r.height 
}

//Implement interfaces in circles
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

//Call variables in the interface if variable has interface type
func measurement (g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perimeter())
}

//Use type assertion to know the runtime type of an interface value. Type switch/type casting is another alternative.
func detectShape(g geometry) {
	if c, okay := g.(circle); okay {
		fmt.Println("Circle with radius is ", c.radius)
	}
}

func main() {
	r := rectangle{width:3, height: 4}
	c := circle{radius: 5}
	
	//Both circle and rectangle implement the geometry interface
	measurement(r)
	measurement(c)

	detectShape(r)
	detectShape(c)
	
}