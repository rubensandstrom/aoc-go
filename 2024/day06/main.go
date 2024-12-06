package main

import (
	"fmt"
	"log"
	"os"
	"bytes"
)

var pos struct { x, y int }
var dir struct { x, y int }

func main() {

    inputFile, err := os.ReadFile("input")
    if err != nil {
        log.Fatal(err)
    }

    input := inputFile[:len(inputFile)-1]


    fmt.Printf("Part one: %d\n", partOne(input))

    fmt.Printf("Part two: %d\n", partTwo(input))
}

func partOne(input []byte) int{


	dir.y = -1
	dir.x = 0

	area := bytes.Split(input, []byte("\n"))
	for row := 0; row < len(area); row++ {
		for col := 0; col < len(area[0]); col++ {
			if area[row][col] == '^' {
				pos.y = row
				pos.x = col
			}
		}
	}

	beenMap := map[struct{x, y int}]struct{x, y int}{}
	for {
		nextY := pos.y + dir.y
	    nextX := pos.x + dir.x

		beenMap[pos] = dir
		if nextY < 0 || nextY >= len(area) || nextX < 0 || nextX >= len(area[0]) {break}

		for area[nextY][nextX] == '#' {
			rot90(&dir)
			nextY = pos.y + dir.y
			nextX = pos.x + dir.x
		}
		pos.y += dir.y
		pos.x += dir.x
	}
	return len(beenMap)
}

func partTwo(input []byte) int{


	dir.y = -1
	dir.x = 0

	sum := 0
	area := bytes.Split(input, []byte("\n"))

	startY := 0;
	startX := 0;

	for row := 0; row < len(area); row++ {
		for col := 0; col < len(area[0]); col++ {
			if area[row][col] == '^' {
				startY = row
				startX = col
			}
		}
	}

	for i := 0; i < len(area); i++ {
		for j := 0; j < len(area[0]); j++ {	

			if i == startY && j == startX { continue }
			if area[i][j] == '#' { continue }


			pos.y = startY
			pos.x = startX

			dir.y = -1
			dir.x = 0

			beenMap := map[struct{x, y int}]struct{x, y int}{}
			area[i][j] = '#'

			for {
				nextY := pos.y + dir.y
				nextX := pos.x + dir.x

				if val, ok := beenMap[pos]; ok {
					if val == dir {
						sum++
						break
					}
				}

				beenMap[pos] = dir

				if nextY < 0 || nextY >= len(area) || nextX < 0 || nextX >= len(area[0]) {break}

				for area[nextY][nextX] == '#' {
					rot90(&dir)
					nextY = pos.y + dir.y
					nextX = pos.x + dir.x
				}

				pos.y += dir.y
				pos.x += dir.x
			}
			area[i][j] = '.'
		}
	}
	return sum
}

func rot90(dir *struct{x, y int}) {
	if dir.y == -1 && dir.x == 0 { 
		dir.y = 0; dir.x = 1
	} else if dir.y == 0 && dir.x == 1 { 
		dir.y = 1; dir.x = 0
	} else if dir.y == 1 && dir.x == 0 { 
		dir.y = 0; dir.x = -1
	} else if dir.y == 0 && dir.x == -1 { 
		dir.y = -1; dir.x = 0
	}
}
