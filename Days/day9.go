package Days

import (
	"log"
	"strconv"
)

func Day9() {
	lines := readFile("data/day9.txt")
	log.Println("Day 9 Part 1: ", day9Part1(lines[0]))
	log.Println("Day 9 Part 2: ", day9Part2(lines[0]))
}

func deFrag(disk *[]int) {
	for i := 0; i < len(*disk); i++ {
		if (*disk)[i] != -1 {
			continue
		}
		// find the right most value that is not -1
		val := -1
		for j := len(*disk) - 1; j > i; j-- {
			if (*disk)[j] != -1 {
				val = (*disk)[j]
				// switch positions, move to front first empty space
				(*disk)[i] = val
				(*disk)[j] = -1
				break
			}
		}
		// nothing left
		if val == -1 {
			break
		}
	}
}

func deFrag2(disk *[]int) {
	for i := len(*disk) - 1; i > 0; i-- {
		if (*disk)[i] == -1 {
			continue
		}
		// check how far the group goes on
		startIdx := i
		var endIdx int
		for j := i - 1; j > 0; j-- {
			if (*disk)[startIdx] != (*disk)[j] {
				endIdx = j
				break
			}
		}
		groupLength := startIdx - endIdx
		// find a free group
		startIdxFree := -1
		endIdxFree := -1
		for z := 0; z < i; z++ {
			val := (*disk)[z]
			if val == -1 && startIdxFree == -1 {
				startIdxFree = z
			}
			if val != -1 && startIdxFree != -1 {
				if z-startIdxFree >= groupLength {
					endIdxFree = z
					break
				}
				// not big enough, reset start
				startIdxFree = -1
			}
		}
		if endIdxFree != -1 {
			for z := 0; z < groupLength; z++ {
				(*disk)[startIdxFree+z] = (*disk)[startIdx-z]
				(*disk)[startIdx-z] = -1
			}
		} else {
			//skip the current group
			i -= (groupLength - 1)
		}
	}
}

func gatherDisk(line string) []int {
	var disk []int
	for i := 0; i < len(line); i++ {
		iInt, _ := strconv.Atoi(string(line[i]))
		var val int
		if i%2 == 0 {
			val = i / 2
		} else {
			val = -1
		}
		for j := 0; j < iInt; j++ {
			disk = append(disk, val)
		}
	}
	return disk
}

func day9Part1(line string) int {
	disk := gatherDisk(line)
	deFrag(&disk)
	result := 0
	for idx, o := range disk {
		if o == -1 {
			break
		}
		result += idx * o
	}
	return result
}

func day9Part2(line string) int {
	disk := gatherDisk(line)
	deFrag2(&disk)
	result := 0
	for idx, o := range disk {
		if o == -1 {
			continue
		}
		result += idx * o
	}
	return result
}
