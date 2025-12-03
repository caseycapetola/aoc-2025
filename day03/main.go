package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func part1(inputFile string) {
	file, _ := os.Open(inputFile)
	defer file.Close()

	res := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		l := 0
		r := 1

		joltage, _ := strconv.Atoi(string(line[l])+string(line[r]))

		for r < len(line) {
			lval, _ := strconv.Atoi(string(line[l]))
			rval, _ := strconv.Atoi(string(line[r]))
			j, _ := strconv.Atoi(string(line[l])+string(line[r]))
			joltage = max(joltage, j)
			if rval > lval && rval + 1 < len(line) {
				l = r
			}
			r += 1
		}
		res += joltage
	}
	fmt.Printf("Part 1: %d\n", res)
}

func byteToInt(b byte) int {
	return int(b-'0')
}

func subsearch(input string, numElements int) string {
	if numElements == 0 {
		return ""
	}
	if len(input) == numElements {
		return input
	}
	index := 0
	max := 0
	for i:=0; i<len(input)-numElements+1; i++ {
		if (byteToInt(input[i]) > max) {
				index = i
				max = byteToInt(input[i])
			}
	}
	return string(input[index]) + subsearch(input[index+1:], numElements-1)
}

func part2(inputFile string, numBatteries int) {
	file, _ := os.Open(inputFile)
	defer file.Close()

	res := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		joltage, _ := strconv.Atoi(subsearch(line, numBatteries))
		res += joltage
		
	}
	fmt.Printf("Part2: %d\n", res)
}

func main() {
	part1("sample_input.txt")
	part1("input.txt")
	part2("sample_input.txt", 12)
	part2("input.txt", 12)
}
