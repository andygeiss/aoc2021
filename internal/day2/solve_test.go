package day2_test

import (
	"path/filepath"
	"testing"

	"github.com/andygeiss/aoc2021/internal/day2"
	"github.com/andygeiss/aoc2021/pkg/assert"
	"github.com/andygeiss/aoc2021/pkg/file"
)

func Test_Solve_Should_Return_0_Given_No_Input_Data(t *testing.T) {
	result := day2.Solve([]string{})
	assert.That(t, result, 0)
}

func Test_Solve_Should_Return_0_Given_Forward_5(t *testing.T) {
	result := day2.Solve([]string{"forward 5"})
	assert.That(t, result, 0)
}

func Test_Solve_Should_Return_25_Given_Forward_5_And_Down_5(t *testing.T) {
	result := day2.Solve([]string{"forward 5", "down 5"})
	assert.That(t, result, 25)
}

func Test_Solve_Should_Return_5_Given_Forward_5_Down_5_Up_4(t *testing.T) {
	result := day2.Solve([]string{"forward 5", "down 5", "up 4"})
	assert.That(t, result, 5)
}

func Test_Solve_Should_Return_0_Given_Forward_5_Down_5_Up_5(t *testing.T) {
	result := day2.Solve([]string{"forward 5", "down 5", "up 5"})
	assert.That(t, result, 0)
}

func Test_Solve_Should_Return_150_Given_Example_Data(t *testing.T) {
	input := file.Parse(filepath.Join("testdata", "example.txt"))
	result := day2.Solve(input)
	assert.That(t, result, 150)
}

func Test_Solve_Should_Return_2102357_Given_Input_Data(t *testing.T) {
	input := file.Parse(filepath.Join("testdata", "input.txt"))
	result := day2.Solve(input)
	assert.That(t, result, 2102357)
}

func Test_SolveTwo_Should_Return_0_Given_Forward_5(t *testing.T) {
	result := day2.SolveTwo([]string{"forward 5"})
	assert.That(t, result, 0)
}

func Test_SolveTwo_Should_Return_0_Given_Forward_5_Down_5(t *testing.T) {
	result := day2.SolveTwo([]string{"forward 5", "down 5"})
	assert.That(t, result, 0)
}

func Test_SolveTwo_Should_Return_520_Given_Forward_5_Down_5_Forward_8(t *testing.T) {
	result := day2.SolveTwo([]string{"forward 5", "down 5", "forward 8"})
	assert.That(t, result, 520)
}

func Test_SolveTwo_Should_Return_520_Given_Forward_5_Down_5_Forward_8_Up_3(t *testing.T) {
	result := day2.SolveTwo([]string{"forward 5", "down 5", "forward 8", "up 3"})
	assert.That(t, result, 520)
}

func Test_SolveTwo_Should_Return_900_Given_Example_Data(t *testing.T) {
	input := file.Parse(filepath.Join("testdata", "example.txt"))
	result := day2.SolveTwo(input)
	assert.That(t, result, 900)
}

func Test_SolveTwo_Should_Return_2101031224_Given_Input_Data(t *testing.T) {
	input := file.Parse(filepath.Join("testdata", "input.txt"))
	result := day2.SolveTwo(input)
	assert.That(t, result, 2101031224)
}
