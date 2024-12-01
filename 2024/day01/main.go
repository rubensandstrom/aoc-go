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
    input := string(inputFile[:len(inputFile)-1])

	var (
		list1 []int
		list2 []int
	)

	for _, rows := range strings.Split(input, "\n") {
		colums := strings.Split(rows, "   ")
		
		val1, _ := strconv.Atoi(colums[0])
		val2, _ := strconv.Atoi(colums[1])

		list1 = append(list1, val1)
		list2 = append(list2, val2)
	}

	sort.Slice(list1, func(i, j int) bool {
		return list1[i] < list1[j]
	})

	sort.Slice(list2, func(i, j int) bool {
		return list2[i] < list2[j]
	})

    fmt.Printf("Part one: %d\n", partOne(list1, list2))
    fmt.Printf("Part two: %d\n", partTwo(list1, list2))
}

func partOne(list1, list2 []int) int{
    sum := 0
	for i := range list1 {
		tmp := list2[i] - list1[i]
		if tmp >= 0 { 
			sum += tmp 
		} else { 
			sum -= tmp 
		}
	}
    return sum
}

func partTwo(list1, list2 []int) int{
    sum := 0
	for _, n := range list1 {

		tmp := 0
		for _, m := range list2 {
			if n == m {tmp += 1}
		}
		sum += n * tmp

	}
    return sum
}
