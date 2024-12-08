package main

import (
	"bytes"
	"fmt"
	"log"
	"os"	
	"aoc/util"
)

func main() {

    inputFile, err := os.ReadFile("input")
    if err != nil {
        log.Fatal(err)
    }
    input := inputFile[:len(inputFile)-1]

    fmt.Printf("Part one: %d\n", partOne(input))
    fmt.Printf("Part two: %d\n", partTwo(input))
}

func partOne(input []byte) int{
    sum := 0
	rows := bytes.Split(input, []byte("\n"))

	for _, x := range rows {
		for y := 0; y <= len(x) - 4; y++ {
			if string(x[y:y+4]) == "XMAS" {sum++}
			if string(x[y:y+4]) == "SAMX" {sum++}
		}
	}

	for _, x := range util.Cols(rows) {
		for y := 0; y <= len(x) - 4; y++ {
			if string(x[y:y+4]) == "XMAS" {sum++}
			if string(x[y:y+4]) == "SAMX" {sum++}
		}
	}
	
	for _, x := range util.MinorDiagonals(rows) {
		if len(x) < 4 {continue}
		for y := 0; y <= len(x) - 4; y++ {
			if string(x[y:y+4]) == "XMAS" {sum++}
			if string(x[y:y+4]) == "SAMX" {sum++}
		}
	}

	for _, x := range util.MajorDiagonals(rows) {
		if len(x) < 4 {continue}
		for y := 0; y <= len(x) - 4; y++ {
			if string(x[y:y+4]) == "XMAS" {sum++}
			if string(x[y:y+4]) == "SAMX" {sum++}
		}
	}

    return sum
}

func partTwo(input []byte) int{
    sum := 0

	rows := bytes.Split(input, []byte("\n"))
	for _, m := range util.Window(rows, 3, 3) {
		if m[1][1] == byte('A') {
			majorDiagStr := string([]byte{m[0][0], m[1][1], m[2][2]})
			minorDiagStr := string([]byte{m[0][2], m[1][1], m[2][0]})

			if (majorDiagStr == "MAS" || majorDiagStr == "SAM") && (minorDiagStr == "MAS" || minorDiagStr == "SAM") {sum++}
		}
	}
    return sum
}

/*
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
*/
