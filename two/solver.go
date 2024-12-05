package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func parseReports(file string) [][]int {
	input, err := os.ReadFile(file)

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	lines := strings.Split(string(input), "\n")

	reports := make([][]int, len(lines))

	for i, line := range lines {
		parts := strings.Fields(line)
		levels := make([]int, len(parts))
		for j, part := range parts {
			num, _ := strconv.Atoi(part)
			levels[j] = num
		}
		reports[i] = levels
	}

	return reports
}
func isReportSafe(levels []int) bool {
	increasing := true
	decreasing := true
	consistent := false

	for i := 1; i < len(levels); i++ {
		if levels[i] > levels[i-1] {
			decreasing = false
		} else if levels[i] < levels[i-1] {
			increasing = false
		}
		diff := int(math.Abs(float64(levels[i] - levels[i-1])))

		consistent = increasing != decreasing
		isSafe := diff >= 1 && diff <= 3 && consistent

		if isSafe == false {
			return false
		}

	}

	return true
}

func isReportSafeWithDamper(levels []int) bool {
	if isReportSafe(levels) {
		return true
	} else {
		for i := range levels {
			newLevels := make([]int, len(levels)-1)
			copy(newLevels[:i], levels[:i])

			copy(newLevels[i:], levels[i+1:])

			// newLevels = append(levels[:i], levels[i+1:]...)
			fmt.Println("index: ", i, "newlevels: ", newLevels, "levels: ", levels)
			if isReportSafe(newLevels) {
				return true
			}
		}
	}
	return false
}

func part1Solve(reports [][]int) int {
	safeReportCount := 0
	for _, levels := range reports {
		isSafe := isReportSafe(levels)

		if isSafe {
			safeReportCount += 1
		}
	}
	return safeReportCount
}

func part2Solve(reports [][]int) int {
	safeReportCount := 0
	for _, levels := range reports {
		isSafe := isReportSafeWithDamper(levels)
		if isSafe {
			safeReportCount += 1
		}
	}
	return safeReportCount
}
