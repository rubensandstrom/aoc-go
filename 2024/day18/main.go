package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type pos struct {row, col int}
type dir struct {row, col int}
var deltas = []dir{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

func main() {

    inputFile, err := os.ReadFile("input")
    if err != nil {
        log.Fatal(err)
    }
    input := string(inputFile[:len(inputFile)-1])

    fmt.Printf("Part one: %d\n", partOne(input))
    fmt.Printf("Part two: %s\n", partTwo(input))
}

func inBounds(next pos, grid [][]byte) bool {
	return next.row >= 0 && next.row < len(grid) && next.col >= 0 && next.col < len(grid[0])
}

func solve(grid [][]byte, start, end pos) (cost int, ok bool) {

	costGrid := [][]int{}
	for _, row := range grid {
		tmp := []int{}
		for range row {
			tmp = append(tmp, math.MaxInt)
		}
		costGrid = append(costGrid, tmp)
	}

	costGrid[0][0] = 0

	toVisit := []pos{start}
	for len(toVisit) > 0 {
		current := toVisit[0]
		toVisit = toVisit[1:]

		for _, d := range deltas {
			next := pos{current.row + d.row, current.col + d.col}
			if inBounds(next, grid) && grid[next.row][next.col] != '#' && costGrid[next.row][next.col] > costGrid[current.row][current.col] + 1{
				toVisit = append(toVisit, next)
				costGrid[next.row][next.col] = costGrid[current.row][current.col] + 1
				grid[next.row][next.col] = 'X'
			}
			if next == end { 
				return costGrid[next.row][next.col], true
			}
		}
	}
	return 0, false
}

func partOne(input string) int{
	grid := [][]byte{}

	for i := 0; i <= 70; i++ {
		tmp := []byte{}
		for j := 0; j <= 70; j++ {
			tmp = append(tmp, '.')
			// grid[i][j] = '.'
		}
		grid = append(grid, tmp)
	}

	for i, row := range strings.Split(input, "\n") {
		if i == 1024 {break}
		tmp := strings.Split(row, ",")
		row, _ := strconv.Atoi(tmp[0])
		col, _ := strconv.Atoi(tmp[1])
		grid[row][col] = '#'
	}


	start := pos{0, 0}
	end := pos{70, 70}

	cost, _ := solve(grid, start, end)
	return cost
}

func partTwo(input string) string{
	grid := [][]byte{}

	for i := 0; i <= 70; i++ {
		tmp := []byte{}
		for j := 0; j <= 70; j++ {
			tmp = append(tmp, '.')
			// grid[i][j] = '.'
		}
		grid = append(grid, tmp)
	}

	start := pos{0, 0}
	end := pos{70, 70}

	for _, row := range strings.Split(input, "\n") {
		tmp := strings.Split(row, ",")
		y, _ := strconv.Atoi(tmp[0])
		x, _ := strconv.Atoi(tmp[1])
		grid[y][x] = '#'
		if _, ok := solve(grid, start, end); !ok {
			return row
		}
	}
	return ""
}
