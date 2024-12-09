package main
import (
	"fmt"
	"log"
	"os"
	"slices"
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

	disk := []int{}
	fileID := 0

	for i := 0; i < len(input); i++ {
		size := int(input[i] - '0')
		if i % 2 == 0 {
			for j := 0; j < size; j++ {
				disk = append(disk, fileID)
			}
			fileID++

		} else {
			for j := 0; j < size; j++ {
				disk = append(disk, -1)
			}
		}
	}

	empty := 0
	last := len(disk) -1
	
	for empty < last {

		if disk[last] == -1 { last-- }
		if disk[empty] != -1 { empty++ }
		if disk[empty] == -1 && disk[last] != -1 {
			disk[empty] = disk[last]
			disk[last] = -1
			empty++
			last--
		}
	}

	for i, val := range disk {
		if val == -1 {
			continue
		}
		sum += i*val
	}

    return sum
}

type block struct {id, size int}

func partTwo(input string) int{
    sum := 0

	disk := []block{}
	fileID := 0

	for i := 0; i < len(input); i++ {
		size := int(input[i] - '0')
		if i % 2 == 0 {
			disk = append(disk, block{id: fileID, size: size})
			fileID++
		} else {
			disk = append(disk, block{id: -1, size: size})
		}
	}

	for i := len(disk) - 1; i >= 0; i-- { // last non empty element
		if disk[i].id != -1 {
			for j := 0; j < i; j++ {
				if disk[j].id == -1 && disk[j].size >= disk[i].size {
					tmp := disk[i]
					disk[i].id = -1
					disk[j].size -= disk[i].size
					disk = slices.Insert(disk, j, tmp)
					break
				}
			}
		}
	}

	i := 0
	for _, val := range disk { // working
		for j := 0; j < val.size; j++ {
			if val.id != -1 {
				sum += val.id * i
			}
			i++
		}
	}

    return sum
}
