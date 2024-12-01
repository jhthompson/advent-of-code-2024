package main

import (
	"log/slog"
	"os"
	"strconv"
	"strings"
	"testing"

	one "jhthompson.ca/advent-of-code-2024/01"
)

type solver func([]string) int

func TestSolve(t *testing.T) {
	slog.SetLogLoggerLevel(slog.LevelDebug)

	verifySolution("data/01.1.sample.in", "data/01.1.sample.out", one.SolvePartOne, t)
	verifySolution("data/01.2.sample.in", "data/01.2.sample.out", one.SolvePartTwo, t)
}

func verifySolution(inputFile string, outputFile string, solve solver, t *testing.T) {
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
