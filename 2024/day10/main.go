package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
)

type pos struct{row, col int}
func main() {

    inputFile, err := os.ReadFile("input")
    if err != nil {
        log.Fatal(err)
    }
    input := string(inputFile[:len(inputFile)-1])

    fmt.Printf("Part one: %d\n", partOne(input))
    fmt.Printf("Part two: %d\n", partTwo(input))
}

func partOne(input string) int{
    sum := 0

	grid := strings.Split(input, "\n")
	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[0]); col++ {
			if grid[row][col] == '0' {
				sum += len(trailScore(grid, pos{row: row, col: col}))
			}
		}
	}

    return sum
}
func partTwo(input string) int{
    sum := 0

	grid := strings.Split(input, "\n")
	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[0]); col++ {
			if grid[row][col] == '0' {
				sum += trailRating(grid, pos{row: row, col: col})
			}
		}
	}

    return sum
}

func trailScore(grid []string, start pos) (tops []pos) {

	if grid[start.row][start.col] == '9' {
		return []pos{pos{row: start.row, col: start.col}}
	}

	var nexts []pos

	if start.row + 1 < len(grid) {
		if grid[start.row + 1][start.col] == grid[start.row][start.col] + 1 {
			nexts = append(nexts, pos{row: start.row + 1, col: start.col})
		}
	}

	if start.row - 1 >= 0 {
		if grid[start.row - 1][start.col] == grid[start.row][start.col] + 1 {
			nexts = append(nexts, pos{row: start.row - 1, col: start.col})
		}
	}

	if start.col + 1 < len(grid[0]) {
		if grid[start.row][start.col + 1] == grid[start.row][start.col] + 1 {
			nexts = append(nexts, pos{row: start.row, col: start.col + 1})
		}
	}

	if start.col - 1 >= 0 {
		if grid[start.row][start.col - 1] == grid[start.row][start.col] + 1 {
			nexts = append(nexts, pos{row: start.row, col: start.col - 1})
		}
	}

	for _, next := range nexts {
		top := trailScore(grid, next)

		for _, t := range top {
			if !slices.Contains(tops, t) {
				tops = append(tops, t)
			}
		}
	}

	return
}

func trailRating(grid []string, start pos) (rating int) {

	if grid[start.row][start.col] == '9' {
		return 1
	}

	var nexts []pos

	if start.row + 1 < len(grid) {
		if grid[start.row + 1][start.col] == grid[start.row][start.col] + 1 {
			nexts = append(nexts, pos{row: start.row + 1, col: start.col})
		}
	}

	if start.row - 1 >= 0 {
		if grid[start.row - 1][start.col] == grid[start.row][start.col] + 1 {
			nexts = append(nexts, pos{row: start.row - 1, col: start.col})
		}
	}

	if start.col + 1 < len(grid[0]) {
		if grid[start.row][start.col + 1] == grid[start.row][start.col] + 1 {
			nexts = append(nexts, pos{row: start.row, col: start.col + 1})
		}
	}

	if start.col - 1 >= 0 {
		if grid[start.row][start.col - 1] == grid[start.row][start.col] + 1 {
			nexts = append(nexts, pos{row: start.row, col: start.col - 1})
		}
	}

	for _, next := range nexts {
		rating += trailRating(grid, next)
	}

	return rating
}
