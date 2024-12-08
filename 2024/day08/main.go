package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type pos struct { row, col int }

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
	antennas := map[byte][]pos{}
	grid := strings.Split(input, "\n")
	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[0]); col++ {
			c := grid[row][col]
			if c >= 'a' && c <= 'z' || c >= 'A' && c <= 'Z' || c >= '0' && c <= '9' {
				antennas[c] = append(antennas[c], pos{row: row, col: col})
			}
		}
	}

	antinodes := map[pos]bool{}
	for _, antenna := range antennas {

		for i := 0; i < len(antenna) - 1; i++ {
			for j := i + 1; j < len(antenna); j++ {
				delta := pos{row: antenna[i].row - antenna[j].row, col: antenna[i].col - antenna[j].col}

				antinode1 := pos{row: antenna[i].row + delta.row, col: antenna[i].col + delta.col}
				antinode2 := pos{row: antenna[j].row - delta.row, col: antenna[j].col - delta.col}
				if antinode1.row >= 0 && antinode1.row < len(grid) && antinode1.col >= 0 && antinode1.col < len(grid[0]) {
					antinodes[antinode1] = true
				}
				if antinode2.row >= 0 && antinode2.row < len(grid) && antinode2.col >= 0 && antinode2.col < len(grid[0]) {
					antinodes[antinode2] = true
				}

			}
		}
	}

    return len(antinodes)
}
func partTwo(input string) int{
	antennas := map[byte][]pos{}
	grid := strings.Split(input, "\n")
	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[0]); col++ {
			c := grid[row][col]
			if c >= 'a' && c <= 'z' || c >= 'A' && c <= 'Z' || c >= '0' && c <= '9' {
				antennas[c] = append(antennas[c], pos{row: row, col: col})
			}
		}
	}

	antinodes := map[pos]bool{}
	for _, antenna := range antennas {

		for i := 0; i < len(antenna) - 1; i++ {
			for j := i + 1; j < len(antenna); j++ {
				delta := pos{row: antenna[i].row - antenna[j].row, col: antenna[i].col - antenna[j].col}

				antinode1 := pos{row: antenna[i].row, col: antenna[i].col}
				for {
					if antinode1.row < 0 || antinode1.row >= len(grid) || antinode1.col < 0 || antinode1.col >= len(grid[0]) {
						break
					}
					antinodes[antinode1] = true
					antinode1.row += delta.row
					antinode1.col += delta.col

				}

				antinode2 := pos{row: antenna[j].row, col: antenna[j].col}
				for {
					if antinode2.row < 0 || antinode2.row >= len(grid) || antinode2.col < 0 || antinode2.col >= len(grid[0]) {
						break
					}
					antinodes[antinode2] = true
					antinode2.row -= delta.row
					antinode2.col -= delta.col
				}
			}
		}
	}

    return len(antinodes)
}
