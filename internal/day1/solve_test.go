package day1_test

import (
	"path/filepath"
	"testing"

	"github.com/andygeiss/aoc2021/internal/day1"
	"github.com/andygeiss/aoc2021/pkg/assert"
	"github.com/andygeiss/aoc2021/pkg/file"
)

func Test_Solve_Should_Return_0_Given_No_Input_Data(t *testing.T) {
	result := day1.SolveOne([]string{})
	assert.That(t, result, 0)
}

func Test_Solve_Should_Return_0_Given_One_Value(t *testing.T) {
	result := day1.SolveOne([]string{"199"})
	assert.That(t, result, 0)
}

func Test_Solve_Should_Return_1_Given_Two_Values(t *testing.T) {
	result := day1.SolveOne([]string{"199", "200"})
	assert.That(t, result, 1)
}

func Test_Solve_Should_Return_2_Given_Three_Values(t *testing.T) {
	result := day1.SolveOne([]string{"199", "200", "208"})
	assert.That(t, result, 2)
}

func Test_Solve_Should_Return_2_Given_Four_Values(t *testing.T) {
	result := day1.SolveOne([]string{"199", "200", "208", "201"})
	assert.That(t, result, 2)
}

func Test_Solve_Should_Return_1195_Given_Example_Data(t *testing.T) {
	input := file.Parse(filepath.Join("testdata", "example.txt"))
	out := day1.SolveOne(input)
	assert.That(t, out, 7)
}

func Test_Solve_Should_Return_1195_Given_Input_Data(t *testing.T) {
	input := file.Parse(filepath.Join("testdata", "input.txt"))
	out := day1.SolveOne(input)
	assert.That(t, out, 1195)
}

func Test_SolveTwo_Should_Return_1_Given_Four_Values(t *testing.T) {
	result := day1.SolveTwo([]string{"199", "200", "208", "210"})
	assert.That(t, result, 1)
}

func Test_SolveTwo_Should_Return_1_Given_Five_Values(t *testing.T) {
	result := day1.SolveTwo([]string{"199", "200", "208", "210", "200"})
	assert.That(t, result, 1)
}

func Test_SolveTwo_Should_Return_1_Given_Six_Values(t *testing.T) {
	result := day1.SolveTwo([]string{"199", "200", "208", "210", "200", "207"})
	assert.That(t, result, 1)
}

func Test_SolveTwo_Should_Return_2_Given_Seven_Values(t *testing.T) {
	result := day1.SolveTwo([]string{"199", "200", "208", "210", "200", "207", "240"})
	assert.That(t, result, 2)
}

func Test_SolveTwo_Should_Return_1235_Given_Input_Data(t *testing.T) {
	input := file.Parse(filepath.Join("testdata", "input.txt"))
	out := day1.SolveTwo(input)
	assert.That(t, out, 1235)
}
