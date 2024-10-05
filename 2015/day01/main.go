package main

import (
    "os"
    "fmt"
)

func main() {
    inputFile, err := os.ReadFile("input")
    if err != nil {
        fmt.Printf("Couldn't read file\n")
    }
    input := string(inputFile[:len(inputFile)-1])

    fmt.Printf("Part one: %d\n", partOne(input))
    fmt.Printf("Part two: %d\n", partTwo(input))
}

func partOne(input string) int {
    sum := 0
    for i := 0; i < len(input) ; i++ {
        if c := input[i]; c == '(' {
            sum++
        } else if c == ')' {
            sum--
        }
    }
    return sum
}

func partTwo(input string) int {
    sum := 0
    i:= 0
    for ; sum >= 0 ; i++ {
        if c := input[i]; c == '(' {
            sum++
        } else if c == ')' {
            sum--
        }
    }

    return i
}
