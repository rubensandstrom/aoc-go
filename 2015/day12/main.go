package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)
func main() {

    input, err := os.ReadFile("input")
    if err != nil {
        log.Fatal(err)
    }

    // input := []byte(`{{"red"{"red":5}[1,"red",5]}}`)
    fmt.Printf("Part one: %d\n", partOne(input))
    fmt.Printf("Part two: %d\n", partTwo(input))
}

func partOne(s []byte) int {

    sum := 0
    for i := 0; i < len(s); i++ {
        if isDigit(s[i]) {
            start := i
            for isDigit(s[i]) {i++}
            i, err := strconv.Atoi(string(s[start:i]))
            if err != nil {
                log.Fatal(err)
            }
            sum += i
        }
    }
    return sum
}

func partTwo(s []byte) int {

    sum := 0
    for i := 0; i < len(s); i++ {
        if s[i] == '{' {
            i = skipIfRed(i, s)
        }
        if isDigit(s[i]) {
            start := i
            for isDigit(s[i]) {i++}
            i, err := strconv.Atoi(string(s[start:i]))
            if err != nil {
                log.Fatal(err)
            }
            sum += i
        }
    }
    return sum
}

func isDigit(b byte) bool {
    return b == '-' || '0' <= b && b <= '9'
}

func skipIfRed(start int, s []byte) int {
    objectDepth := 1
    arrayDepth := 0
    red := false
    offset := start + 1
    for ; objectDepth > 0; offset++ {
        switch s[offset] {
        case '{':
            objectDepth++
        case '}':
            objectDepth--
        case '[':
            arrayDepth++
        case ']':
            arrayDepth--
        }
        if offset < start + 4 {continue}
        if objectDepth != 1 || arrayDepth != 0 {continue}

        if string(s[offset-4:offset+1]) == "\"red\"" {
            red = true
        }
    }
    if red == true {
        return offset
    }
    return start
}

