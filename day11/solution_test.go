package day11_test

import (
	"github.com/svfoxat/aoc24/day11"
	"testing"
)

func TestSimple(t *testing.T) {
	dataString := "0 1 10 99 999"
	result := day11.Run([]byte(dataString), 1)

	desired := []int64{1, 2024, 1, 0, 9, 9, 2021976}

	if len(desired) != int(result) {
		t.Fatalf("result=%d, desired=%d", result, len(desired))
	}

}

func TestLonger(t *testing.T) {
	dataString := "125 17"
	result := day11.Run([]byte(dataString), 6)

	desired := []int64{2097446912, 14168, 4048, 2, 0, 2, 4, 40, 48, 2024, 40, 48, 80, 96, 2, 8, 6, 7, 6, 0, 3, 2}
	if len(desired) != int(result) {
		t.Fatalf("result=%d, desired=%d", result, len(desired))
	}
}

func TestPart1(t *testing.T) {
	dataString := "64554 35 906 6 6960985 5755 975820 0"
	result := day11.Run([]byte(dataString), 25)
	if 175006 != result {
		t.Fatalf("result=%d, desired=%d", result, 175006)
	}
}

func TestPart2(t *testing.T) {
	dataString := "64554 35 906 6 6960985 5755 975820 0"
	result := day11.Run([]byte(dataString), 75)
	if 207961583799296 != result {
		t.Fatalf("result=%d, desired=%d", result, 207961583799296)
	}
}

func BenchmarkPart1(b *testing.B) {
	dataString := "64554 35 906 6 6960985 5755 975820 0"
	for i := 0; i < b.N; i++ {
		day11.Run([]byte(dataString), 25)
	}
}
func BenchmarkPart2(b *testing.B) {
	dataString := "64554 35 906 6 6960985 5755 975820 0"
	for i := 0; i < b.N; i++ {
		day11.Run([]byte(dataString), 75)
	}
}
