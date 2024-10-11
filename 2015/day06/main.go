package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// TODO: There is probably some smart math way to count with ranges without making grids to keep trak of these.
var grid [1000][1000]bool
var gradualGrid [1000][1000]int

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
    lines := strings.Split(input, "\n")
    for _, line := range(lines) {
        words := strings.Split(line, " ")
        if words[0] == "turn" {

            rc0 := strings.Split(words[2], ",")
            rc1 := strings.Split(words[4], ",")

            rowStart, _ := strconv.Atoi(rc0[0])
            colStart, _ := strconv.Atoi(rc0[1])
            rowStop, _ := strconv.Atoi(rc1[0])
            colStop, _ := strconv.Atoi(rc1[1])
            if words[1] == "on" {
                turnOn(rowStart, colStart, rowStop, colStop)
            } else if words[1] == "off" {
                turnOff(rowStart, colStart, rowStop, colStop)
            }

        } else if words[0] == "toggle" {
            rc0 := strings.Split(words[1], ",")
            rc1 := strings.Split(words[3], ",")

            rowStart, _ := strconv.Atoi(rc0[0])
            colStart, _ := strconv.Atoi(rc0[1])
            rowStop, _ := strconv.Atoi(rc1[0])
            colStop, _ := strconv.Atoi(rc1[1])

            toggle(rowStart, colStart, rowStop, colStop)
        }
    }

    for row := 0; row < len(grid); row++ {
        for col := 0; col < len(grid[0]); col++ {

            if grid[row][col] {
                sum++
            }
        }
    }


    return sum
}

func partTwo(input string) int{
    sum := 0
    lines := strings.Split(input, "\n")
    for _, line := range(lines) {
        words := strings.Split(line, " ")
        if words[0] == "turn" {

            rc0 := strings.Split(words[2], ",")
            rc1 := strings.Split(words[4], ",")

            rowStart, _ := strconv.Atoi(rc0[0])
            colStart, _ := strconv.Atoi(rc0[1])
            rowStop, _ := strconv.Atoi(rc1[0])
            colStop, _ := strconv.Atoi(rc1[1])
            if words[1] == "on" {
                turnUp(rowStart, colStart, rowStop, colStop)
            } else if words[1] == "off" {
                turnDown(rowStart, colStart, rowStop, colStop)
            }

        } else if words[0] == "toggle" {
            start := strings.Split(words[1], ",")
            stop := strings.Split(words[3], ",")

            rowStart, _ := strconv.Atoi(start[0])
            colStart, _ := strconv.Atoi(start[1])
            rowStop, _ := strconv.Atoi(stop[0])
            colStop, _ := strconv.Atoi(stop[1])

            turnUpTwo(rowStart, colStart, rowStop, colStop)
        }
    }

    for row := 0; row < len(grid); row++ {
        for col := 0; col < len(grid[0]); col++ {
            sum += gradualGrid[row][col]
        }
    }


    return sum
}

func turnOn(colStart, rowStart, colStop, rowStop int) {
    for row := rowStart; row <= rowStop; row++ {
        for col := colStart; col <= colStop; col++ {
            grid[row][col] = true
        }
    }
}

func turnOff(colStart, rowStart, colStop, rowStop int) {
    for row := rowStart; row <= rowStop; row++ {
        for col := colStart; col <= colStop; col++ {
            grid[row][col] = false
        }
    }
}

func toggle(colStart, rowStart, colStop, rowStop int) {
    for row := rowStart; row <= rowStop; row++ {
        for col := colStart; col <= colStop; col++ {
            grid[row][col] = !grid[row][col]
        }
    }
}

func turnUp(colStart, rowStart, colStop, rowStop int) {
    for row := rowStart; row <= rowStop; row++ {
        for col := colStart; col <= colStop; col++ {
            gradualGrid[row][col]++
        }
    }
}

func turnUpTwo(colStart, rowStart, colStop, rowStop int) {
    for row := rowStart; row <= rowStop; row++ {
        for col := colStart; col <= colStop; col++ {
            gradualGrid[row][col] += 2
        }
    }
}

func turnDown(colStart, rowStart, colStop, rowStop int) {
    for row := rowStart; row <= rowStop; row++ {
        for col := colStart; col <= colStop; col++ {
            if gradualGrid[row][col] == 0 {continue}
            gradualGrid[row][col]--
        }
    }
}
