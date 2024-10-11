package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

    inputFile, err := os.ReadFile("input")
    if err != nil {
        log.Fatal(err)
    }
    input := inputFile[:len(inputFile)-1]

    fmt.Printf("Part one: %s\n", partOne(input))
    fmt.Printf("Part two: %s\n", partTwo(input))
}

func partOne(input []byte) []byte {
    s := input
    for !hasStraight(s) || hasConfusing(s) || countPairs(s) < 2 {
        s = incString(s)
    }
    return s
}
func partTwo(input []byte) []byte {
    s := input
    for !hasStraight(s) || hasConfusing(s) || countPairs(s) < 2 {
        s = incString(s)
    }
    s = incString(s)
    for !hasStraight(s) || hasConfusing(s) || countPairs(s) < 2 {
        s = incString(s)
    }
    return s
}

func incString(s []byte) []byte {
    for i := len(s) -1; i >= 0; i-- {
        if s[i] != 'z' {
            s[i]++
            break
        }
        s[i] = 'a'
    }
    return s
}

func hasStraight(s []byte) bool {
    for i := 2; i < len(s); i++ {
        if s[i-2]+1 == s[i-1] && s[i-1] + 1 == s[i] {
            return true
        }
    } 
    return false
}

func hasConfusing(s []byte) bool {
    for _, c := range(s) {
        if c == 'i' || c == 'o' || c == 'l' {
            return true
        }
    }
    return false
}

func countPairs(s []byte) int {
    sum := 0
    for i := 1; i < len(s); i++ {
        if s[i-1] == s[i] {
            sum++
            i++
        }
    }
    return sum
}
