package main

import (
	"bufio"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	// Part 1
	var sum, prev int
	for scanner.Scan() {
		cur := scanner.Text()
		curInt, err := strconv.Atoi(cur)
		if err != nil {
			panic(err)
		}
		if prev == 0 {
			prev = curInt
			continue
		}

		if curInt > prev {
			sum++
		}
		prev = curInt
	}

	println(sum)

	// Part 2
	file2, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file2.Close()

	sum = 0
	prev = 0
	var tmp, tmp2, prev2, prev3 int
	scanner2 := bufio.NewScanner(file2)
	for scanner2.Scan() {
		cur := scanner2.Text()
		curInt, err := strconv.Atoi(cur)
		if err != nil {
			panic(err)
		}
		if prev == 0 {
			prev = curInt
			continue
		}
		if prev2 == 0 {
			prev2 = curInt
			continue
		}
		if prev3 == 0 {
			prev3 = curInt
			tmp = prev + prev2 + prev3
			continue
		}

		prev = prev2
		prev2 = prev3
		prev3 = curInt
		tmp2 = prev + prev2 + prev3

		if tmp2 > tmp {
			sum++
		}
		tmp = tmp2
	}

	println(sum)
}
