package main

import (
	"advent_of_code_2024_go/src/utils"
	"testing"
)

func Test02_2(t *testing.T) {
	lines := utils.MustReadlines("input_test.txt")

	expected := 4
	got := Do(lines)

	if got != expected {
		t.Errorf("Do day01/part2 | want %d, got %d", expected, got)
	}
}
