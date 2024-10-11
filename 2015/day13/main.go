package main

import (
	// "fmt"
	// "log"
	"day13/graph"
	"fmt"
	// "os"
	// "strconv"
	// "strings"
)

func main() {
	// inputFile, err := os.ReadFile("input")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// input := string(inputFile[:len(inputFile)-1])
	//
	// table := graph.Graph {}
	// lines := strings.Split(input, "\n")
	//
	// for _, line := range lines {
	// 	words := strings.Split(line[:len(line)-1], " ")
	// 	weight, _ := strconv.Atoi(words[3])
	// 	if words[2] == "lose" {
	// 		weight = -weight
	//
	// 	}
	// 	table.AddEdge(words[0], words[10], weight)
	// }
	//
	// fmt.Printf("Part one: %d\n", partOne(table))
	// fmt.Printf("Part two: %d\n", partTwo(input))

	g := graph.Graph {}
	g.AddEdge("a", "b", 10)
	g.AddEdge("b", "a", 20)
	g.AddEdge("a", "c", 20)
	g.AddEdge("c", "b", 20)
	g.AddEdge("c", "a", 20)
	g.AddEdge("b", "c", 20)

	if cost, err := g.TSP("a"); err != nil {
		panic(err)
	} else {
		fmt.Printf("%d\n", cost)
	}
}

func partOne(table graph.Graph) int{
	sum := 0
	table.Print()
	return sum
}

func partTwo(input string) int{
	sum := 0
	return sum
}
