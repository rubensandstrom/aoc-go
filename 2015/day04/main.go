package main

import (
	"crypto/md5"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

func main() {
    
    inputFile, err := os.ReadFile("input")
    if err != nil {
        log.Fatal(err)
    }
    input := string(inputFile)
    input = strings.TrimSpace(input[:len(inputFile)-1])

    fmt.Printf("Part one: %d\n", partOne(input))
    fmt.Printf("Part two: %d\n", partTwo(input))
}

func partOne(input string) int {

    for i := 0; i < math.MaxInt32; i++{
        key := fmt.Sprintf("%s%d", input, i)
        hashed := fmt.Sprintf("%x", md5.Sum([]byte(key)))

        if strings.HasPrefix(hashed[0:5], "00000") {
            return i
        }
    }
    return -1
}

func partTwo(input string) int {
    for i := 0; i < math.MaxInt32; i++{
        key := fmt.Sprintf("%s%d", input, i)
        hashed := fmt.Sprintf("%x", md5.Sum([]byte(key)))

        if strings.HasPrefix(hashed[0:6], "000000") {
            return i
        }
    }
    return -1
}
