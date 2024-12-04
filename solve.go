package main

import (
	"bufio"
	"log/slog"
	"os"

	one "jhthompson.ca/advent-of-code-2024/01"
	two "jhthompson.ca/advent-of-code-2024/02"
	three "jhthompson.ca/advent-of-code-2024/03"
)

func main() {
	slog.SetLogLoggerLevel(slog.LevelInfo)

	p1 := ReadLines("data/01.in")
	one.SolvePartOne(p1)
	one.SolvePartTwo(p1)

	p2 := ReadLines("data/02.in")
	two.SolvePartOne(p2)
	two.SolvePartTwo(p2)

	p3 := ReadLines("data/03.in")
	three.SolvePartOne(p3)
	three.SolvePartTwo(p3)
}

func ReadLines(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := make([]string, 0)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}
