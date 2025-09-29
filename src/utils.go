package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func Readlines(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, strings.TrimSpace(scanner.Text()))
	}

	return lines, scanner.Err()
}

func MustReadlines(filename string) []string {
	lines, err := Readlines(filename)
	if err != nil {
		panic(err)
	}
	return lines
}

func mustAtoi(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return n
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
