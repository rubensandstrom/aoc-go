package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
)

type pos struct { x, y int }
func main() {

    inputFile, err := os.ReadFile("input")
    if err != nil {
        log.Fatal(err)
    }
    input := inputFile[:len(inputFile)-1]

	a := bytes.Split(input, []byte("\n\n"))

	movements := a[1]

	warehouse1 := bytes.Split(a[0], []byte("\n"))
	warehouse2 := [][]byte{}

	for _, row := range bytes.Split(a[0], []byte("\n")) {
		tmp := []byte{}
		for _, col := range row {
			switch col {
			case 'O': tmp = append(tmp, '[', ']');
			case '#': tmp = append(tmp, '#', '#')
			case '@': tmp = append(tmp, '@', '.')
			case '.': tmp = append(tmp, '.', '.')
			}
		}
		warehouse2 = append(warehouse2, tmp)
	}

    fmt.Printf("Part one: %d\n", partOne(warehouse1, movements))
    fmt.Printf("Part two: %d\n", partTwo(warehouse2, movements))
}

func partOne(warehouse [][]byte, movements []byte) int{
    sum := 0

	robot := struct {x, y int}{}
	BREAK:
	for i, row := range warehouse {
		for j, col := range row {
			if col == '@' {
				robot.y = i
				robot.x = j
				break BREAK
			}
		}
	}

	OUTER:
	for _, c := range movements {
		switch c {
		case '\n': continue
		case '<': {
			i := robot.x - 1
			for ; i >= 0; i-- {
				if warehouse[robot.y][i] == '#' { continue OUTER }
				if warehouse[robot.y][i] == '.' { break }
			}
			for ; i < robot.x; i++ {
				warehouse[robot.y][i] = warehouse[robot.y][i + 1]
			}
			warehouse[robot.y][robot.x] = '.'
			robot.x--
		}

		case '>': {
			i := robot.x + 1
			for ; i < len(warehouse[0]); i++ {
				if warehouse[robot.y][i] == '#' { continue OUTER }
				if warehouse[robot.y][i] == '.' { break }
			}
			for ; i > robot.x; i-- {
				warehouse[robot.y][i] = warehouse[robot.y][i - 1]
			}
			warehouse[robot.y][robot.x] = '.'
			robot.x++
		}

		case 'v': {
			i := robot.y + 1
			for ; i < len(warehouse); i++ {
				if warehouse[i][robot.x] == '#' { continue OUTER }
				if warehouse[i][robot.x] == '.' { break }
			}
			for ; i > robot.y; i-- {
				warehouse[i][robot.x] = warehouse[i - 1][robot.x]
			}
			warehouse[robot.y][robot.x] = '.'
			robot.y++
		}

		case '^': {
			i := robot.y - 1
			for ; i >= 0; i-- {
				if warehouse[i][robot.x] == '#' { continue OUTER }
				if warehouse[i][robot.x] == '.' { break }
			}
			for ; i < robot.y; i++ {
				warehouse[i][robot.x] = warehouse[i + 1][robot.x]
			}
			warehouse[robot.y][robot.x] = '.'
			robot.y--
		}
		}
	}

	for i, row := range warehouse {
		for j, col := range row {
			if col == 'O' {
				sum += (100 * i) + j
			}
		}
	}
    return sum
}

func partTwo(warehouse [][]byte, movements []byte) int{
    sum := 0

	robot := struct {x, y int}{}
	BREAK:
	for i, row := range warehouse {
		for j, col := range row {
			if col == '@' {
				robot.y = i
				robot.x = j
				break BREAK
			}
		}
	}

	OUTER:
	for _, c := range movements {

		switch c {
		case '\n': continue
		case '<': {
			i := robot.x - 1
			for ; i >= 0; i-- {
				if warehouse[robot.y][i] == '#' { continue OUTER }
				if warehouse[robot.y][i] == '.' { break }
			}
			for ; i < robot.x; i++ {
				warehouse[robot.y][i] = warehouse[robot.y][i + 1]
			}
			warehouse[robot.y][robot.x] = '.'
			robot.x--
		}

		case '>': {
			i := robot.x + 1
			for ; i < len(warehouse[0]); i++ {
				if warehouse[robot.y][i] == '#' { continue OUTER }
				if warehouse[robot.y][i] == '.' { break }
			}
			for ; i > robot.x; i-- {
				warehouse[robot.y][i] = warehouse[robot.y][i - 1]
			}
			warehouse[robot.y][robot.x] = '.'
			robot.x++
		}

		case 'v': {
			toMove := []pos{}
			allowed := false
			switch warehouse[robot.y + 1][robot.x] {
			case '#': continue OUTER
			case '[': {
				toMove, allowed = pushDown(robot.x, robot.y + 1, warehouse)
			}
			case ']': {
				toMove, allowed = pushDown(robot.x - 1, robot.y + 1, warehouse)
			}
			case '.': {
				warehouse[robot.y][robot.x] = '.'
				warehouse[robot.y + 1][robot.x] = '@'
				robot.y++
			}
			}

			if allowed {
				for _, box := range toMove {
					warehouse[box.y][box.x] = '.'
					warehouse[box.y][box.x + 1] = '.'
				}
				for _, box := range toMove {
					warehouse[box.y + 1][box.x] = '['
					warehouse[box.y + 1][box.x + 1] = ']'
				}
				warehouse[robot.y + 1][robot.x] = '@'
				warehouse[robot.y][robot.x] = '.'
				robot.y++
			}
		}

		case '^': {
			toMove := []pos{}
			allowed := false
			switch warehouse[robot.y - 1][robot.x] {
			case '[': {
				toMove, allowed = pushUp(robot.x, robot.y - 1, warehouse)
			}
			case ']': {
				toMove, allowed = pushUp(robot.x - 1, robot.y - 1, warehouse)
			}
			case '#': continue OUTER
			case '.': {
				warehouse[robot.y][robot.x] = '.'
				warehouse[robot.y - 1][robot.x] = '@'
				robot.y--
			}
			}
			if allowed {
				for _, box := range toMove {
					warehouse[box.y][box.x] = '.'
					warehouse[box.y][box.x + 1] = '.'
				}
				for _, box := range toMove {
					warehouse[box.y - 1][box.x] = '['
					warehouse[box.y - 1][box.x + 1] = ']'
				}
				warehouse[robot.y - 1][robot.x] = '@'
				warehouse[robot.y][robot.x] = '.'
				robot.y--
			}
		}
		}
	}

	for i, row := range warehouse {
		for j, col := range row {
			if col == '[' {
				sum += (100 * i) + j
			}
		}
	}
    return sum
}

func pushUp(x, y int, warehouse[][]byte) (toMove []pos, allowed bool){

	if warehouse[y - 1][x] == '#' || warehouse[y - 1][x + 1] == '#' {
		return nil, false
	} else {

		if warehouse[y - 1][x] == '[' {
			boxes, a := pushUp(x, y - 1, warehouse)
			if a {
				toMove = append(toMove, boxes...)
			} else {
				return nil, false
			}
		} else if warehouse[y - 1][x] == ']' {
			boxes, a := pushUp(x - 1, y - 1, warehouse) 
			if a {
				toMove = append(toMove, boxes...)
			} else {
				return nil, false
			}
		}

		if warehouse[y - 1][x + 1] == '[' {
			boxes, a := pushUp(x + 1, y - 1, warehouse) 
			if a {
				toMove = append(toMove, boxes...)
			} else {
				return nil, false
			}
		}
	}


	toMove = append(toMove, pos{x, y})
	return toMove, true

}


func pushDown(x, y int, warehouse[][]byte) (toMove []pos, allowed bool){

	if warehouse[y + 1][x] == '#' || warehouse[y + 1][x + 1] == '#' {
		return nil, false
	} else {
		if warehouse[y + 1][x] == '[' {
			boxes, a := pushDown(x, y + 1, warehouse)
			if a {
				toMove = append(toMove, boxes...)
			} else {
				return nil, false
			}
		} else if warehouse[y + 1][x] == ']' {
			boxes, a := pushDown(x - 1, y + 1, warehouse) 
			if a {
				toMove = append(toMove, boxes...)
			} else {
				return nil, false
			}
		}
		if warehouse[y + 1][x + 1] == '[' {
			boxes, a := pushDown(x + 1, y + 1, warehouse) 
			if a {
				toMove = append(toMove, boxes...)
			} else {
				return nil, false
			}
		}
	}

	toMove = append(toMove, pos{x, y})
	return toMove, true
}
