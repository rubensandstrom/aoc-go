package main

import (
	"fmt"
	"log"
	"os"
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

func matchDesign(design string, patterns map[string]bool, minLen, maxLen int, memo map[string]int) int {
	if len(design) == 0 { return 1 }
	if val, ok := memo[design]; ok && val > 0 {
		return memo[design]
	} else if ok && val == 0 {
		return 0
	}
	for i := minLen; i <= min(maxLen, len(design)); i++ {
		if patterns[design[:i]] {
			memo[design] += matchDesign(design[i:], patterns, minLen, maxLen, memo)
		}
	}
	return memo[design]
}

func partOne(input string) int{
    sum := 0
	tmp := strings.Split(input, "\n\n")
	minLen := 1000
	maxLen := 0
	patterns := map[string]bool{}
	for _, pattern := range strings.Split(tmp[0], ", ") {
		if len(pattern) < minLen {
			minLen = len(pattern)
		}
		if len(pattern) > maxLen {
			maxLen = len(pattern)
		}
		patterns[pattern] = true
	}
	designs := strings.Split(tmp[1], "\n")
	for _, design := range designs {
		if val := matchDesign(design, patterns, minLen, maxLen, map[string]int{}); val > 0 {
			sum++
		}
	} 
    return sum
}

func partTwo(input string) int{
    sum := 0
	tmp := strings.Split(input, "\n\n")
	minLen := 1000
	maxLen := 0
	patterns := map[string]bool{}
	for _, pattern := range strings.Split(tmp[0], ", ") {
		if len(pattern) < minLen {
			minLen = len(pattern)
		}
		if len(pattern) > maxLen {
			maxLen = len(pattern)
		}
		patterns[pattern] = true
	}
	designs := strings.Split(tmp[1], "\n")
	for _, design := range designs {
		sum += matchDesign(design, patterns, minLen, maxLen, map[string]int{})
	}
    return sum
}
