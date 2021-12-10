package day3_test

import (
	"testing"

	"github.com/andygeiss/aoc2021/internal/day3"
	"github.com/andygeiss/aoc2021/pkg/assert"
)

func Test_Epsilon_Should_Return_0_Given_0_0_1(t *testing.T) {
	result := day3.Epsilon([]string{"0", "0", "1"})
	assert.That(t, result, 1)
}

func Test_Gamma_Should_Return_0_Given_0_0_1(t *testing.T) {
	result := day3.Gamma([]string{"0", "0", "1"})
	assert.That(t, result, 0)
}

func Test_Solve_Should_Return_0_Given_0_0_1(t *testing.T) {
	result := day3.Solve([]string{"0", "0", "1"})
	assert.That(t, result, 0)
}
