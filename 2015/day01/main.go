package main

import (
	"fmt"
	"os"
)

func main() {
    input, err := os.ReadFile("input")
    if err != nil {
        fmt.Printf("Couldn't read file\n")
    }

    fmt.Printf("Part one: %d\n", partOne(input))
    fmt.Printf("Part two: %d\n", partTwo(input))
}

func partOne(input []byte) int {
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

func partTwo(input []byte) int {
    sum := 0
    i:= 0
    for ; sum >= 0 ; i++ {
        if c := input[i]; c == '(' {
            sum++
        } else {
            sum--
        }
    }

    return i
}
