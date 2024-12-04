package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func countOccurrences(arr []int, num int) int {
	count := 0
	for _, val := range arr {
		if val == num {
			count++
		}
	}
	return count
}

func parseFile(file string) ([]int, []int) {
	input, err := os.ReadFile(file)

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	lines := strings.Split(string(input), "\n")

	leftNumbers := make([]int, len(lines))
	rightNumbers := make([]int, len(lines))

	for i, line := range lines {
		parts := strings.Fields(line)
		left, _ := strconv.Atoi(parts[0])
		right, _ := strconv.Atoi(parts[1])
		leftNumbers[i] = left
		rightNumbers[i] = right
	}

	return leftNumbers, rightNumbers
}

func part1Solve(leftNumbers, rightNumbers []int) int {
	sort.Ints(leftNumbers)
	sort.Ints(rightNumbers)

	totalDistance := 0

	for i := range leftNumbers {
		indexDistance := math.Abs(float64(leftNumbers[i] - rightNumbers[i]))
		totalDistance += int(indexDistance)
	}

	return totalDistance
}

func part2Solve(leftNumbers, rightNumbers []int) int {
	similarityScore := 0

	for i := range leftNumbers {
		occurrences := countOccurrences(rightNumbers, leftNumbers[i])
		similarityScore += int(leftNumbers[i] * occurrences)
	}

	return similarityScore
}
