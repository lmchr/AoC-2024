package Days

import (
	"log"
	"slices"
	"strconv"
	"strings"
)

func Day5() {
	lines := readFile("data/day5.txt")
	correctLinesIdx, result := day5Part1(lines)
	log.Println("Day 5 Part 1: ", result)
	log.Println("Day 5 Part 2: ", day5Part2(lines, correctLinesIdx))
}

func readInput(lines []string) (map[int][]int, [][]int) {
	rules := make(map[int][]int)
	var pages [][]int
	toggle := false
	for _, line := range lines {
		if line == "" {
			toggle = true
			continue
		}
		if !toggle {
			split := strings.Split(line, "|")
			val1, _ := strconv.Atoi(split[0])
			val2, _ := strconv.Atoi(split[1])
			rules[val1] = append(rules[val1], val2)
		} else {
			splits := strings.Split(line, ",")
			var splitsInt = []int{}
			for _, split := range splits {
				j, _ := strconv.Atoi(split)
				splitsInt = append(splitsInt, j)
			}
			pages = append(pages, splitsInt)
		}
	}
	return rules, pages
}

func day5Part1(lines []string) ([]int, int) {
	var correctLinesIdx []int
	rules, pages := readInput(lines)
	for idxPage, page := range pages {
		pageValid := true
		for idx, p := range page {
			if idx == 0 {
				continue
			}
			// check if there are rules for the current number
			if pageNumbersMustAppearAfter, ok := rules[p]; ok {
				pageNumbersAfter := page[idx+1:]
				for _, pageNumberMustAppearAfter := range pageNumbersMustAppearAfter {
					// both pages are present && the second number of the rule after the first one
					if slices.Contains(page, pageNumberMustAppearAfter) && !slices.Contains(pageNumbersAfter, pageNumberMustAppearAfter) {
						pageValid = false
						break
					}
				}
			}
			if !pageValid {
				break
			}
		}
		if pageValid {
			correctLinesIdx = append(correctLinesIdx, idxPage)
		}
	}
	middleNumbersSum := 0
	for idx, page := range pages {
		if slices.Contains(correctLinesIdx, idx) {
			middleNumbersSum += page[len(page)/2]
		}
	}
	return correctLinesIdx, middleNumbersSum
}

func day5Part2(lines []string, correctLinesIdx []int) int {
	rules, pages := readInput(lines)
	// get correct lines
	middleNumbersSum := 0
	for idx, page := range pages {
		// skip correct lines since we can't fix those
		if slices.Contains(correctLinesIdx, idx) {
			continue
		}
		for idx, p := range page {
			if idx == 0 {
				continue
			}
			// check if there are rules for the current number
			if pageNumbersMustAppearAfter, ok := rules[p]; ok {
				var idxOfPageNumbersMustAppearAfter []int
				for _, pageNumberMustAppearAfter := range pageNumbersMustAppearAfter {
					idy := indexOf(pageNumberMustAppearAfter, page)
					if idy != -1 {
						idxOfPageNumbersMustAppearAfter = append(idxOfPageNumbersMustAppearAfter, idy)
					}
				}
				if len(idxOfPageNumbersMustAppearAfter) > 0 {
					// swap?
					minn := slices.Min(idxOfPageNumbersMustAppearAfter)
					if minn < idx {
						page = removeFromList(page, idx)
						page = slices.Insert(page, minn, p)
					}
				}
			}
		}
		middleNumbersSum += page[len(page)/2]
	}
	return middleNumbersSum
}
