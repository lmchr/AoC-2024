package Days

import (
	"log"
	"strconv"
	"strings"
)

func Day2() {
	lines := readFile("data/day2.txt")
	log.Println("Day 2 Part 1: ", day2Part1(lines))
	log.Println("Day 2 Part 2: ", day2Part2(lines))
}

func preprocess(lines []string) [][]int {
	var out [][]int
	for _, line := range lines {
		fields := strings.Fields(line)
		var fieldsI []int
		for _, field := range fields {
			conv, _ := strconv.Atoi(field)
			fieldsI = append(fieldsI, conv)
		}
		out = append(out, fieldsI)
	}
	return out
}

func isSafe(line []int) bool {
	decreasing := line[0] > line[1]
	diff := 0
	safeIter := true
	for idy := range line[:len(line)-1] {
		diff = line[idy] - line[idy+1]
		if diff == 0 || // no diff
			diff > 3 || //diff too big
			diff < -3 || // diff too small
			(decreasing && diff < 0) || // was decreasing but is now increasing
			(!decreasing && diff > 0) { // was increasing but is now decreasing
			safeIter = false
			break
		}
	}
	return safeIter
}

func day2Part1(lines []string) int {
	linesInt := preprocess(lines)
	safeCount := 0
	safe := false
	for _, line := range linesInt {
		safe = isSafe(line)
		if safe {
			safeCount += 1
		}
	}
	return safeCount
}

func day2Part2(lines []string) int {
	linesInt := preprocess(lines)
	safeCount := 0
	safe := false
	for _, line := range linesInt {
		safe = false
		var tmp []int
		for idx := range line {
			// remove elements one by one
			tmp = make([]int, len(line))
			copy(tmp, line)
			tmp = removeFromList(tmp, idx)
			if isSafe(tmp) {
				safe = true
				break
			}
		}
		if safe {
			safeCount += 1
		}
	}
	return safeCount
}
