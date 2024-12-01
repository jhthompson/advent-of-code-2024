package one

import (
	"log/slog"
	"slices"
	"strconv"
	"strings"
)

func SolvePartOne(lines []string) int {
	left, right := readData(lines)
	slices.Sort(left)
	slices.Sort(right)

	total := 0

	for i, l := range left {
		r := right[i]

		diff := abs(l - r)
		slog.Debug("01.1", "left", l, "right", r, "diff", diff)
		total += diff
	}

	slog.Info("01.1", "result", total)
	return total
}

func SolvePartTwo(lines []string) int {
	left, right := readData(lines)

	similarity := 0

	counter := make(map[int]int)
	for _, c := range right {
		counter[c] += 1
	}

	for _, v := range left {
		score := 0

		score = v * counter[v]

		slog.Debug("01.2", "left number", v, "right count", counter[v], "score", score)
		similarity += score
	}

	slog.Info("01.2", "result", similarity)
	return similarity
}

func readData(lines []string) ([]int, []int) {
	left := make([]int, 0)
	right := make([]int, 0)

	for _, line := range lines {
		data := strings.Fields(line)

		l, err := strconv.Atoi(data[0])
		if err != nil {
			panic(err)
		}

		r, err := strconv.Atoi(data[1])
		if err != nil {
			panic(err)
		}

		left = append(left, l)
		right = append(right, r)
	}

	return left, right
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
