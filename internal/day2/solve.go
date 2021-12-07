package day2

import (
	"strconv"
	"strings"
)

// Solve ...
func Solve(input []string) (result int) {
	depth := 0
	position := 0
	for _, in := range input {
		parts := strings.Split(in, " ")
		value, _ := strconv.Atoi(parts[1])
		switch parts[0] {
		case "down":
			depth += value
		case "forward":
			position += value
		case "up":
			depth -= value
		}
	}
	return depth * position
}

// SolveTwo ..
func SolveTwo(input []string) (result int) {
	aim := 0
	depth := 0
	position := 0
	for _, in := range input {
		parts := strings.Split(in, " ")
		value, _ := strconv.Atoi(parts[1])
		switch parts[0] {
		case "down":
			aim += value
		case "forward":
			position += value
			depth += aim * value
		case "up":
			aim -= value
		}
	}
	return depth * position
}
