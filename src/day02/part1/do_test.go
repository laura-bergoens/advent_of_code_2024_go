package main

import (
	"advent_of_code_2024_go/src/utils"
	"fmt"
	"testing"
)

func Test02_1(t *testing.T) {
	lines := utils.MustReadlines("input_test.txt")

	expected := 2
	fmt.Println(lines)
	got := Do(lines)

	if got != expected {
		t.Errorf("Do day01/part2 | want %d, got %d", expected, got)
	}
}
