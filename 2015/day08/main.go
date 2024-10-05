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

func partOne(input string) int{
    stringLiteral := 0
    inMemoryString := 0

    lines := strings.Split(input, "\n")
    for _, line := range(lines) {
        stringLiteral += len(line)
        inMemoryString += calcInMemoryString(line)
    }
    return stringLiteral - inMemoryString
}

func partTwo(input string) int{
    stringLiteral := 0
    encodedString := 0

    lines := strings.Split(input, "\n")
    for _, line := range(lines) {
        stringLiteral += len(line)
        encodedString += calcEncodedString(line)
    }
    return encodedString - stringLiteral
}

func calcInMemoryString(s string) int {
    sum := 0
    for i := 1; i < len(s) - 1; i++ {
        if s[i] == '\\' {
            switch s[i + 1] {
            case '\\': i += 1
            case '"': i += 1
            case 'x': i += 3
            }
        }
        sum++
    }
    return sum
}

func calcEncodedString(s string) int {
    sum := 2
    for _, c := range(s) {
        if c == '"' || c == '\\' {sum++}
        sum++
    }
    return sum
}
