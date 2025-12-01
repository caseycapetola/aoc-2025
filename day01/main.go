package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func part2() {
	dialPosition := 50
	res := 0
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		
		var dir int
		if line[0] == 'L' {
			dir = -1
		} else {
			dir = 1
		}
		mag, err := strconv.Atoi(line[1:])
		if err != nil {
			panic(err)
		}
		for mag > 0 {
			dialPosition = (dialPosition + dir) % 100
			if dialPosition == 0 {
				res += 1
			}
			mag -= 1
		}
	}

	fmt.Printf("Part 2 Result: %d\n", res)
}

func main() {
	dialPosition := 50
	res := 0

	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		
		var dir int
		if line[0] == 'L' {
			dir = -1
		} else {
			dir = 1
		}
		mag, err := strconv.Atoi(line[1:])
		if err != nil {
			panic(err)
		}
		movement := dir * mag
		
		dialPosition = (dialPosition + movement) % 100
		if dialPosition == 0 {
			res += 1
		}
	}

	fmt.Printf("Result: %d\n", res)
	part2()
}