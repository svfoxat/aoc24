package day2

import (
	"log/slog"
	"os"
	"strconv"
	"strings"
)

func Part1() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(data), "\n")
	safeLines := 0
	dampednedSafeLines := 0

	for _, v := range lines {
		if len(v) == 0 {
			break
		}
		numberStrings := strings.Split(v, " ")
		numbers := []int64{}

		for _, v := range numberStrings {
			n, err := strconv.ParseInt(v, 10, 0)
			if err != nil {
				panic(err)
			}
			numbers = append(numbers, n)
		}

		isSafe := isLineSafe(numbers)
		if isSafe {
			safeLines++
		}

		if dampen(numbers) {
			dampednedSafeLines++
		}
	}
	slog.Info("finished", "result", safeLines, "with damping", dampednedSafeLines)
}

func dampen(numbers []int64) bool {
	for i := 0; i < len(numbers); i++ {
		temp := make([]int64, len(numbers))
		copy(temp, numbers)
		if isLineSafe(RemoveIndex(temp, i)) {
			return true
		}
	}
	return false
}

func isLineSafe(numbers []int64) bool {
	increasing := numbers[0] < numbers[1]
	errs := []int64{}

	for i := 0; i <= len(numbers)-1; i++ {
		if i == len(numbers)-1 {
			break
		}
		safe := isPairSafe(numbers[i], numbers[i+1], increasing)
		if !safe {
			errs = append(errs, int64(i))
		}
	}

	return len(errs) == 0
}

func isPairSafe(a, b int64, increasing bool) bool {
	if (a < b) != increasing {
		return false
	}

	delta := a - b
	if delta < 0 {
		delta *= -1
	}

	if delta < 1 || delta > 3 {
		return false
	}

	return true
}

func RemoveIndex(s []int64, index int) []int64 {
	return append(s[:index], s[index+1:]...)
}
