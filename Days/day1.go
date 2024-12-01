package Days

import (
	"log"
	"sort"
	"strconv"
	"strings"
)

func Day1() {
	lines := readFile("data/day1.txt")
	log.Println("Day 1 Part 1: ", part1(lines))
	log.Println("Day 1 Part 2: ", part2(lines))
}

func getLists(lines []string) ([]int, []int) {
	var a, b []int
	for _, line := range lines {
		var x = strings.Split(line, " ")
		ele1, _ := strconv.Atoi(x[0])
		ele2, _ := strconv.Atoi(x[len(x)-1])
		a = append(a, ele1)
		b = append(b, ele2)
	}
	return a, b
}

func part1(lines []string) int {
	a, b := getLists(lines)
	sort.Ints(a)
	sort.Ints(b)
	res := 0
	for idx := range a {
		diff := a[idx] - b[idx]
		if diff < 0 {
			diff *= -1
		}
		res += diff
	}
	return res
}

func part2(lines []string) int {
	a, b := getLists(lines)
	sort.Ints(a)
	m := make(map[int]int)
	for _, ele := range b {
		m[ele] += 1
	}
	res := 0
	for _, ele := range a {
		res += ele * m[ele]
	}
	return res
}
