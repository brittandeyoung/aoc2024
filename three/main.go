package main

import "fmt"

func main() {
	muliplications := parseMultiplications("./three/inputs/main.txt", false)

	solutionPart1 := part1Solve(muliplications)
	fmt.Println("Part1 solution:", solutionPart1)

	muliplicationsPart2 := parseMultiplications("./three/inputs/main.txt", true)

	solutionPart2 := part1Solve(muliplicationsPart2)
	fmt.Println("Part2 solution:", solutionPart2)
}
