package main

import "fmt"

func main() {
	reports := parseReports("./two/inputs/main.txt")

	solutionPart1 := part1Solve(reports)
	fmt.Println("Part1 solution:", solutionPart1)

	solutionPart2 := part2Solve(reports)
	fmt.Println("Part2 solution:", solutionPart2)
}
