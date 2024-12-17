package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {

    inputFile, err := os.ReadFile("test")
    if err != nil {
        log.Fatal(err)
    }
    input := string(inputFile[:len(inputFile)-1])

    fmt.Printf("Part one: %s\n", partOne(input))
    fmt.Printf("Part two: %d\n", partTwo(input))
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

func partOne(input string) string {

	r, _ := regexp.Compile(`\d+`)

	A := 0
	B := 0
	C := 0

	tmp := strings.Split(input, "\n\n")
	regs := r.FindAllString(tmp[0], 3)
	A, _ = strconv.Atoi(regs[0])
	B, _ = strconv.Atoi(regs[1])
	C, _ = strconv.Atoi(regs[2])

	prog := []int{}
	for _, op  := range strings.Split(tmp[1][9:], ","){
		tmp, _ := strconv.Atoi(op)
		prog = append(prog, tmp)
	}
	out := strings.Builder{}

	for op := 0; op < len(prog); op++ {
		switch prog[op] {
		case 0: {
			tmp := cmb(prog[op + 1], &A, &B, &C)
			A = A / int(math.Pow(2, float64(tmp)))
			op++
			continue
		}
		case 1: {
			B = B ^ prog[op + 1]
			op++
			continue
		}
		case 2: {
			B = mod(cmb(prog[op + 1], &A, &B, &C), 8)
			op++
			continue
		}
		case 3: {
			if A == 0 {
				op++
				continue
			}
			op = prog[op + 1] - 1
		}
		case 4: {
			B = B ^ C
			op ++
			continue
		}
		case 5: {
			tmp := cmb(prog[op + 1], &A, &B, &C)
			tmp = mod(tmp, 8)
			out.WriteString(strconv.Itoa(tmp))
			out.WriteByte(',')
			op++
		}
		case 6: {
			tmp := cmb(prog[op + 1], &A, &B, &C)
			B = A / int(math.Pow(2, float64(tmp)))
			op++
			continue
		}
		case 7: {
			tmp := cmb(prog[op + 1], &A, &B, &C)
			C = A / int(math.Pow(2, float64(tmp)))
			op++
			continue
		}
		}
	}


	return out.String()[:out.Len()-1]
}

func partTwo(input string) int{
	r, _ := regexp.Compile(`\d+`)

	A := 0
	B := 0
	C := 0

	tmp := strings.Split(input, "\n\n")
	regs := r.FindAllString(tmp[0], 3)

	progStr := tmp[1][9:]
	prog := []int{}
	for _, op  := range strings.Split(progStr, ","){
		tmp, _ := strconv.Atoi(op)
		prog = append(prog, tmp)
	}
	out := strings.Builder{}

	
	a := int(math.Pow(8, float64(len(prog) - 1)))
	OUT:
	for {
		A = a

		B, _ = strconv.Atoi(regs[1])
		C, _ = strconv.Atoi(regs[2])


		for op := 0; op < len(prog); op++ {
			switch prog[op] {
			case 0: {
				tmp := cmb(prog[op + 1], &A, &B, &C)
				A = A / int(math.Pow(2, float64(tmp)))
				op++
				continue
			}
			case 1: {
				B = B ^ prog[op + 1]
				op++
				continue
			}
			case 2: {
				B = mod(cmb(prog[op + 1], &A, &B, &C), 8)
				op++
				continue
			}
			case 3: {
				if A == 0 {
					op++
					continue
				}
				op = prog[op + 1] - 1
			}
			case 4: {
				B = B ^ C
				op ++
				continue
			}
			case 5: {
				tmp := cmb(prog[op + 1], &A, &B, &C)
				tmp = mod(tmp, 8)
				out.WriteString(strconv.Itoa(tmp))
				out.WriteByte(',')
				op++
			}
			case 6: {
				tmp := cmb(prog[op + 1], &A, &B, &C)
				B = A / int(math.Pow(2, float64(tmp)))
				op++
				continue
			}
			case 7: {
				tmp := cmb(prog[op + 1], &A, &B, &C)
				C = A / int(math.Pow(2, float64(tmp)))
				op++
				continue
			}
			}
		}


		if out.Len() > 0 && out.String()[:out.Len()-1] == progStr { break }

		t1 := strings.Split(progStr, ",")
		t2 := strings.Split(out.String()[:out.Len()-1], ",")

		for i := len(t1) - 1; i >= 0; i-- {
			if t1[i] != t2[i] {
				a += int(math.Pow(8, float64(i)))
				out.Reset()
				continue OUT
			}

		}
		
	}

	

	return a
}


func mod(a, b int) int {
	tmp := a % b
	if tmp < 0 { return -tmp }
	return tmp
}
