package main

import (
	"log/slog"
	"os"
	"strconv"
	"strings"
	"testing"

	one "jhthompson.ca/advent-of-code-2024/01"
	two "jhthompson.ca/advent-of-code-2024/02"
	three "jhthompson.ca/advent-of-code-2024/03"
	four "jhthompson.ca/advent-of-code-2024/04"
)

type Solver func([]string) int

func TestSolve(t *testing.T) {
	slog.SetLogLoggerLevel(slog.LevelDebug)

	verifySolution("data/01.1.sample.in", "data/01.1.sample.out", one.SolvePartOne, t)
	verifySolution("data/01.2.sample.in", "data/01.2.sample.out", one.SolvePartTwo, t)

	verifySolution("data/02.1.sample.in", "data/02.1.sample.out", two.SolvePartOne, t)
	verifySolution("data/02.2.sample.in", "data/02.2.sample.out", two.SolvePartTwo, t)

	verifySolution("data/03.1.sample.in", "data/03.1.sample.out", three.SolvePartOne, t)
	verifySolution("data/03.2.sample.in", "data/03.2.sample.out", three.SolvePartTwo, t)

	verifySolution("data/04.1.sample.in", "data/04.1.sample.out", four.SolvePartOne, t)
	verifySolution("data/04.2.sample.in", "data/04.2.sample.out", four.SolvePartTwo, t)
}

func verifySolution(inputFile string, outputFile string, solve Solver, t *testing.T) {
	t.Helper()

	lines := ReadLines(inputFile)

	got := solve(lines)
	want := readInt(outputFile)

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func readInt(outputFile string) int {
	out, err := os.ReadFile(outputFile)
	if err != nil {
		panic(err)
	}

	want, err := strconv.Atoi(strings.TrimSpace(string(out)))
	if err != nil {
		panic(err)
	}

	return want
}
