package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
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

func partOne(input string) int {

    sum := 0

    lines := strings.Split(input, "\n")
    for _, line := range(lines) {
        box := strings.Split(line, "x")
        if len(box) != 3 {break}

        l, _ := strconv.Atoi(box[0])
        w, _ := strconv.Atoi(box[1])
        h, _ := strconv.Atoi(box[2])

        lw := l*w
        wh := w*h
        hl := h*l

        sum += 2*lw + 2*wh + 2*hl + min(lw, wh, hl)
    }
    return sum

}

func partTwo(input string) int {

    sum := 0

    lines := strings.Split(input, "\n")
    for _, line := range(lines) {
        boxString := strings.Split(line, "x")
        if len(boxString) != 3 {break}

        l, _ := strconv.Atoi(boxString[0])
        w, _ := strconv.Atoi(boxString[1])
        h, _ := strconv.Atoi(boxString[2])

        box := []int {l, w, h}

        sort.SliceStable(box, func(i, j int) bool {
            return box[i] < box[j]
        })

        sum += 2*box[0] + 2*box[1] + box[0]*box[1]*box[2]
    }
    return sum
}
