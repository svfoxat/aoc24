package day11

import (
	"strconv"
	"strings"
)

func Run(data []byte, blinks int) int64 {
	input := string(data)
	stoneStrings := strings.Split(input, " ")

	var stones = map[int64]int64{}
	for _, s := range stoneStrings {
		i, err := strconv.ParseInt(s, 10, 0)
		if err != nil {
			panic(err)
		}
		stones[i] += 1
	}

	for _ = range blinks {
		stones = blink(stones)
	}
	return stoneCount(stones)
}

func blink(stones map[int64]int64) map[int64]int64 {
	newStones := map[int64]int64{}

	for val, count := range stones {
		if val == 0 {
			newStones[1] += count
			continue
		}
		if intLen(val)%2 == 0 {
			stoneStr := strconv.FormatInt(val, 10)
			lStr := stoneStr[:len(stoneStr)/2]
			rStr := stoneStr[len(stoneStr)/2 : len(stoneStr)]
			l, _ := strconv.ParseInt(lStr, 10, 64)
			r, _ := strconv.ParseInt(rStr, 10, 64)

			newStones[l] += count
			newStones[r] += count
			continue
		}
		newStones[val*2024] += count
	}
	return newStones
}

func stoneCount(m map[int64]int64) int64 {
	var count int64
	count = 0
	for _, v := range m {
		count += v
	}
	return count
}

func intLen(num int64) int {
	return len(strconv.Itoa(int(num)))
}
