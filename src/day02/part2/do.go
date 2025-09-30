package main

import (
	"advent_of_code_2024_go/src/utils"
	"fmt"
	"strings"
)

func main() {
	lines := utils.MustReadlines("src/day02/part2/input.txt")
	res := Do(lines)
	fmt.Printf("The answer: %d\n", res)
}

func Do(lines []string) int {
	safeCnt := 0
	for _, n := range lines {
		parts := strings.Split(n, " ")
		allPoss := generate_all_poss(parts)
		for _, poss := range allPoss {
			decreasing := false
			previous := 0
			isSafe := true
			for i, m := range poss {
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
				break
			}
		}
	}
	return safeCnt
}

func generate_all_poss(input []string) [][]string {
	res := make([][]string, len(input)+1)
	for i := range input {
		poss := make([]string, len(input)-1)
		copy(poss, input[:i])
		copy(poss[i:], input[i+1:])
		res[i] = poss
	}
	res[len(input)] = input
	return res
}
