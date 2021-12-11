package day3_test

import (
	"path/filepath"
	"testing"

	"github.com/andygeiss/aoc2021/internal/day3"
	"github.com/andygeiss/aoc2021/pkg/assert"
	"github.com/andygeiss/aoc2021/pkg/file"
)

// --- Part One ---

func Test_BitMask_Should_Return_1_Given_1(t *testing.T) {
	result := day3.BitMask(1)
	assert.That(t, result, 1)
}

func Test_BitMask_Should_Return_3_Given_2(t *testing.T) {
	result := day3.BitMask(2)
	assert.That(t, result, 3)
}

func Test_BitMask_Should_Return_7_Given_3(t *testing.T) {
	result := day3.BitMask(3)
	assert.That(t, result, 7)
}

func Test_BitMask_Should_Return_15_Given_4(t *testing.T) {
	result := day3.BitMask(4)
	assert.That(t, result, 15)
}

func Test_CommonBit_Should_Return_1_Given_1(t *testing.T) {
	result := day3.CommonBit("1")
	assert.That(t, result, 1)
}

func Test_CommonBit_Should_Return_1_Given_101(t *testing.T) {
	result := day3.CommonBit("101")
	assert.That(t, result, 1)
}

func Test_CommonBit_Should_Return_0_Given_1001010(t *testing.T) {
	result := day3.CommonBit("1001010")
	assert.That(t, result, 0)
}

func Test_Epsilon_Should_Return_0_Given_0_0_1(t *testing.T) {
	result := day3.Epsilon([]string{"0", "0", "1"})
	assert.That(t, result, 1)
}

func Test_Epsilon_Should_Return_0_Given_1_1_0(t *testing.T) {
	result := day3.Epsilon([]string{"1", "1", "0"})
	assert.That(t, result, 0)
}

func Test_Gamma_Should_Return_0_Given_0_0_1(t *testing.T) {
	result := day3.Gamma([]string{"0", "0", "1"})
	assert.That(t, result, 0)
}

func Test_Gamma_Should_Return_0_Given_1_1_0(t *testing.T) {
	result := day3.Gamma([]string{"1", "1", "0"})
	assert.That(t, result, 1)
}

func Test_Gamma_Should_Return_2_Given_10_10_00(t *testing.T) {
	result := day3.Gamma([]string{"10", "10", "00"})
	assert.That(t, result, 2)
}

func Test_Gamma_Should_Return_6_Given_110_110_100(t *testing.T) {
	result := day3.Gamma([]string{"110", "110", "100"})
	assert.That(t, result, 6)
}

func Test_PowerConsumption_Should_Return_0_Given_0_0_1(t *testing.T) {
	result := day3.PowerConsumption([]string{"0", "0", "1"})
	assert.That(t, result, 0)
}

func Test_PowerConsumption_Should_Return_2_Given_10_10_00(t *testing.T) {
	result := day3.PowerConsumption([]string{"10", "10", "00"})
	assert.That(t, result, 2)
}

func Test_PowerConsumption_Should_Return_198_Given_Example_Data(t *testing.T) {
	input := file.Parse(filepath.Join("testdata", "example.txt"))
	result := day3.PowerConsumption(input)
	assert.That(t, result, 198)
}

func Test_PowerConsumption_Should_Return_852500_Given_Input_Data(t *testing.T) {
	input := file.Parse(filepath.Join("testdata", "input.txt"))
	result := day3.PowerConsumption(input)
	assert.That(t, result, 852500)
}
