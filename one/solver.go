package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func solve(file string) int {
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

	sort.Ints(leftNumbers)
	sort.Ints(rightNumbers)

	totalDistance := 0

	for i := range lines {
		indexDistance := math.Abs(float64(leftNumbers[i] - rightNumbers[i]))
		totalDistance += int(indexDistance)
	}

	return totalDistance
}
