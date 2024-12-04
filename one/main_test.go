package main

import "testing"

func TestPart1(t *testing.T) {
	leftNumbers, rightNumbers := parseFile("./inputs/sample.txt")

	solution := part1Solve(leftNumbers, rightNumbers)
	providedSolution := 11
	if solution != providedSolution {
		t.Fatalf(`solution of sample data (%d) did not match provided sample answer (%d)`, solution, providedSolution)
	}
}

func TestPart2(t *testing.T) {
	leftNumbers, rightNumbers := parseFile("./inputs/sample.txt")

	solution := part2Solve(leftNumbers, rightNumbers)
	providedSolution := 31
	if solution != providedSolution {
		t.Fatalf(`solution of sample data (%d) did not match provided sample answer (%d)`, solution, providedSolution)
	}
}
