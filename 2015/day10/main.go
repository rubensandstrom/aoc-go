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

    fmt.Printf("Part one: %d\n", partOne(input))
    fmt.Printf("Part two: %d\n", partTwo(input))
}

func partOne(input string) int{
    s := input
    for i:= 0; i < 40; i++ {
        s = lookAndSay(s)
    }
    return len(s)
}

func partTwo(input string) int{
    s := input
    for i:= 0; i < 50; i++ {
        s = lookAndSay(s)
    }
    return len(s)
}

func lookAndSay(s string) string {

    b := strings.Builder {}

    count := 1
    i := 1
    for ; i < len(s); i++ {
        if s[i-1] == s[i] {
            count++
        } else {
            b.Write([]byte(strconv.Itoa(count)))
            b.WriteByte(s[i-1])
            count = 1
        }
    } 
    b.Write([]byte(strconv.Itoa(count)))
    b.WriteByte(s[i-1])
    return b.String()
}
