package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

func main() {
    inputFile, err := os.ReadFile("input")
    if err != nil {
        log.Fatal(err)
    }
    input := string(inputFile[:len(inputFile)-1])

    fmt.Printf("Part one: %s\n", partOne(input))
    fmt.Printf("Part two: %d\n", partTwo(input))
}

func mod(a, b int) int {
	tmp := a % b
	if tmp < 0 { return -tmp }
	return tmp
}

func cmb(i int, A, B, C *int) int {
	if i <= 3 {
		return i
	}
	switch i {
	case 4: return *A
	case 5: return *B
	case 6: return *C
	}
	return -1
}

func execute (A, B, C int, prog []int) (out []int) {
	for op := 0; op < len(prog); op += 2 {
		switch prog[op] {
		case 0: {
			tmp := cmb(prog[op + 1], &A, &B, &C)
			A /= int(math.Pow(2, float64(tmp)))
		}
		case 1: {
			B = B ^ prog[op + 1]
		}
		case 2: {
			tmp := cmb(prog[op + 1], &A, &B, &C)
			B = mod(tmp, 8)
		}
		case 3: {
			if A == 0 {
				continue
			}
			op = prog[op + 1] - 2
		}
		case 4: {
			B = B ^ C
		}
		case 5: {
			tmp := cmb(prog[op + 1], &A, &B, &C)
			tmp = mod(tmp, 8)
			out = append(out, tmp)
		}
		case 6: {
			tmp := cmb(prog[op + 1], &A, &B, &C)
			B = A / int(math.Pow(2, float64(tmp)))
		}
		case 7: {
			tmp := cmb(prog[op + 1], &A, &B, &C)
			C = A / int(math.Pow(2, float64(tmp)))
		}
		}
	}
	return out
}

func partOne(input string) string {

	r, _ := regexp.Compile(`\d+`)

	tmp := strings.Split(input, "\n\n")
	regs := r.FindAllString(tmp[0], 3)

	prog := []int{}
	for _, op  := range strings.Split(tmp[1][9:], ","){
		tmp, _ := strconv.Atoi(op)
		prog = append(prog, tmp)
	}

	A, _ := strconv.Atoi(regs[0])
	B, _ := strconv.Atoi(regs[1])
	C, _ := strconv.Atoi(regs[2])

	out := execute(A, B, C, prog)
	outStr := strings.Builder{}

	for _, o := range out {
		tmp := strconv.Itoa(o)
		outStr.WriteString(tmp)
		outStr.WriteRune(',')
	}
	return outStr.String()[:outStr.Len() - 1]
}

func partTwo(input string) int{

	r, _ := regexp.Compile(`\d+`)

	tmp := strings.Split(input, "\n\n")
	regs := r.FindAllString(tmp[0], 3)

	prog := []int{}
	for _, op  := range strings.Split(tmp[1][9:], ","){
		tmp, _ := strconv.Atoi(op)
		prog = append(prog, tmp)
	}

	A := int(math.Pow(8, float64(len(prog) - 1)))
	B, _ := strconv.Atoi(regs[1])
	C, _ := strconv.Atoi(regs[2])

	OUT:
	for {
		out := execute(A, B, C, prog)

		if slices.Equal(out, prog) { break }
		for i := len(prog) - 1; i >= 0; i-- {
			if prog[i] != out[i] {
				A += int(math.Pow(8, float64(i)))
				continue OUT
			}
		}
	}
	return A
}
