package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func mod(a, b int) int {
	tmp := a % b
	if tmp < 0 {
		return tmp + b
	}
	return tmp
}

func main() {

    inputFile, err := os.ReadFile("input")
    if err != nil {
        log.Fatal(err)
    }
    input := string(inputFile[:len(inputFile)-1])

    fmt.Printf("Part one: %d\n", partOne(input))
    partTwo(input)
}

type robot struct {posX, posY, velX, velY int}

func partOne(input string) int{
	robots := []robot{}
	re, _ := regexp.Compile(`-?\d+`)
	for _, line := range strings.Split(input, "\n") {
		instr := re.FindAllString(line, -1)
		r := robot{}
		r.posX, _ = strconv.Atoi(instr[0])
		r.posY, _ = strconv.Atoi(instr[1])
		r.velX, _ = strconv.Atoi(instr[2])
		r.velY, _ = strconv.Atoi(instr[3])

		robots = append(robots, r)
	}

	for i := 0; i < 100; i++ {
		for j := range robots {
			robots[j].posX = mod(robots[j].posX + robots[j].velX, 101)
			robots[j].posY = mod(robots[j].posY + robots[j].velY, 103)
		}
	}

	var first, second, third, fourth int
	for _, robot := range robots {
		if robot.posX >=  51 && robot.posX <= 100 && robot.posY >=  0 && robot.posY <=  50 {first  ++}
		if robot.posX >=   0 && robot.posX <=  49 && robot.posY >=  0 && robot.posY <=  50 {second ++}
		if robot.posX >=   0 && robot.posX <=  49 && robot.posY >= 52 && robot.posY <= 102 {third ++}
		if robot.posX >=  51 && robot.posX <= 100 && robot.posY >= 52 && robot.posY <= 102 {fourth ++}
	}

    return first * second * third * fourth
}

func partTwo(input string) {
	robots := []robot{}
	re, _ := regexp.Compile(`-?\d+`)
	for _, line := range strings.Split(input, "\n") {
		instr := re.FindAllString(line, -1)
		r := robot{}
		r.posX, _ = strconv.Atoi(instr[0])
		r.posY, _ = strconv.Atoi(instr[1])
		r.velX, _ = strconv.Atoi(instr[2])
		r.velY, _ = strconv.Atoi(instr[3])

		robots = append(robots, r)
	}

	matrix := [103][101]byte{}

	for i := 1; i < 1000000; i++ {

		for row := 0; row < 103; row++ {
			for col := 0; col < 101; col++ {
				matrix[row][col] = '.'
			}
		}

		for j := range robots {
			robots[j].posX = mod(robots[j].posX + robots[j].velX, 101)
			robots[j].posY = mod(robots[j].posY + robots[j].velY, 103)
			matrix[robots[j].posY][robots[j].posX] = 'X'

		}

		score := 0
		for row := 1; row < 102; row++ {
			for col := 1; col < 100; col++ {

				sum := 0
				for a := -1; a <= 1; a++ {
					for b := -1; b <= 1; b++ {

						if matrix[row + a][col + b] == 'X' { sum++ }
					}
				}
				if sum == 9 { score++ }
			}
		}
 
		if  score > 3 {
			fmt.Printf("%d\n",i)
			for row := 0; row < 103; row++ {
				for col := 0; col < 101; col++ {
					fmt.Printf("%c", matrix[row][col])
				}
				fmt.Printf("\n")
			}
			fmt.Printf("\n")
		}
	}

}
