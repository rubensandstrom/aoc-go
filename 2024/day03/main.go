package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"regexp"
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
    sum := 0
	r1, _ := regexp.Compile(`mul\(\d+,\d+\)`)
	r2, _ := regexp.Compile(`\d+`)
	for _, op := range r1.FindAllString(input, -1) {
		factors := r2.FindAllString(op, -1)

		a, _ := strconv.Atoi(factors[0])
		b, _ := strconv.Atoi(factors[1])

		sum += a * b
	}

    return sum
}
func partTwo(input string) int{
    sum := 0
	r1, _ := regexp.Compile(`mul\(\d+,\d+\)|do\(\)|don't\(\)`)
	r2, _ := regexp.Compile(`\d+`)

	enable := true
	for _, op := range r1.FindAllString(input, -1) {
		if op == "don't()" {
			enable = false
			continue
		}
		if op == "do()" {
			enable = true
			continue
		}
		if enable {
			factors := r2.FindAllString(op, -1)

			a, _ := strconv.Atoi(factors[0])
			b, _ := strconv.Atoi(factors[1])

			sum += a * b
		}
	}
    return sum
}
