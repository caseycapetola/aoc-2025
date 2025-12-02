package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func part1(input_file string) {
	file, err := os.Open(input_file)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	content := scanner.Text()

	res := 0

	contentArr := strings.Split(content, ",")
	for _, val := range contentArr {
		v := strings.Split(val, "-")
		l, err := strconv.Atoi(v[0])
		if err != nil {
			panic(err)
		}

		r, err := strconv.Atoi(v[1])
		if err != nil {
			panic(err)
		}
		
		for i := l; i<r+1; i++ {
			val := strconv.Itoa(i)
			mid := len(val)/2
			if val[:mid] == val[mid:] {
				res += i
			}
		}
	}
	fmt.Printf("Part 1: %d\n", res)
}

func part2(input_file string) {
	file, err := os.Open(input_file)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	content := scanner.Text()

	res := 0

	contentArr := strings.Split(content, ",")
	
	for _, val := range contentArr {
		invalidCount := 0
		v := strings.Split(val, "-")
		l, err := strconv.Atoi(v[0])
		if err != nil {
			panic(err)
		}

		r, err := strconv.Atoi(v[1])
		if err != nil {
			panic(err)
		}
		
		for i := l; i<r+1; i++ {
			val := strconv.Itoa(i)
			for j := 1; j<len(val)+1; j++ {
				substring := val[:j]

				arr := strings.Split(val, substring)
				if len(arr) < 3 {
					continue
				}

				flag := true
				for _, v := range arr {
					if v != "" {
						flag = false
					}
				}

				if flag {
					res += i
					invalidCount += 1
					break
				}
			}
		}
	}
	fmt.Printf("Part 2: %d\n", res)
}

func main() {
	part1("sample_input.txt")
	part1("input.txt")
	part2("sample_input.txt")
	part2("input.txt")
}