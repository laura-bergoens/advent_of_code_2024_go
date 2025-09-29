package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	lines := MustReadlines("./inputs/01_1.txt")
	slice := lines[:]
	left := make([]int, len(slice))
	right := make([]int, len(slice))
	for i, n := range slice {
		parts := strings.Split(n, "   ")
		left[i] = mustAtoi(parts[0])
		right[i] = mustAtoi(parts[1])
	}
	sort.Ints(left)
	sort.Ints(right)
	sum := 0
	for i := range len(left) {
		sum += Abs(left[i] - right[i])
	}
	fmt.Printf("The answer: %d\n", sum)
}
