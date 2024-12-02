package main

import (
	"os"
	"log"
	"testing"
)

func TestDay02A(t *testing.T) {
    inputFile, err := os.ReadFile("example")
    if err != nil {
        log.Fatal(err)
    }
    input := string(inputFile[:len(inputFile)-1])


}
