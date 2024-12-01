package day1

import (
	"cmp"
	"log/slog"
	"os"
	"slices"
	"strconv"
	"strings"
)

func Part1() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(input), "\n")
	var left, right []int

	for _, v := range lines {
		numbers := strings.Split(v, "   ")
		if len(numbers) != 2 {
			break
		}

		leftInt, err := strconv.ParseInt(numbers[0], 10, 32)
		if err != nil {
			panic(err)
		}

		rightInt, err := strconv.ParseInt(numbers[1], 10, 32)
		if err != nil {
			panic(err)
		}

		left = append(left, int(leftInt))
		right = append(right, int(rightInt))
	}

	slices.SortFunc(left, func(a, b int) int {
		return cmp.Compare(a, b)
	})
	slices.SortFunc(right, func(a, b int) int {
		return cmp.Compare(a, b)
	})

	totalDistance := 0
	for i := range left {
		d := left[i] - right[i]
		if d < 0 {
			d = d * -1
		}
		totalDistance += d
	}

	slog.Info("finished", "result", totalDistance)
}

func Part2() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(input), "\n")
	var left, right []int64
	histogram := map[int64]int64{}

	for _, v := range lines {
		numbers := strings.Split(v, "   ")
		if len(numbers) != 2 {
			break
		}
		leftInt, err := strconv.ParseInt(numbers[0], 10, 0)
		if err != nil {
			panic(err)
		}

		rightInt, err := strconv.ParseInt(numbers[1], 10, 0)
		if err != nil {
			panic(err)
		}

		left = append(left, leftInt)
		right = append(right, rightInt)

		val, ok := histogram[rightInt]
		if !ok {
			histogram[rightInt] = 0
		}

		histogram[rightInt] = val + 1
	}

	var score int64
	score = 0

	for _, v := range left {
		c, ok := histogram[v]
		if !ok {
			continue
		}
		if c == 0 {
			continue
		}

		score += v * c
	}

	slog.Info("finished", "result", score)
}
