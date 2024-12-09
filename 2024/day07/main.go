package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	PLUS = iota
	TIMES
	CONCAT
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
    sum := 0

	for _, equation := range strings.Split(input, "\n") {
		words := strings.Split(equation, " ")
		result,err := strconv.Atoi(words[0][:len(words[0]) - 1])
		if err != nil {
			log.Fatal(err)
		}
		tmpSum, err := strconv.Atoi(words[1])
		if eval(result, tmpSum, words[2:], PLUS) || eval(result, tmpSum, words[2:], TIMES) {
			sum += result
		}
	}
    return sum
}

func partTwo(input string) int {
    sum := 0
	for _, equation := range strings.Split(input, "\n") {
		words := strings.Split(equation, " ")
		result,err := strconv.Atoi(words[0][:len(words[0]) - 1])
		if err != nil {
			log.Fatal(err)
		}
		tmpSum, err := strconv.Atoi(words[1])
		if eval2(result, tmpSum, words[2:], PLUS) || eval2(result, tmpSum, words[2:], TIMES) || eval2(result, tmpSum, words[2:], CONCAT) {
			sum += result
		}
	}
    return sum
}

func eval(result, sum int, operands []string, operator int) bool {

	if sum > result {
		return false
	}

	if len(operands) == 0{
		if sum == result {
			return true
		}
		return false
	}

	val, _ := strconv.Atoi(operands[0])
	switch operator {
	case PLUS: { sum += val }
	case TIMES: {sum *= val }
	}

	if eval(result, sum, operands[1:], PLUS) || eval(result, sum, operands[1:], TIMES) {
		return true
	}
	return false
}

func eval2(result, sum int, operands []string, operator int) bool {
	
	if sum > result {
		return false
	}

	if len(operands) == 0{
		if sum == result {
			return true
		}
		return false
	}

	val, _ := strconv.Atoi(operands[0])
	switch operator {
		case PLUS: { sum += val }
		case TIMES: {sum *= val }
		case CONCAT: {
			sumStr := strconv.Itoa(sum)
			sumStr += operands[0]
			sum, _ = strconv.Atoi(sumStr)
		}
	}

	if eval2(result, sum, operands[1:], PLUS) || eval2(result, sum, operands[1:], TIMES) || eval2(result, sum, operands[1:], CONCAT) {
		return true
	}
	return false
}
