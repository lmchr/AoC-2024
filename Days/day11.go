package Days

import (
	"log"
	"strconv"
	"strings"
)

func Day11() {
	lines := readFile("data/day11.txt")[0]
	var linesInt []int
	for _, line := range strings.Split(lines, " ") {
		linesSplitInt, _ := strconv.Atoi(line)
		linesInt = append(linesInt, linesSplitInt)
	}
	log.Println("Day 11 Part 1: ", day11(linesInt, 25))
	log.Println("Day 11 Part 2: ", day11(linesInt, 75))
}

func recurse(currNum int, numStones *int, depthRemaining int, splitCache map[int][]int) {
	if depthRemaining <= 0 {
		return
	}
	if val, ok := splitCache[currNum]; ok {
		if len(val) == 2 {
			(*numStones) += 1
			recurse(val[0], numStones, depthRemaining-1, splitCache)
			recurse(val[1], numStones, depthRemaining-1, splitCache)
		} else {
			recurse(val[0], numStones, depthRemaining-1, splitCache)
		}
	} else {
		if currNum == 0 {
			m := make([]int, 1)
			m[0] = 1
			splitCache[currNum] = m
			recurse(1, numStones, depthRemaining-1, splitCache)
		} else {
			x := string(strconv.Itoa(currNum))
			if len(x)%2 == 0 {
				firstHalf := x[:len(x)/2]
				secondHalf := x[len(x)/2:]
				firstHalfInt, _ := strconv.Atoi(firstHalf)
				secondHalfInt, _ := strconv.Atoi(secondHalf)
				(*numStones) += 1
				m := make([]int, 2)
				m[0] = firstHalfInt
				m[1] = secondHalfInt
				splitCache[currNum] = m
				recurse(firstHalfInt, numStones, depthRemaining-1, splitCache)
				recurse(secondHalfInt, numStones, depthRemaining-1, splitCache)
			} else {
				m := make([]int, 1)
				m[0] = currNum * 2024
				splitCache[currNum] = m
				recurse(currNum*2024, numStones, depthRemaining-1, splitCache)
			}
		}
	}
}

func day11(lines []int, depth int) int {
	res := 0
	splitCache := make(map[int][]int)
	for i := len(lines) - 1; i >= 0; i-- {
		log.Println(i)
		val := lines[i]
		res += 1
		recurse(val, &res, depth, splitCache)
	}
	return res
}
