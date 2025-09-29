package main

import (
	"advent_of_code_2024_go/src/utils"
	"testing"
)

func Test01_1(t *testing.T) {
	lines := utils.MustReadlines("input_test.txt")

	expected := 11
	got := Do(lines)

	if got != expected {
		t.Errorf("Do 01_01 | want %d, got %d", expected, got)
	}
}
