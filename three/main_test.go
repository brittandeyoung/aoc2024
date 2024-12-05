package main

import "testing"

func TestPart1(t *testing.T) {
	muliplications := parseMultiplications("./inputs/sample.txt", false)

	solution := part1Solve(muliplications)
	providedSolution := 161
	if solution != providedSolution {
		t.Fatalf(`Part1: solution of sample data (%d) did not match provided sample answer (%d)`, solution, providedSolution)
	}

}

func TestPart2(t *testing.T) {
	muliplications := parseMultiplications("./inputs/sample2.txt", true)

	solution := part2Solve(muliplications)
	providedSolution := 48
	if solution != providedSolution {
		t.Fatalf(`Part2: solution of sample data (%d) did not match provided sample answer (%d)`, solution, providedSolution)
	}

}
