package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

type dir struct {x, y int}
type pos struct {x, y int}

var deltas = []dir{ {0, 1}, {1, 0}, {0, -1}, {-1, 0} }

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
	maze := strings.Split(input, "\n")
	start := pos{}
	for y, row := range maze {
		for x, col := range row {
			if col == 'S' {
				start = pos{y: y, x: x}
			}
		}
	}
	return race(start, dir{x: 1, y: 0}, 0, map[pos]int{}, maze, math.MaxInt)
}

func partTwo(input string) int{
	maze := strings.Split(input, "\n")
	start := pos{}
	for y, row := range maze {
		for x, col := range row {
			if col == 'S' {
				start = pos{y: y, x: x}
			}
		}
	}
	seatMap := map[pos]bool{}
	seats(start, dir{x: 1, y: 0}, 0, map[pos]int{}, maze, math.MaxInt, []pos{}, &seatMap)
	// fmt.Printf("%v", seatMap)
	return len(seatMap)
}

func abs(a int) int {
	if a < 0 { return - a }
	return a
}

func race( cp pos, cd dir, cs int, visited map[pos]int, maze []string, minScore int) (score int) {
	if cs > minScore {return cs }
	if maze[cp.y][cp.x] == 'E' {
		if cs < minScore {
			minScore = cs
		}
		return cs
	}
	if val, ok := visited[cp]; ok && val + 1000 < cs {
		return math.MaxInt
	} else {
		visited[cp] = cs
	}
	for _, d := range deltas {
		if d.x + cd.x == 0 && d.y + cd.y == 0 { continue }
		newPos := pos{x: cp.x + d.x, y: cp.y + d.y}
		if maze[newPos.y][newPos.x] == '#' {continue}

		s := 0
		if d.x == cd.x && d.y == cd.y {
			s = race(newPos, d, cs + 1, visited, maze, minScore)
		} else {
			s = race(newPos, d, cs + 1001, visited, maze, minScore)
		}
		if s < minScore { minScore = s }
	}
	return minScore
} 


func seats( cp pos, cd dir, cs int, visited map[pos]int, maze []string, minScore int, path []pos, seatMap *map[pos]bool) (score int) {

	path = append(path, cp)

	if cs > minScore { return cs }
	if maze[cp.y][cp.x] == 'E' {
		if cs < minScore {
			*seatMap = map[pos]bool{}
			minScore = cs
		}

		if cs == minScore {
			for _, p := range path {
				(*seatMap)[p] = true
			}
		}
		return cs
	}
	if val, ok := visited[cp]; ok && val + 1000 < cs {
		return math.MaxInt
	} else {
		visited[cp] = cs
	}
	for _, d := range deltas {
		if d.x + cd.x == 0 && d.y + cd.y == 0 { continue }
		newPos := pos{x: cp.x + d.x, y: cp.y + d.y}
		if maze[newPos.y][newPos.x] == '#' {continue}

		s := 0
		if d.x == cd.x && d.y == cd.y {
			s = seats(newPos, d, cs + 1, visited, maze, minScore, path, seatMap)
		} else {
			s = seats(newPos, d, cs + 1001, visited, maze, minScore, path, seatMap)
		}
		if s < minScore { minScore = s }
	}
	return minScore
} 
