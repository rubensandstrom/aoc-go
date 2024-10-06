package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Graph map[string] Edge
type Edge map[string] int

func (g Graph) AddEdge(from, to string, weight int) {
    if _, ok := g[from]; !ok {
        g[from] = Edge{}
    }
    g[from][to] = weight
}

var cities = Graph{}

func main() {

    inputFile, err := os.ReadFile("input")
    if err != nil {
        log.Fatal(err)
    }
    input := string(inputFile[:len(inputFile)-1])
    lines := strings.Split(input, "\n")
    for _, line := range(lines) {

        words := strings.Split(line, " ")
        weight, err := strconv.Atoi(words[4])
        if err != nil {
            log.Fatal(err)
        }
        cities.AddEdge(words[0], words[2], weight)
        cities.AddEdge(words[2], words[0], weight)
    }
    fmt.Printf("Part one: %d\n", partOne())
    fmt.Printf("Part two: %d\n", partTwo())

}

func partOne() int{

    notVisited := make([]string, 0, len(cities))
    for k := range cities {
        notVisited = append(notVisited, k)
    }

    return minTravel(notVisited[0], notVisited[1:], 0)
}

func partTwo() int{

    notVisited := make([]string, 0, len(cities))
    for k := range cities {
        notVisited = append(notVisited, k)
    }

    return maxTravel(notVisited[0], notVisited[1:], 0)
}

// BUG: Undeterministic.
func minTravel(start string, toVisit []string, currentCost int) int {

    if len(toVisit) == 0 {
        return currentCost
    }

    costs := []int {}
    for i, next := range(toVisit) {
        nextToVisit := slices.Concat(toVisit[:i], toVisit[i+1:])
        cost := minTravel(next, nextToVisit, currentCost + cities[start][next])
        costs = append(costs, cost)
    }

    return slices.Min(costs)
}

// BUG: Undeterministic.
func maxTravel(start string, toVisit []string, currentCost int) int {

    if len(toVisit) == 0 {
        return currentCost
    }

    costs := []int {}
    for i, next := range(toVisit) {
        nextToVisit := slices.Concat(toVisit[:i], toVisit[i+1:])
        cost := maxTravel(next, nextToVisit, currentCost + cities[start][next])
        costs = append(costs, cost)
    }

    return slices.Max(costs)
}
