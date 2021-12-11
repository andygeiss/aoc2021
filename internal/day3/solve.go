package day3

import (
	"math"

	"github.com/andygeiss/aoc2021/pkg/slice"
)

// --- Part One ---

// BitMask ...
func BitMask(len int) (res uint) {
	return uint(math.Pow(2, float64(len))) - 1
}

// CommonBit ...
func CommonBit(s string) (bit uint) {
	count := 0
	n := len(s)
	for i := 0; i < n; i++ {
		if s[i:i+1] == "1" {
			count++
		}
	}
	if count > n/2 {
		return 1
	}
	return 0
}

// Epsilon ...
func Epsilon(s []string) (res uint) {
	return Gamma(s) ^ BitMask(len(s[0]))
}

// Gamma ...
func Gamma(s []string) (res uint) {
	res = 0
	n := len(s[0])
	for i := 0; i < n; i++ {
		val := slice.StringsVertical(s, i)
		bit := CommonBit(val)
		res += bit * uint(math.Pow(2, float64(n-1-i)))
	}
	return res
}

// PowerConsumption ...
func PowerConsumption(s []string) (res uint) {
	return Gamma(s) * Epsilon(s)
}
