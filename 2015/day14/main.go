package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

    inputFile, err := os.ReadFile("input")
    if err != nil {
		log.Fatal(err)
    }
    input := string(inputFile[:len(inputFile)-1])

    // fmt.Printf("Part one: %d\n", partOne(input))
    fmt.Printf("Part one: %d\n", partOne(input))
    fmt.Printf("Part two: %d\n", partTwo(input))
}

func partOne(input string) int {

	sum := [9]int{}
	for l, line := range strings.Split(input, "\n") {
		words := strings.Split(line, " ")

		speed, _ := strconv.Atoi(words[3])
		fly, _ := strconv.Atoi(words[6])
		rest, _ := strconv.Atoi(words[13])

		j := 0
		for i := 0; i < 2503; i++ {
			if j < fly {
				sum[l] += speed
				j++
				continue
			}
			if j < fly + rest {
				j++
				continue
			}
			j = 0
		}
	}
	return max(sum[0], sum[1], sum[2], sum[3], sum[4], sum[5], sum[6], sum[7], sum[8])
}

func partTwo(input string) int{
    sum := 0
    return sum
}

