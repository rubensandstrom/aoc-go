package main

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func TestPartOne(t *testing.T) {

	input := `Comet can fly 14 km/s for 10 seconds, but then must rest for 127 seconds.
	Dancer can fly 16 km/s for 11 seconds, but then must rest for 162 seconds.`

	sum := [2]int{}
	for l, line := range strings.Split(input, "\n") {
		words := strings.Split(line, " ")

		speed, _ := strconv.Atoi(words[3])
		fly, _ := strconv.Atoi(words[6])
		rest, _ := strconv.Atoi(words[13])

		j := 0
		for i := 0; i < 1000; i++ {
			if j < fly {
				sum[l] += speed
				j++
				continue
			}
			if j < fly + rest {
				j++
				continue
			}
			j = 0
		}
	}
	fmt.Printf("%v\n", sum)
	if sum[0] != 1120 {
		t.Fatalf("expected 1120, got %d\n", sum[0])
	}
	if sum[1] != 1056 {
		t.Fatalf("expected 1056, got %d\n", sum[1])
	}
}
