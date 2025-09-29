package main

import (
	"advent_of_code_2024_go/src/utils"
	"fmt"
	"sort"
	"strings"
)

func main() {
	lines := utils.MustReadlines("input.txt")
	res := Do(lines)
	fmt.Printf("The answer: %d\n", res)
}

func Do(lines []string) int {
	left := make([]int, len(lines))
	right := make([]int, len(lines))
	for i, n := range lines {
		parts := strings.Split(n, "   ")
		left[i] = utils.MustAtoi(parts[0])
		right[i] = utils.MustAtoi(parts[1])
	}
	sort.Ints(left)
	sort.Ints(right)
	sum := 0
	for i := range left {
		sum += utils.Abs(left[i] - right[i])
	}
	return sum
}
