package main

import "fmt"

func main() {
	leftNumbers, rightNumbers := parseFile("./one/inputs/main.txt")

	solutionPart1 := part1Solve(leftNumbers, rightNumbers)
	fmt.Println("Part1 solution:", solutionPart1)

	solutionPart2 := part2Solve(leftNumbers, rightNumbers)
	fmt.Println("Part2 solution:", solutionPart2)
}
