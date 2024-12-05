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

	safetyManual := strings.Split(input, "\n\n")
	rules := safetyManual[0]

	rulesMap := map[string][]string{}
	for _, rule := range strings.Split(rules, "\n") {
		parts := strings.Split(rule, "|")
		rulesMap[parts[0]] = append(rulesMap[parts[0]], parts[1])
	}

	updates := safetyManual[1]
	updatesList := [][]string{}
	for _, update := range strings.Split(updates, "\n") {
		updatesList = append(updatesList, strings.Split(update, ","))
	}

    fmt.Printf("Part one: %d\n", partOne(rulesMap, updatesList))
    fmt.Printf("Part two: %d\n", partTwo(rulesMap, updatesList))
}

func partOne(rulesMap map[string][]string, updatesList [][]string) int{
    sum := 0

	for _, update := range updatesList {
		val, _ := value(rulesMap, update)
		sum += val
	}
    return sum
}
func partTwo(rulesMap map[string][]string, updatesList [][]string) int{
    sum := 0

	for _, update := range updatesList {
		_, valid := value(rulesMap, update)
		if valid { continue }
		sort.Slice(update, func(i, j int) bool {
			for _, k := range rulesMap[update[i]] {
				if update[j] == k {
					return false
				}
			}
			return true
		})
		val, _ := strconv.Atoi(update[(len(update) -1)/2])

		sum += val
	}
    return sum
}

func value(rulesMap map[string][]string, update []string) (int, bool) {
	for i := 0; i < len(update); i++ {
		for j := 0; j <= i; j++ {
			for _, k := range rulesMap[update[i]] {
				if update[j] == k {
					return 0, false
				}
			}
		}
	}
	val, _ := strconv.Atoi(update[(len(update) -1)/2])
	return val, true
}
