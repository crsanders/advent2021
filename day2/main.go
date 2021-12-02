package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	part1, err := part1("input.txt")
	if err != nil {
		panic(err)
	}
	println("Part 1 Total:", part1)

	part2, err := part2("input.txt")
	if err != nil {
		panic(err)
	}
	println("Part 2 Total:", part2)
}

func part1(inputFile string) (total int, err error) {
	file, err := os.Open(inputFile)
	if err != nil {
		err = fmt.Errorf("could not open file: %w", err)
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var h, d int
	for scanner.Scan() {
		text := scanner.Text()
		slice := strings.Split(text, " ")

		dir := slice[0]
		dis := slice[1]
		disInt, convertErr := strconv.Atoi(dis)
		if err != nil {
			err = fmt.Errorf("could not convert string to int: %w", convertErr)
			return
		}

		switch dir {
		case "forward":
			h += disInt
		case "down":
			d += disInt
		case "up":
			d -= disInt
		}
	}

	total = h * d
	return
}

func part2(inputFile string) (total int, err error) {
	file, err := os.Open(inputFile)
	if err != nil {
		err = fmt.Errorf("could not open file: %w", err)
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var h, d, a int
	for scanner.Scan() {
		text := scanner.Text()
		slice := strings.Split(text, " ")

		dir := slice[0]
		dis := slice[1]
		disInt, convertErr := strconv.Atoi(dis)
		if err != nil {
			err = fmt.Errorf("could not convert string to int: %w", convertErr)
			return
		}

		switch dir {
		case "forward":
			h += disInt
			d += a * disInt
		case "down":
			a += disInt
		case "up":
			a -= disInt
		}
	}

	total = h * d
	return
}
