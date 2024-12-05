package four

import (
	"log/slog"
)

func SolvePartOne(lines []string) int {
	word := []string{"X", "M", "A", "S"}
	width := 0
	height := 0
	total := 0

	// read data from lines into 2d array
	height = len(lines)
	grid := make([][]string, height)
	for y, line := range lines {
		width = len(line)
		grid[y] = make([]string, width)
		for x, char := range line {
			grid[y][x] = string(char)
		}
	}

	// scan through each element in the grid, and look in each potential direction for the word
	for y, row := range grid {
		for x, char := range row {

			// early exit
			if char != word[0] {
				continue
			}

			// check each possible direction
			// follow each of the following steps until word is found or not
			steps := []struct {
				x int
				y int
			}{
				{1, 0},   // →
				{1, 1},   // ↘
				{0, 1},   // ↓
				{-1, 1},  // ↙
				{-1, 0},  // ←
				{-1, -1}, // ↖
				{0, -1},  // ↑
				{1, -1},  // ↗
			}

			for _, step := range steps {
				// follow the steps until word is found or not
				complete := false
				wordY := y
				wordX := x

				for i, letter := range word {
					// check if we are out of bounds
					if wordX < 0 || wordX >= width || wordY < 0 || wordY >= height {
						break
					}

					// check if the letter is correct
					if grid[wordY][wordX] != letter {
						break
					}

					// check if we're at the end of the word
					if i == len(word)-1 {
						complete = true
						break
					}

					// move to the next step
					wordX += step.x
					wordY += step.y
				}

				if complete {
					total++
				}
			}
		}
	}

	slog.Info("04.1", "result", total)
	return total
}

func SolvePartTwo(lines []string) int {
	width := 0
	height := 0
	total := 0

	// read data from lines into 2d array
	height = len(lines)
	grid := make([][]string, height)
	for y, line := range lines {
		width = len(line)
		grid[y] = make([]string, width)
		for x, char := range line {
			grid[y][x] = string(char)
		}
	}

	// start the search with A
	// then check for:
	// 		M top left, S bottom right || S top left, M bottom right
	// AND	M top right, S bottom left || S top right, M bottom left
	for y, row := range grid {
		for x, char := range row {

			// early exit
			if char != "A" {
				continue
			}

			// check out of bounds
			if x-1 < 0 || x+1 >= width || y-1 < 0 || y+1 >= height {
				continue
			}

			topLeft := grid[y-1][x-1]
			bottomRight := grid[y+1][x+1]
			topRight := grid[y-1][x+1]
			bottomLeft := grid[y+1][x-1]

			if (topLeft == "M" && bottomRight == "S" || topLeft == "S" && bottomRight == "M") && (topRight == "M" && bottomLeft == "S" || topRight == "S" && bottomLeft == "M") {
				total++
			}
		}
	}

	slog.Info("04.2", "result", total)
	return total
}
