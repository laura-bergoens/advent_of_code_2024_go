package main

import (
	"advent_of_code_2024_go/src/utils"
	"fmt"
	"strings"
)

func main() {
	lines := utils.MustReadlines("src/day02/part1/input.txt")
	res := Do(lines)
	fmt.Printf("The answer: %d\n", res)
}

func Do(lines []string) int {
	safeCnt := 0
	for _, n := range lines {
		parts := strings.Split(n, " ")
		decreasing := false
		previous := 0
		isSafe := true
		for i, m := range parts {
			iM := utils.MustAtoi(m)
			if i == 0 {
				previous = iM
				continue
			}
			if i == 1 {
				if previous == iM {
					isSafe = false
					break
				}
				decreasing = previous > iM
			}

			if decreasing && iM >= previous {
				isSafe = false
				break
			}
			if !decreasing && iM <= previous {
				isSafe = false
				break
			}
			abs := utils.Abs(previous - iM)
			if abs < 1 || abs > 3 {
				isSafe = false
				break
			}
			previous = iM
		}
		if isSafe {
			safeCnt++
		}
	}
	return safeCnt
}
