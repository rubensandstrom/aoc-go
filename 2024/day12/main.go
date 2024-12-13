package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type pos struct { row, col int}
type dir struct { row, col int}

var delta = []dir{
	{0, -1}, {-1, 0}, {0, 1}, {1, 0},
}

func main() {

    inputFile, err := os.ReadFile("input")
    if err != nil {
        log.Fatal(err)
    }
    input := string(inputFile[:len(inputFile)-1])

	garden := strings.Split(input, "\n")
    fmt.Printf("Part one: %d\n", partOne(garden))
    fmt.Printf("Part two: %d\n", partTwo(garden))
}

func partOne(garden []string) int{
    sum := 0
	visited := map[pos]bool{}

	for row := 0; row < len(garden); row++ {
		for col := 0; col < len(garden[0]); col++ {
			if !visited[pos{row, col}] {
				area, border, region := findPerimiter(pos{row, col}, garden)

				for key := range region {
					visited[key] = true
				}

				sum += area * border
			}
		}
	}

    return sum
}

func partTwo(garden []string) int{
    sum := 0
	visited := map[pos]bool{}

	for row := 0; row < len(garden); row++ {
		for col := 0; col < len(garden[0]); col++ {
			if !visited[pos{row, col}] {
				area, border, region := findSides(pos{row, col}, garden)

				for key := range region {
					visited[key] = true
				}

				sum += area * border
			}
		}
	}
    return sum
}

func inBounds(p pos, garden []string) bool {
	return p.row >= 0 && p.row < len(garden) && p.col >= 0 && p.col < len(garden[0])
}

func findPerimiter(p pos, g []string) (area, perimiter int, visited map[pos]bool){

	plant := g[p.row][p.col]
	visited = map[pos]bool{}
	toVisit := []pos{p}

	for len(toVisit) > 0 {
		tmp := toVisit[len(toVisit) - 1]
		toVisit = toVisit[:len(toVisit) - 1]

		for _, d := range delta {
			neighbor := pos{row: tmp.row + d.row, col: tmp.col + d.col }
			if inBounds(neighbor, g) && g[neighbor.row][neighbor.col] == plant {
				if !visited[neighbor]  {
					toVisit = append(toVisit, neighbor)
					visited[neighbor] = true
				}
			} else {
				perimiter++
			}
		}
		visited[tmp] = true
		area++
	}
	return
}

func findSides(p pos, g []string) (area, sides int, visited map[pos]bool) {

	plant := g[p.row][p.col]
	visited = map[pos]bool{}
	toVisit := []pos{p}

	for len(toVisit) > 0 {
		tmp := toVisit[len(toVisit) - 1]
		toVisit = toVisit[:len(toVisit) - 1]

		for _, d := range delta {
			neighbor := pos{row: tmp.row + d.row, col: tmp.col + d.col }
			if inBounds(neighbor, g) && g[neighbor.row][neighbor.col] == plant {
				if !visited[neighbor]  {
					toVisit = append(toVisit, neighbor)
					visited[neighbor] = true
				}
			} 
		}
		visited[tmp] = true
		area++

		window := [3][3]bool{}

		for i := -1; i <= 1; i++ {
			for j := -1; j <= 1; j++ {
				neighbor := pos{row: tmp.row + i, col: tmp.col + j}
				window[i + 1][j + 1] = inBounds(neighbor, g) && g[neighbor.row][neighbor.col] == plant
			}
		}

		// Inner corners
		if window[1][0] && window[0][1] && !window[0][0] {
			sides++
		}

		if window[0][1] && window[1][2] && !window[0][2] {
			sides++
		}

		if window[1][2] && window[2][1] && !window[2][2] {
			sides++
		}

		if window[2][1] && window[1][0] && !window[2][0] {
			sides++
		}

		// Outer corners
		if !window[1][0] && !window[0][1] && !window[1][2]  && !window[2][1] {
			sides += 4
		} else if !window[1][0] && !window[0][1] && !window[1][2] || 
		          !window[0][1] && !window[1][2] && !window[2][1] || 
		          !window[1][2] && !window[2][1] && !window[1][0] || 
		          !window[2][1] && !window[1][0] && !window[0][1] {
			sides += 2
		} else if !window[1][0] && !window[0][1] || !window[0][1] && !window[1][2] || 
		          !window[1][2] && !window[2][1] || !window[2][1] && !window[1][0] {
			sides += 1
		}
	}

	return
}
