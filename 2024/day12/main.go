package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type pos struct { row, col int}
type dir struct { row, col int}



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
				area, region := findArea(pos{row, col}, garden)
				perimiter := findPerimiter(pos{row, col}, garden)

				for key := range region {
					visited[key] = true
				}

				sum += area * perimiter
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
				area, region := findArea(pos{row, col}, garden)
				sides := findSides(pos{row, col}, garden)

				for key := range region {
					visited[key] = true
				}

				sum += area * sides
			}
		}
	}

    return sum
}

func findArea(p pos, g []string) (area int, visited map[pos]bool ){
	sum := 0

	plant := g[p.row][p.col]
	visited = map[pos]bool{}

	toVisit := []pos{}
	toVisit = append(toVisit, p)

	for len(toVisit) > 0 {
		tmp := toVisit[len(toVisit) - 1]
		toVisit = toVisit[:len(toVisit) - 1]

		if tmp.col - 1 >= 0 {
			if !visited[pos{tmp.row, tmp.col - 1}] && g[tmp.row][tmp.col - 1] == plant {
				toVisit = append(toVisit, pos{tmp.row, tmp.col - 1})
				visited[pos{tmp.row, tmp.col - 1}] = true
			}
		}

		if tmp.col + 1 < len(g[0]) {
			if !visited[pos{tmp.row, tmp.col + 1}] && g[tmp.row][tmp.col + 1] == plant {
				toVisit = append(toVisit, pos{tmp.row, tmp.col + 1})
				visited[pos{tmp.row, tmp.col + 1}] = true
			}
		}

		if tmp.row - 1 >= 0 {
			if !visited[pos{tmp.row -1, tmp.col}] && g[tmp.row - 1][tmp.col] == plant {
				toVisit = append(toVisit, pos{tmp.row - 1, tmp.col})
				visited[pos{tmp.row -1, tmp.col}] = true
			}
		}

		if tmp.row + 1 < len(g) {
			if !visited[pos{tmp.row + 1, tmp.col}] && g[tmp.row + 1][tmp.col] == plant {
				toVisit = append(toVisit, pos{tmp.row + 1, tmp.col})
				visited[pos{tmp.row + 1, tmp.col}] = true
			}
		}
		visited[tmp] = true
		sum++
	}

	return sum, visited
}

func findPerimiter(p pos, g []string) (perimiter int) {
	sum := 0

	plant := g[p.row][p.col]
	visited := map[pos]bool{}

	toVisit := []pos{p}

	for len(toVisit) > 0 {
		tmpSum := 0
		tmp := toVisit[len(toVisit) - 1]
		toVisit = toVisit[:len(toVisit) - 1]

		if tmp.col - 1 >= 0 {
			if !visited[pos{tmp.row, tmp.col - 1}] {
				if g[tmp.row][tmp.col - 1] == plant {
					toVisit = append(toVisit, pos{tmp.row, tmp.col - 1})
					visited[pos{tmp.row, tmp.col - 1}] = true
				} else {
					tmpSum++
				}
			}
		} else {
			tmpSum++
		}

		if tmp.col + 1 < len(g[0]) {
			if !visited[pos{tmp.row, tmp.col + 1}] {
				if g[tmp.row][tmp.col + 1] == plant {
					toVisit = append(toVisit, pos{tmp.row, tmp.col + 1})
					visited[pos{tmp.row, tmp.col + 1}] = true
				} else {
					tmpSum ++
				}
			}
		} else {
			tmpSum++
		}

		if tmp.row - 1 >= 0 {
			if !visited[pos{tmp.row -1, tmp.col}] {
			    if g[tmp.row - 1][tmp.col] == plant {
					toVisit = append(toVisit, pos{tmp.row - 1, tmp.col})
					visited[pos{tmp.row -1, tmp.col}] = true
				} else {
					tmpSum++
				}
			}
		} else {
			tmpSum++
		}

		if tmp.row + 1 < len(g) {
			if !visited[pos{tmp.row + 1, tmp.col}] {
				if g[tmp.row + 1][tmp.col] == plant {
					toVisit = append(toVisit, pos{tmp.row + 1, tmp.col})
					visited[pos{tmp.row + 1, tmp.col}] = true
				} else {
					tmpSum++
				}
			}
		} else {
			tmpSum++
		}

		sum += tmpSum
		visited[tmp] = true
	}

	return sum
}


func findSides(p pos, g []string) (perimiter int) {
	sum := 0

	plant := g[p.row][p.col]
	visited := map[pos]bool{}

	toVisit := []pos{p}


	for len(toVisit) > 0 {

		left      := false
		right     := false
		up        := false
		down      := false
		upLeft    := false
		upRight   := false
		downLeft  := false
		downRight := false

		tmp := toVisit[len(toVisit) - 1]
		toVisit = toVisit[:len(toVisit) - 1]

		if tmp.col - 1 >= 0 {
			if !visited[pos{tmp.row, tmp.col - 1}] {
				if g[tmp.row][tmp.col - 1] == plant {
					toVisit = append(toVisit, pos{tmp.row, tmp.col - 1})
					visited[pos{tmp.row, tmp.col - 1}] = true
				} else {
					left = true
				}			
			}
		} else {
			left = true
		}

		if tmp.col + 1 < len(g[0]) {
			if !visited[pos{tmp.row, tmp.col + 1}] {
				if g[tmp.row][tmp.col + 1] == plant {
					toVisit = append(toVisit, pos{tmp.row, tmp.col + 1})
					visited[pos{tmp.row, tmp.col + 1}] = true
				} else {
					right = true
				}
			}
		} else {
			right = true
		}

		if tmp.row - 1 >= 0 {
			if !visited[pos{tmp.row -1, tmp.col}] {
			    if g[tmp.row - 1][tmp.col] == plant {
					toVisit = append(toVisit, pos{tmp.row - 1, tmp.col})
					visited[pos{tmp.row -1, tmp.col}] = true
				} else {
					up = true
				}
			}
		} else {
			up = true
		}

		if tmp.row + 1 < len(g) {
			if !visited[pos{tmp.row + 1, tmp.col}] {
				if g[tmp.row + 1][tmp.col] == plant {
					toVisit = append(toVisit, pos{tmp.row + 1, tmp.col})
					visited[pos{tmp.row + 1, tmp.col}] = true
				} else {
					down = true
				}
			}
		} else {
			down = true
		}

		visited[tmp] = true

		if tmp.row - 1 >= 0 && tmp.col - 1 >= 0 {
			upLeft = g[tmp.row - 1][tmp.col - 1] != plant
		} else {
			upLeft = true
		}

		if tmp.row - 1 >= 0 && tmp.col + 1 < len(g[0]) {
			upRight = g[tmp.row - 1][tmp.col + 1] != plant
		} else {
			upRight = true
		}

		if tmp.row + 1 < len(g) && tmp.col - 1 >= 0 {
			downLeft = g[tmp.row + 1][tmp.col - 1] != plant
		} else {
			downLeft = true
		}

		if tmp.row + 1 < len(g) && tmp.col + 1 < len(g[0]) {
			downRight = g[tmp.row + 1][tmp.col + 1] != plant
		} else {
			downRight = true
		}


		if !left && !up && upLeft {
			sum++
		}

		if !up && !right && upRight {
			sum++
		}

		if !right && !down && downRight {
			sum++
		}

		if !down && !left && downLeft {
			sum++
		}

		if left && up && right  && down {
			sum += 4
		} else if left && up && right || up && right && down || right && down && left || down && left && up {
			sum += 2
		} else if left && up || up && right || right && down || down && left {
			sum += 1
		}
	}

	return sum
}
