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

	stones := map[int]int{}
	for _, stone := range strings.Split(input, " ") {
		s, _ := strconv.Atoi(stone)
		stones[s]++ 
	}

    fmt.Printf("Part one: %d\n", partOne(stones))
    fmt.Printf("Part two: %d\n", partTwo(stones))
}

func partOne(stones map[int]int) int{
	sum := 0
	for i := 0; i < 25; i++ {
		stones = blink(stones)
	}

	for _, count := range stones {
		sum += count
	}
	return sum
}

func partTwo(stones map[int]int) int{
	sum := 0
	for i := 0; i < 75; i++ {
		stones = blink(stones)
	}
	for _, count := range stones {
		sum += count
	}
	return sum
}


func blink(stones map[int]int) map[int]int {
	next := map[int]int{}
	for stone, count := range stones {

		if stone == 0 { 
			next[1] += count
			continue
		}
		s := strconv.Itoa(stone)
		if len(s) % 2 == 0 {
			left, _ := strconv.Atoi(s[:len(s)/2])
			right, _ := strconv.Atoi(s[len(s)/2:])

			next[left] += count
			next[right] += count
			continue
		}

		next[stone * 2024] += count
	}
	return next
}
