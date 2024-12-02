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

	data := [][]int{}
	for _, row := range strings.Split(input, "\n") {

		report := []int{}
		for _, level := range strings.Split(row, " ") {
			tmp, _ := strconv.Atoi(level)
			report = append(report, tmp)
		}
		data = append(data, report)

	}
    fmt.Printf("Part one: %d\n", partOne(data))
    fmt.Printf("Part two: %d\n", partTwo(data))
}

func partOne(data [][]int) int{
    sum := 0
	for _, report := range data {
		increasing := report[0] < report[1]
		if safeReport(report, increasing) { sum ++ }
	}
    return sum
}
func partTwo(data [][]int) int{
    sum := 0
	for _, report := range data {

		increasing := report[0] < report[1]
		if safeReport(report, increasing) { 
			sum ++ 
			continue
		}

		for i := 0; i < len(report); i++ {

			filteredReport := []int{}
			filteredReport = append(filteredReport, report[:i]...)
			filteredReport = append(filteredReport, report[i+1:]...)
			

			increasing := filteredReport[0] < filteredReport[1]
			if safeReport(filteredReport, increasing) { 
				sum ++ 
				break
			}

		}
	}
    return sum
}

func safeReport(report []int, increment bool) bool {
	for i := 1; i < len(report); i++ {
		tmp := 0
		if increment {
			tmp = report[i] - report[i - 1]
		} else {
			tmp = report[i - 1] - report[i]
		}

		if tmp < 1 || tmp > 3 {
			return false
		}
	}
	return true
}
