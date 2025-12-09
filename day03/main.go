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

type Battery struct {
	index int
	value int
}

// func iterativePart2(inputFile string, numBatteries int) {
// 	file, _ := os.Open(inputFile)
// 	defer file.Close()

// 	res := 0

// 	scanner := bufio.NewScanner(file)
// 	for scanner.Scan() {
// 		line := scanner.Text()

// 		l, r := 0, 0

// 		batteries := []byte{}

// 		for r < len(line) {
// 			temp := min(numBatteries, len(line)-r)
// 			fmt.Printf("Batteries: %s; value to compare: %d; temp: %d\n", string(batteries[:]), byteToInt(line[r]), temp)
// 			if len(batteries) + (len(line)-r) == numBatteries { // Append the remaining content
// 				fmt.Println("APPENDING REMAINING CONTENT")
// 				for r < len(line) {
// 					batteries = append(batteries, line[r])
// 					r += 1
// 				}
// 			} else if len(batteries) > 0 && byteToInt(batteries[numBatteries-temp]) < byteToInt(line[r]) {
// 				batteries = append(batteries[:numBatteries-temp], line[r])
// 			} else if len(batteries) < numBatteries {
// 				fmt.Println("IN DEFAULT")
// 				l = r
// 				for len(line)-r > temp {
// 					if byteToInt(line[r]) > byteToInt(line[l]) {
// 						l = r
// 						break
// 					}
// 					r += 1
// 				}
				
// 				batteries = append(batteries, line[l])
// 			}
// 			r+=1
// 		}

// 		// Let l = the index of the start of the non-increasing subsequence
// 		for r < len(line) {
// 			temp := min(numBatteries, len(line)-r)

// 			if len(batteries) > 0 && byteToInt(batteries[numBatteries-temp]) < byteToInt(line[r]) { // Handles remainder stuff
// 				batteries = append(batteries[:numBatteries-temp], line[r])
// 			} else if len(batteries) == 0 { // Need to add an initial value to the slice
// 				batteries = append(batteries, line[r])
// 			} else if byteToInt(line[r]) > byteToInt(line[r-1]){
// 				// We will walk through our array to find the optimal place to put this value
// 				for byteToInt(line[l]) >= byteToInt(line[r]) &&  {
// 					l += 1
// 				}
// 			}

// 			r += 1
// 		}

// 		joltage, _ := strconv.Atoi(string(batteries[:]))
// 		fmt.Printf("Joltage: %d\n", joltage)
// 		res += joltage
		
// 	}
// 	fmt.Printf("Part2: %d\n", res)
// }

func main() {
	part1("sample_input.txt")
	part1("input.txt")
	part2("sample_input.txt", 12)
	part2("input.txt", 12)
	// iterativePart2("one_line.txt", 12)
}
