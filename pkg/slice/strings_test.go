package slice_test

import (
	"testing"

	"github.com/andygeiss/aoc2021/pkg/assert"
	"github.com/andygeiss/aoc2021/pkg/slice"
)

func Test_StringsVertical_Should_Return_0_Given_0(t *testing.T) {
	in := []string{"0"}
	out := slice.StringsVertical(in, 0)
	assert.That(t, out, "0")
}

func Test_StringsVertical_Should_Return_1_Given_1(t *testing.T) {
	in := []string{"1"}
	out := slice.StringsVertical(in, 0)
	assert.That(t, out, "1")
}

func Test_StringsVertical_Should_Return_0_1_Given_0_1(t *testing.T) {
	in := []string{"0", "1"}
	out := slice.StringsVertical(in, 0)
	assert.That(t, out, "01")
}

func Test_StringsVertical_Should_Return_0_1_1_Given_0_1_1(t *testing.T) {
	in := []string{"0", "1", "1"}
	out := slice.StringsVertical(in, 0)
	assert.That(t, out, "011")
}

func Test_StringsVertical_Should_Return_0_1_1_Given_01_11_10(t *testing.T) {
	in := []string{"01", "11", "10"}
	out := slice.StringsVertical(in, 0)
	assert.That(t, out, "011")
}

func Test_StringsVertical_Should_Return_1_1_0_Given_01_11_10(t *testing.T) {
	in := []string{"01", "11", "10"}
	out := slice.StringsVertical(in, 1)
	assert.That(t, out, "110")
}
