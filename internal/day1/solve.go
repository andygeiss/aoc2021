package day1

import "strconv"

// SolveOne ...
func SolveOne(input []string) (result int) {
	previous := 0
	for i := 1; i < len(input); i++ {
		value, _ := strconv.Atoi(input[i])
		if value > previous {
			result++
		}
		previous = value
	}
	return
}

// SolveTwo ...
func SolveTwo(input []string) (result int) {
	previous := 0
	for i := 0; i < len(input)-2; i++ {
		value1, _ := strconv.Atoi(input[i])
		value2, _ := strconv.Atoi(input[i+1])
		value3, _ := strconv.Atoi(input[i+2])
		sum := value1 + value2 + value3
		if previous > 0 && sum > previous {
			result++
		}
		previous = sum
	}
	return
}
