package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var wires = map[string] []string {}
var liveWires = map[string] int {}

func main() {

    inputFile, err := os.ReadFile("input")
    if err != nil {
        log.Fatal(err)
    }
    input := string(inputFile[:len(inputFile)-1])

    lines := strings.Split(input, "\n")
    for _, line := range(lines) {
        words := strings.Split(line, " ")
        wires[words[len(words)-1]] = words[:len(words)-2]
    }

    fmt.Printf("Part one: %d\n", partOne())
    fmt.Printf("Part two: %d\n", partTwo())
}

func partOne() int{
    return wire("a")
}

func partTwo() int{
    value := wire("a")

    liveWires = make(map[string]int)
    liveWires["b"] = value
    return wire("a")
}

func isNumber(c byte) bool {
    return '0' <= c && c <= '9'
}

func wire(s string) int {

    if val, ok := liveWires[s]; ok {
        return val
    }

    rule := wires[s]

    if len(rule) == 1 {
        liveWires[s] = wireOrNumber(rule[0])
    } else if rule[0] == "NOT" {
        liveWires[s] = not(rule[1])
    } else {

        switch rule[1] {
        case "AND":
            liveWires[s] = and(rule[0], rule[2])
        case "OR":
            liveWires[s] = or(rule[0], rule[2])
        case "RSHIFT":
            liveWires[s] = rshift(rule[0], rule[2])
        case "LSHIFT":
            liveWires[s] = lshift(rule[0], rule[2])
        }
    }
    val, _ := liveWires[s]
    return val
}

func number(s string) int {
    n, err := strconv.Atoi(s)
    if err != nil {
        log.Fatal(err)
    }
    return n
}

func wireOrNumber(s string) int {
    if isNumber(s[0]) {
        return number(s)
    } else {
        return wire(s)
    }
}

func and(l, r string) int {
    lvalue := wireOrNumber(l)
    rvalue := wireOrNumber(r)
    return lvalue & rvalue
}

func or(l, r string) int {
    lvalue := wireOrNumber(l)
    rvalue := wireOrNumber(r)
    return lvalue | rvalue
}

func not(s string) int {
    value := wireOrNumber(s)
    return ^value
}

func lshift(l, r string) int {
    lvalue := wireOrNumber(l)
    rvalue := wireOrNumber(r)
    return lvalue << rvalue
}

func rshift(l, r string) int {
    lvalue := wireOrNumber(l)
    rvalue := wireOrNumber(r)
    return lvalue >> rvalue
}
