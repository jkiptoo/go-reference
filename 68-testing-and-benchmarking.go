// Writing unit tests in Go
// Place testing code always in the same package as the code it is testing
package main

import (
	"fmt"
	"testing"
)

//Test implementation of an integer minimum
//If code to be tested is in a file named customers.go; the name of the test file will be customers-test.go
func IntegerMinimum(a, b int) int {
	if a < b {
		return a
	}
	return b
}

//Begin a test function starting with test name
func TestIntegerMinimum(t *testing.T) {
	answer := IntegerMinimum(3, -3)
	if answer != -3 {
		
		//Use t.Error to report test errors while executing the test
		//t.Fatal stops the test once test failures are detected
		t.Errorf("Integer Minimum(3, -3) = %d; Want -3", answer)
	}
}
//Use table-driven test style to list test inputs and expected outputs in a table then pass a loop through them.
func TestIntegerMinimumTable(t *testing.T) {
	var tests = []struct {
		a, b int
		want int
	}{
		{0, 1, 0},
		{1, 0, 0},
		{2, 0, 0},
		{3, -3, -3},
		{0, -2, -2},
		{-2, 0, -2},
		{0, -1, -1},
		{-1, 0, -1},
	}
	//Run a for loop through the table to test the logic
	for _, runtests := range tests {
		testname := fmt.Sprintf("%d, %d", runtests.a, runtests.b)

		//Use t.Run to run subtests for each table entry
		t.Run(testname, func(t *testing.T) {
			answer := IntegerMinimum(runtests.a, runtests.b)
			if answer != runtests.want {
				t.Errorf("Got %d, want %d", answer, runtests.want)
			} 
		})
	}
}

//Use Benchmark to run tests and estimate the runtime of a single iteration. Code required for the code to run but not needed to be tested should go before the loop.
func BenchmarkIntegerMinimum(benchmark *testing.B) {
	for benchmark.Loop() {
		
		//Benchmark runner executes loop many times to determine reasonable estimate of a single iteration
		IntegerMinimum(1, 3)
	}
} 

//To run all tests current project in verbose mode, use 'go test -v'

//To run all benchmarks in the current project, use 'bench' flag to filter benchmark function names with a regular expression. Use 'go test -bench=.'
