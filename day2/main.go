package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Part 1
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var h, d int
	for scanner.Scan() {
		text := scanner.Text()
		slice := strings.Split(text, " ")

		dir := slice[0]
		dis := slice[1]
		disInt, err := strconv.Atoi(dis)
		if err != nil {
			panic(err)
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

	total := h * d
	println(total)

	// Part 2
	file2, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner2 := bufio.NewScanner(file2)

	h = 0
	d = 0
	total = 0
	var a int
	for scanner2.Scan() {
		text := scanner2.Text()
		slice := strings.Split(text, " ")

		dir := slice[0]
		dis := slice[1]
		disInt, err := strconv.Atoi(dis)
		if err != nil {
			panic(err)
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
	println(total)
}
