package Days

import (
	"log"
	"regexp"
	"strconv"
	"strings"
)

func Day3() {
	lines := readFile("data/day3.txt")
	log.Println("Day 3 Part 1: ", day3Part1(lines))
	log.Println("Day 3 Part 2: ", day3Part2(lines))
}

func day3Part1(lines []string) int {
	r, _ := regexp.Compile(`mul\(\d{1,3}\,\d{1,3}\)`)
	res := 0
	line := strings.Join(lines, "")
	matches := r.FindAllString(line, -1)
	for _, match := range matches {
		match = strings.TrimPrefix(match, "mul(")
		match = strings.TrimSuffix(match, ")")
		numbers := strings.Split(match, ",")
		num1, _ := strconv.Atoi(numbers[0])
		num2, _ := strconv.Atoi(numbers[1])
		res += num1 * num2
	}
	return res
}

func day3Part2(lines []string) int {
	r, _ := regexp.Compile(`mul\(\d{1,3}\,\d{1,3}\)`)
	res := 0
	line := strings.Join(lines, "")
	matches := r.FindAllStringSubmatchIndex(line, -1)
	for _, matchTuple := range matches {
		matchStart := matchTuple[0]
		// find previous enable / disable. Only search before this match
		enableIdx := strings.LastIndex(line[:matchStart], "do()")
		disableIdx := strings.LastIndex(line[:matchStart], "don't()")
		if disableIdx == -1 || disableIdx < enableIdx {
			matchEnd := matchTuple[1]
			match := line[matchStart:matchEnd]
			match = strings.TrimPrefix(match, "mul(")
			match = strings.TrimSuffix(match, ")")
			numbers := strings.Split(match, ",")
			num1, _ := strconv.Atoi(numbers[0])
			num2, _ := strconv.Atoi(numbers[1])
			res += num1 * num2
		}

	}
	return res
}
