// Using math/rand/v2 package to generate random numbers
package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	//Use rand.IntN to return random integers.
	//Here, we limit to 100
	//Format as n, 0 <= n < 100

	printLine := fmt.Println
	print := fmt.Print
	printLine(rand.IntN(100), ",")
	printLine(rand.IntN(100))
	printLine()

	//Use rand.Float64 to return random floats
	printLine(rand.Float64())

	//Generate random floats in other ranges
	print((rand.Float64()*5)+5, ",")
	print((rand.Float64() * 5) + 5)
	printLine()

	//Create a new "rand.Source" and pass to a new constructor  
	newSeed := rand.NewPCG(45, 1024)
	newRand := rand.New(newSeed)
	print(newRand.IntN(100), ",")
	print(newRand.IntN(100))
	printLine()

	//Use NewPCG creates a new PCG that utilises two uint64 numbers
	newerSeed := rand.NewPCG(45, 1024)
	newerRand := rand.New(newerSeed)
	print(newerRand.IntN(100), ",")
	print(newerRand.IntN(100))
	printLine()

}