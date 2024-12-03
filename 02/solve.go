package two

import (
	"log/slog"
	"strconv"
	"strings"
)

func SolvePartOne(lines []string) int {
	reports := readData(lines)

	totalSafe := 0

	for _, levels := range reports {
		safe := isSafe(levels, 0)

		slog.Debug("02.1", "levels", levels, "safe", safe)
		if safe {
			totalSafe++
		}
	}

	slog.Info("02.1", "result", totalSafe)
	return totalSafe
}

// tried:
//   - 314
func SolvePartTwo(lines []string) int {
	reports := readData(lines)

	totalSafe := 0

	for _, levels := range reports {
		safe := isSafe(levels, 1)

		slog.Debug("02.2", "levels", levels, "safe", safe)
		if safe {
			totalSafe++
		}
	}

	slog.Info("02.2", "result", totalSafe)
	return totalSafe
}

func readData(lines []string) [][]int {
	reports := make([][]int, 0)

	for _, line := range lines {
		data := strings.Fields(line)
		levels := make([]int, 0)

		for _, d := range data {
			l, err := strconv.Atoi(d)
			if err != nil {
				panic(err)
			}

			levels = append(levels, l)
		}

		reports = append(reports, levels)
	}

	return reports
}

func isSafe(levels []int, tolerance int) bool {
	slog.Debug("02.isSafe", "levels", levels, "tolerance", tolerance)

	increasing := levels[0] < levels[1]
	decreasing := levels[0] > levels[1]

	if !increasing && !decreasing {
		if tolerance > 0 {
			return isSafe(RemoveIndex(levels, 0), tolerance-1)
		} else {
			return false
		}
	}

	const MinDiff = 1
	const MaxDiff = 3

	for i := 1; i < len(levels); i++ {
		invalidStep := false

		if increasing {
			invalidStep = levels[i-1] > levels[i] || levels[i]-levels[i-1] < MinDiff || levels[i]-levels[i-1] > MaxDiff
		} else if decreasing {
			invalidStep = levels[i-1] < levels[i] || levels[i-1]-levels[i] < MinDiff || levels[i-1]-levels[i] > MaxDiff
		}

		// if we encounter an error, try solving with new slices with potential problematic element(s) removed
		if invalidStep {
			if tolerance > 0 {
				first := RemoveIndex(levels, i-1)
				second := RemoveIndex(levels, i)

				// this is for when we have something like 1 0 1 2 3
				// the previous two cases will try:
				// 1 0 2 3
				// 1 1 2 3
				// neither of which are valid, but we can make a valid one:
				// 0 1 2 3
				third := []int{1, 1, 1}
				if i-2 >= 0 {
					third = RemoveIndex(levels, i-2)
				}

				return isSafe(first, tolerance-1) || isSafe(second, tolerance-1) || isSafe(third, tolerance-1)
			} else {
				return false
			}
		}
	}

	return true
}

func RemoveIndex(s []int, index int) []int {
	ret := make([]int, 0)
	ret = append(ret, s[:index]...)
	return append(ret, s[index+1:]...)
}
