package main

import (
	"bufio"
	"os"
	"strconv"
)

func main() {
	// Part 1
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

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
	var tmp, tmp2, prev2, prev3, i int
	scanner2 := bufio.NewScanner(file2)
	for scanner2.Scan() {
		cur := scanner2.Text()
		curInt, err := strconv.Atoi(cur)
		if err != nil {
			panic(err)
		}
		if i == 0 {
			prev = curInt
			i++
			continue
		}
		if i == 1 {
			prev2 = curInt
			i++
			continue
		}
		if i == 2 {
			prev3 = curInt
			tmp = prev + prev2 + prev3
			i++
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
