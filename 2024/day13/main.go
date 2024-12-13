package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
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

type equation struct {
	a, b, ans int
}

func partOne(input string) int{
	re, _ := regexp.Compile(`\d+`)
	sum := 0
	for _, machine := range strings.Split(input, "\n\n") {

		lines := strings.Split(machine, "\n")

		a   := re.FindAllString(lines[0], -1)
		b   := re.FindAllString(lines[1], -1)
		ans := re.FindAllString(lines[2], -1)

		x := equation{}
		y := equation{}
	
		x.a, _ = strconv.Atoi(a[0]); x.b, _ = strconv.Atoi(b[0]); x.ans, _ = strconv.Atoi(ans[0])
		y.a, _ = strconv.Atoi(a[1]); y.b, _ = strconv.Atoi(b[1]); y.ans, _ = strconv.Atoi(ans[1])

		A, B := solve(x, y)
		sum += A * 3 + B
	}
	return sum
}

func partTwo(input string) int{
	re, _ := regexp.Compile(`\d+`)
	sum := 0
	for _, machine := range strings.Split(input, "\n\n") {

		lines := strings.Split(machine, "\n")

		a   := re.FindAllString(lines[0], -1)
		b   := re.FindAllString(lines[1], -1)
		ans := re.FindAllString(lines[2], -1)

		x := equation{}
		y := equation{}
	
		x.a, _ = strconv.Atoi(a[0]); x.b, _ = strconv.Atoi(b[0]); x.ans, _ = strconv.Atoi(ans[0])
		y.a, _ = strconv.Atoi(a[1]); y.b, _ = strconv.Atoi(b[1]); y.ans, _ = strconv.Atoi(ans[1])
		
		x.ans += 10000000000000; y.ans += 10000000000000

		A, B := solve(x, y)
		sum += A * 3 + B
	}
	return sum
}


func solve(x, y equation) (a, b int) {

	tmpX := x.a
	tmpY := y.a

	x.a *= tmpY; x.b *= tmpY; x.ans *= tmpY
	y.a *= tmpX; y.b *= tmpX; y.ans *= tmpX

	y.a -= x.a; y.b -= x.b; y.ans -= x.ans


	if y.ans % y.b != 0 {
		return 0, 0
	}
	y.ans /= y.b
	x.ans -= (x.b * y.ans)

	if x.ans % x.a != 0 {
		return 0, 0
	}
	x.ans /= x.a
	return x.ans, y.ans
}
