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
    
    input, err := os.ReadFile("input")
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Part one: %d\n", partOne(input))
    fmt.Printf("Part two: %d\n", partTwo(input))
}

func partOne(input []byte) int {

    for i := 0; i < math.MaxInt32; i++{
        key := fmt.Sprintf("%s%d", input, i)
        hashed := fmt.Sprintf("%x", md5.Sum([]byte(key)))

        if strings.HasPrefix(hashed, "00000") {
            return i
        }
    }
    return -1
}

func partTwo(input []byte) int {
    for i := 0; i < math.MaxInt32; i++{
        key := fmt.Sprintf("%s%d", input, i)
        hashed := fmt.Sprintf("%x", md5.Sum([]byte(key)))

        if strings.HasPrefix(hashed, "000000") {
            return i
        }
    }
    return -1
}
