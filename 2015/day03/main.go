package main

import (
	"fmt"
	"log"
	"os"
)

type Coord struct {
    x int
    y int
}

func main() {
    inputFile, err := os.ReadFile("input")
    if err != nil {
        log.Fatal(err)
    }
    input := string(inputFile)

    fmt.Printf("Part one: %d\n", partOne(input))
    fmt.Printf("Part two: %d\n", partTwo(input))

}

func partOne(input string) int {
    santa := Coord{}
    house := map[Coord] int {}
    house[santa]++
    for _, c := range(input) {
        switch c {
        case '>': santa.x++
        case '<': santa.x--
        case '^': santa.y++
        case 'v': santa.y--
        }
        house[santa]++
    }
    return len(house)
}

func partTwo(input string) int {
    santa := Coord{}
    roboSanta := Coord{}
    house := map[Coord] int {}
    house[santa]++
    for i, c := range(input) {
        if i % 2 == 0 {
            switch c {
            case '>': santa.x++
            case '<': santa.x--
            case '^': santa.y++
            case 'v': santa.y--
            }
            house[santa]++
        } else {
            switch c {
            case '>': roboSanta.x++
            case '<': roboSanta.x--
            case '^': roboSanta.y++
            case 'v': roboSanta.y--
            }
            house[roboSanta]++
        }
    }
    return len(house)
}
