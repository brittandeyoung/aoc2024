package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func parseMultiplications(file string, enableFunctions bool) [][]int {
	input, err := os.ReadFile(file)

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	inputString := string(input)
	if enableFunctions {
		stripRe := regexp.MustCompile(`(?s)don't\(\)(.*?)do\((.*?)\)`)
		inputString = stripRe.ReplaceAllString(inputString, "")
		stripReFinal := regexp.MustCompile(`(?s)don't\(\)(.*)`)
		inputString = stripReFinal.ReplaceAllString(inputString, "")
	}
	re := regexp.MustCompile(`mul\(([\d.]+),([\d.]+)\)`)
	matches := re.FindAllStringSubmatch(inputString, -1)
	multiplications := [][]int{}

	for _, match := range matches {
		left, _ := strconv.Atoi(match[1])
		right, _ := strconv.Atoi(match[2])
		multiplication := []int{left, right}
		multiplications = append(multiplications, multiplication)
	}
	return multiplications
}

func multiplyNumbers(numbers []int) int {
	product := 1
	for _, num := range numbers {
		product *= num
	}
	return product
}

func part1Solve(multiplications [][]int) int {
	stock := 0
	for _, v := range multiplications {
		product := multiplyNumbers(v)
		stock += product
	}

	return stock
}

func part2Solve(multiplications [][]int) int {
	stock := 0
	for _, v := range multiplications {
		product := multiplyNumbers(v)
		stock += product
	}

	return stock
}
