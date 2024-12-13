package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)


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
	r, _ := regexp.Compile(`\d+`)
	sum := 0

	
	for _, machine := range strings.Split(input, "\n\n") {
		a := struct     {x, y int}{x: 0, y: 0}
		b := struct     {x, y int}{x: 0, y: 0}
		prize := struct {x, y int}{x: 0, y: 0}

		lines := strings.Split(machine, "\n")

		as := r.FindAllString(lines[0], -1)
		a.x, _ = strconv.Atoi(as[0])
		a.y, _ = strconv.Atoi(as[1])

		bs := r.FindAllString(lines[1], -1)
		b.x, _ = strconv.Atoi(bs[0])
		b.y, _ = strconv.Atoi(bs[1])

		ps := r.FindAllString(lines[2], -1)
		prize.x, _ = strconv.Atoi(ps[0])
		prize.y, _ = strconv.Atoi(ps[1])

		tmpX := a.x
		tmpY := a.y

		a.x *= tmpY
		b.x *= tmpY
		prize.x *= tmpY

		a.y *= tmpX
		b.y *= tmpX
		prize.y *= tmpX
		
		a.y -= a.x
		b.y -= b.x
		prize.y -= prize.x

		if prize.y % b.y != 0 { continue }
		prize.y /= b.y
		b.y = 1

		prize.x -= (b.x * prize.y)
		b.x = 0
		
		if prize.x % a.x != 0 { continue }
		prize.x /= a.x

		sum += prize.x * 3 + prize.y
		// fmt.Printf( "%d %d\n", A, B)
	}
	return sum
}

func partTwo(input string) int{
	r, _ := regexp.Compile(`\d+`)
	sum := 0

	
	for _, machine := range strings.Split(input, "\n\n") {
		a := struct     {x, y int}{x: 0, y: 0}
		b := struct     {x, y int}{x: 0, y: 0}
		prize := struct {x, y int}{x: 0, y: 0}

		lines := strings.Split(machine, "\n")

		as := r.FindAllString(lines[0], -1)
		a.x, _ = strconv.Atoi(as[0])
		a.y, _ = strconv.Atoi(as[1])

		bs := r.FindAllString(lines[1], -1)
		b.x, _ = strconv.Atoi(bs[0])
		b.y, _ = strconv.Atoi(bs[1])

		ps := r.FindAllString(lines[2], -1)
		prize.x, _ = strconv.Atoi(ps[0])
		prize.x += 10000000000000
		prize.y, _ = strconv.Atoi(ps[1])
		prize.y += 10000000000000

		tmpX := a.x
		tmpY := a.y

		a.x *= tmpY
		b.x *= tmpY
		prize.x *= tmpY

		a.y *= tmpX
		b.y *= tmpX
		prize.y *= tmpX
		
		a.y -= a.x
		b.y -= b.x
		prize.y -= prize.x

		if prize.y % b.y != 0 { continue }
		prize.y /= b.y
		b.y = 1

		prize.x -= (b.x * prize.y)
		b.x = 0
		
		if prize.x % a.x != 0 { continue }
		prize.x /= a.x

		sum += prize.x * 3 + prize.y
		// fmt.Printf( "%d %d\n", A, B)
	}
	return sum
}
