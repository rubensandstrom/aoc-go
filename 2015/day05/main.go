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
    input := string(inputFile)

    fmt.Printf("Part one: %d\n", partOne(input))
    fmt.Printf("Part two: %d\n", partTwo(input))
}

func partOne(s string) int {

    sum := 0
    lines := strings.Split(s, "\n")
    for _, line := range(lines) {
        if hasDouble(line) && !hasForbidden(line) && countVowels(line) >= 3 {
            sum++
        }
    }
    return sum

}
func partTwo(s string) int {

    sum := 0
    lines := strings.Split(s, "\n")
    for _, line := range(lines) {
        if hasPairTwice(line) && hasSandwich(line) {
            sum++
        }
    }
    return sum
}

func hasDouble(s string) bool {
    for i := 1; i < len(s); i++ {
        if s[i-1] == s[i] {
            return true
        }
    }
    return false
}

func hasForbidden(s string) bool {
    for i := 2; i <= len(s); i++ {
        t := s[i-2:i]
        if t == "ab" || t == "cd" || t == "pq" || t == "xy" {
            return true
        }
    }
    return false
}

func countVowels(s string) int {
    sum := 0
    vowels := "aeiou"
    for _, c := range(s) {
        for _, d := range(vowels) {
            if c == d { sum++ }
        }
    }
    return sum
}

func hasPairTwice(s string) bool {
    for i := 2; i <= len(s); i++ {
        for j := i+2; j <= len(s); j++ {
            if s[i-2:i] == s[j-2:j] {return true}
        }
    }
    return false
}

func hasSandwich(s string) bool {
    for i := 2; i < len(s); i++ {
        if s[i-2] == s[i] { return true}
    }
    return false
}
