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
	log.Println("Day 11 Part 1: ", day11(linesInt, true))
	log.Println("Day 11 Part 2: ", day11(linesInt, false))
}

func recurse(currNum int, depthRemaining int, stonesLastDepth *[]int) {
	if depthRemaining <= 0 {
		(*stonesLastDepth) = append(*stonesLastDepth, currNum)
		return
	}
	if currNum == 0 {
		recurse(1, depthRemaining-1, stonesLastDepth)
	} else {
		x := string(strconv.Itoa(currNum))
		if len(x)%2 == 0 {
			firstHalf := x[:len(x)/2]
			secondHalf := x[len(x)/2:]
			firstHalfInt, _ := strconv.Atoi(firstHalf)
			secondHalfInt, _ := strconv.Atoi(secondHalf)
			recurse(firstHalfInt, depthRemaining-1, stonesLastDepth)
			recurse(secondHalfInt, depthRemaining-1, stonesLastDepth)
		} else {
			recurse(currNum*2024, depthRemaining-1, stonesLastDepth)
		}
	}
}

func day11(lines []int, part1 bool) int {
	res := 0
	var depth int
	if part1 {
		depth = 25
	} else {
		depth = 75
	}
	cacheFor25Layers := make(map[int][]int)
	for i := 0; i < len(lines); i++ {
		stonei32 := lines[i]
		// #1: depth 1-25
		stones1 := make([]int, 0)
		recurse(stonei32, depth, &stones1)
		if part1 {
			res += len(stones1)
			continue
		}
		cacheFor25Layers[stonei32] = stones1
		// #2: depth 26-50
		for _, stone := range stones1 {
			tmp, ok := cacheFor25Layers[stone]
			if !ok {
				// cache missing. calculate now
				stonesLastDepth := make([]int, 0)
				recurse(stone, depth, &stonesLastDepth)
				cacheFor25Layers[stone] = stonesLastDepth
				tmp = stonesLastDepth
			}
			// #3: depth 51-75
			for _, stonex := range tmp {
				tmp, ok := cacheFor25Layers[stonex]
				if !ok {
					// cache missing. calculate now
					stonesLastDepth := make([]int, 0)
					recurse(stonex, depth, &stonesLastDepth)
					cacheFor25Layers[stonex] = stonesLastDepth
					tmp = stonesLastDepth
				}
				res += len(tmp)
			}
		}
	}
	return res
}
