package Days

import (
	"fmt"
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
	log.Println("Day 11 Part 1: ", day11_i32(linesInt, 25))
	// log.Println("Day 11 Part 2: ", day11(linesInt, 75))
}

func recurse(currNum int, depthRemaining int, stonesLastDepth *[]uint8, intSignedLookup map[int]uint8, intSignedLookupReverse map[uint8]int) {
	if depthRemaining <= 0 {
		(*stonesLastDepth) = append(*stonesLastDepth, addOrGetFromCache(intSignedLookup, intSignedLookupReverse, currNum))
		return
	}
	if currNum == 0 {
		m := make([]uint8, 1)
		m[0] = addOrGetFromCache(intSignedLookup, intSignedLookupReverse, 1)
		recurse(1, depthRemaining-1, stonesLastDepth, intSignedLookup, intSignedLookupReverse)
	} else {
		x := string(strconv.Itoa(currNum))
		if len(x)%2 == 0 {
			firstHalf := x[:len(x)/2]
			secondHalf := x[len(x)/2:]
			firstHalfInt, _ := strconv.Atoi(firstHalf)
			secondHalfInt, _ := strconv.Atoi(secondHalf)
			m := make([]uint8, 2)
			m[0] = addOrGetFromCache(intSignedLookup, intSignedLookupReverse, firstHalfInt)
			m[1] = addOrGetFromCache(intSignedLookup, intSignedLookupReverse, secondHalfInt)
			recurse(firstHalfInt, depthRemaining-1, stonesLastDepth, intSignedLookup, intSignedLookupReverse)
			recurse(secondHalfInt, depthRemaining-1, stonesLastDepth, intSignedLookup, intSignedLookupReverse)
		} else {
			m := make([]uint8, 1)
			currNum2024 := currNum * 2024
			uint8Val := addOrGetFromCache(intSignedLookup, intSignedLookupReverse, currNum2024)
			m[0] = uint8Val
			recurse(currNum2024, depthRemaining-1, stonesLastDepth, intSignedLookup, intSignedLookupReverse)
		}
	}
}

func addOrGetFromCache(intSignedLookup map[int]uint8, intSignedLookupReverse map[uint8]int, val int) uint8 {
	if valUint8, ok := intSignedLookup[val]; ok {
		return valUint8
	} else {
		// get highest value
		high := uint8(0)
		for _, v := range intSignedLookup {
			if v > high {
				high = v
			}
		}
		// auto increment
		newVal := high + 1
		intSignedLookup[val] = newVal
		intSignedLookupReverse[newVal] = val
		return intSignedLookup[val]
	}
}

func reverseCache(intSignedLookupReverse map[uint8]int, val uint8) int {
	if val, ok := intSignedLookupReverse[val]; ok {
		return val
	}
	panic(fmt.Sprintf("%v not in cache", val))
}

func day11(lines []int, depth int) int {
	res := 0
	cacheFor25Layers := make(map[uint8][]uint8)
	intSignedLookup := make(map[int]uint8)
	intSignedLookup[0] = 0
	intSignedLookup[1] = 1
	intSignedLookupReverse := make(map[uint8]int)
	intSignedLookupReverse[0] = 0
	intSignedLookupReverse[1] = 1
	for i := 0; i < len(lines); i++ {
		fmt.Println(i, " / ", len(lines)-1)
		stonei32 := lines[i]
		// #1: depth 1-25
		stones1 := make([]uint8, 0)
		recurse(stonei32, depth, &stones1, intSignedLookup, intSignedLookupReverse)
		cacheFor25Layers[addOrGetFromCache(intSignedLookup, intSignedLookupReverse, stonei32)] = stones1
		// #2: depth 26-50
		for _, stone := range stones1 {
			tmp, ok := cacheFor25Layers[stone]
			if !ok {
				// cache missing. calculate now
				stonesLastDepth := make([]uint8, 0)
				recurse(reverseCache(intSignedLookupReverse, stone), depth, &stonesLastDepth, intSignedLookup, intSignedLookupReverse)
				cacheFor25Layers[stone] = stonesLastDepth
				tmp = stonesLastDepth
			}
			// #3: depth 51-75
			for _, stonex := range tmp {
				tmp, ok := cacheFor25Layers[stonex]
				if !ok {
					// cache missing. calculate now
					stonesLastDepth := make([]uint8, 0)
					recurse(reverseCache(intSignedLookupReverse, stonex), depth, &stonesLastDepth, intSignedLookup, intSignedLookupReverse)
					cacheFor25Layers[stonex] = stonesLastDepth
					tmp = stonesLastDepth
				}
				res += len(tmp)
			}
		}
	}
	return res
}

func recurse_i32(currNum int, depthRemaining int, stonesLastDepth *[]int) {
	if depthRemaining <= 0 {
		(*stonesLastDepth) = append(*stonesLastDepth, currNum)
		return
	}
	if currNum == 0 {
		recurse_i32(1, depthRemaining-1, stonesLastDepth)
	} else {
		x := string(strconv.Itoa(currNum))
		if len(x)%2 == 0 {
			firstHalf := x[:len(x)/2]
			secondHalf := x[len(x)/2:]
			firstHalfInt, _ := strconv.Atoi(firstHalf)
			secondHalfInt, _ := strconv.Atoi(secondHalf)
			recurse_i32(firstHalfInt, depthRemaining-1, stonesLastDepth)
			recurse_i32(secondHalfInt, depthRemaining-1, stonesLastDepth)
		} else {
			recurse_i32(currNum*2024, depthRemaining-1, stonesLastDepth)
		}
	}
}

func day11_i32(lines []int, depth int) int {
	res := 0
	cacheFor25Layers := make(map[int][]int)
	for i := 0; i < len(lines); i++ {
		fmt.Println(i, " / ", len(lines)-1)
		stonei32 := lines[i]
		// #1: depth 1-25
		stones1 := make([]int, 0)
		recurse_i32(stonei32, depth, &stones1)
		cacheFor25Layers[stonei32] = stones1
		// #2: depth 26-50
		for _, stone := range stones1 {
			tmp, ok := cacheFor25Layers[stone]
			if !ok {
				// cache missing. calculate now
				stonesLastDepth := make([]int, 0)
				recurse_i32(stone, depth, &stonesLastDepth)
				cacheFor25Layers[stone] = stonesLastDepth
				tmp = stonesLastDepth
			}
			// #3: depth 51-75
			for _, stonex := range tmp {
				tmp, ok := cacheFor25Layers[stonex]
				if !ok {
					// cache missing. calculate now
					stonesLastDepth := make([]int, 0)
					recurse_i32(stonex, depth, &stonesLastDepth)
					cacheFor25Layers[stonex] = stonesLastDepth
					tmp = stonesLastDepth
				}
				res += len(tmp)
			}
		}
	}
	return res
}
