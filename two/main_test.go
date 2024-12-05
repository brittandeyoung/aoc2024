package main

import "testing"

func TestPart1(t *testing.T) {
	reports := parseReports("./inputs/sample.txt")

	solution := part1Solve(reports)
	providedSolution := 2
	if solution != providedSolution {
		t.Fatalf(`Part1: solution of sample data (%d) did not match provided sample answer (%d)`, solution, providedSolution)
	}
}

func TestPart2(t *testing.T) {
	reports := parseReports("./inputs/sample.txt")

	solution := part2Solve(reports)
	providedSolution := 4
	if solution != providedSolution {
		t.Fatalf(`Part2: solution of sample data (%d) did not match provided sample answer (%d)`, solution, providedSolution)
	}
}
