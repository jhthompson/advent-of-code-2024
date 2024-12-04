package three

import (
	"log/slog"
	"regexp"
	"strconv"
)

func SolvePartOne(lines []string) int {
	sum := 0
	product := 0
	re, _ := regexp.Compile(`mul\((?<X>\d+),(?<Y>\d+)\)`)

	for _, line := range lines {
		matches := re.FindAllStringSubmatch(line, -1)

		for _, match := range matches {
			X, _ := strconv.Atoi(match[re.SubexpIndex("X")])
			Y, _ := strconv.Atoi(match[re.SubexpIndex("Y")])

			product = X * Y
			sum += product
		}
	}

	slog.Info("03.1", "result", sum)
	return sum
}

func SolvePartTwo(lines []string) int {
	sum := 0
	product := 0
	enabled := true

	re, _ := regexp.Compile(`mul\((?<X>\d+),(?<Y>\d+)\)|(?<DO>do\(\))|(?<DONT>don't\(\))`)

	for _, line := range lines {
		matches := re.FindAllStringSubmatch(line, -1)

		for _, match := range matches {
			X, _ := strconv.Atoi(match[re.SubexpIndex("X")])
			Y, _ := strconv.Atoi(match[re.SubexpIndex("Y")])

			do := match[re.SubexpIndex("DO")]
			dont := match[re.SubexpIndex("DONT")]

			if do != "" {
				enabled = true
			} else if dont != "" {
				enabled = false
			}

			if enabled {
				product = X * Y
				sum += product
			}
		}
	}

	slog.Info("03.2", "result", sum)
	return sum
}
