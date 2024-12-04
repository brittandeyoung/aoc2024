package main

import "testing"

func TestHelloName(t *testing.T) {
	solution := solve("./inputs/sample.txt")
	providedSolution := 11
	if solution != providedSolution {
		t.Fatalf(`solution of sample data (%d) did not match provided sample answer (%d)`, solution, providedSolution)
	}
}
