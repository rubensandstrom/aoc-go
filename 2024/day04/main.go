package main

import (
	"fmt"
	"log"
	"os"
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
    sum := 0
	rows := strings.Split(input, "\n")
	cols := []string{}

	for col := 0; col < len(rows[0]); col++ {
		tmp := strings.Builder{}
		for row := 0; row < len(rows); row++ {
			tmp.WriteByte(rows[row][col])
		}
		cols = append(cols, tmp.String())
	}

	majorDiags := []string{}
	minorDiags := []string{}
	for k := 0; k < len(rows) + len(cols) - 1; k++ {
		diagonal := strings.Builder{}
		for i := 0; i < len(rows); i++ {
			j := k - i
			if j >= 0 && j < len(cols) {
				diagonal.WriteByte(rows[i][j])
			}
		}
		if len(diagonal.String()) > 0 {
			majorDiags = append(majorDiags, diagonal.String())
		}
	}

	for k := 0; k < len(rows) + len(cols) - 1; k++ {
		diagonal := strings.Builder{}
		for i := 0; i < len(rows); i++ {
			j := k - (len(rows) - 1 - i)
			if j >= 0 && j < len(cols) {
				diagonal.WriteByte(rows[i][j])
			}
		}
		if len(diagonal.String()) > 0 {
			minorDiags = append(minorDiags, diagonal.String())
		}
	}

	for _, x := range rows {
		for y := 0; y <= len(x) - 4; y++ {
			if x[y:y+4] == "XMAS" {sum++}
			if x[y:y+4] == "SAMX" {sum++}
		}
	}

	for _, x := range cols {
		for y := 0; y <= len(x) - 4; y++ {
			if x[y:y+4] == "XMAS" {sum++}
			if x[y:y+4] == "SAMX" {sum++}
		}
	}
	
	for _, x := range minorDiags {
		if len(x) < 4 {continue}
		for y := 0; y <= len(x) - 4; y++ {
			if x[y:y+4] == "XMAS" {sum++}
			if x[y:y+4] == "SAMX" {sum++}
		}
	}

	for _, x := range majorDiags {
		if len(x) < 4 {continue}
		for y := 0; y <= len(x) - 4; y++ {
			if x[y:y+4] == "XMAS" {sum++}
			if x[y:y+4] == "SAMX" {sum++}
		}
	}

    return sum
}

func partTwo(input string) int{
    sum := 0

	rows := strings.Split(input, "\n")
	for row := 1; row < len(rows) - 1; row++ {
		for col := 1; col < len(rows[0]) - 1; col++ {
			if rows[row][col] == byte('A') {
				majorDiag := strings.Builder{}
				majorDiag.WriteByte(rows[row - 1][col - 1])
				majorDiag.WriteByte(rows[row][col])
				majorDiag.WriteByte(rows[row + 1][col + 1])

				minorDiag := strings.Builder{}
				minorDiag.WriteByte(rows[row - 1][col + 1])
				minorDiag.WriteByte(rows[row][col])
				minorDiag.WriteByte(rows[row + 1][col - 1])

				majorDiagStr := majorDiag.String()
				minorDiagStr := minorDiag.String()


				if (majorDiagStr == "MAS" || majorDiagStr == "SAM") && (minorDiagStr == "MAS" || minorDiagStr == "SAM") {sum++}

			}
		}
	}
    return sum
}
