package Days

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

type Line struct {
	number  int
	numbers []int
}

func Day7() {
	lines := readFile("data/day7.txt")
	log.Println("Day 7 Part 1: ", day7Part1And2(lines, 2))
	log.Println("Day 7 Part 2: ", day7Part1And2(lines, 3))
}

func itertoolsCombinations(nNumbers int, numUniqueValues int) [][]int {
	// e.g. nNumbers == 5, numUniqueValues == 2
	// calculate numUniqueValues**nNumbers = 2**5 = 32 (=11111 in binary)
	// iterate from 0 to 32-1 (00000 to 11111) which will create all possible combinations
	var combinations [][]int
	for i := range IntPow(numUniqueValues, nNumbers) {
		combination := make([]int, nNumbers)
		binaryRepresentation := strconv.FormatInt(int64(i), numUniqueValues)
		for j := 0; j < len(binaryRepresentation); j++ {
			val, _ := strconv.Atoi(string(binaryRepresentation[j]))
			combination[len(binaryRepresentation)-1-j] = val
		}
		combinations = append(combinations, combination)
	}
	return combinations
}

func parseInput(lines []string) []Line {
	parsed := make([]Line, len(lines))
	for idx, line := range lines {
		var numbers []int
		splitColon := strings.Split(line, ": ")
		for _, s := range strings.Split(splitColon[1], " ") {
			s_, _ := strconv.Atoi(string(s))
			numbers = append(numbers, s_)
		}
		splitColon_, _ := strconv.Atoi(string(splitColon[0]))
		parse := Line{number: splitColon_, numbers: numbers}
		parsed[idx] = parse
	}
	return parsed
}

func day7Part1And2(lines []string, base int) int {
	parsed := parseInput(lines)
	// precompute operatorCombinations in binary/ternary numeral system
	operatorCombinationsLookup := make(map[int][][]int)
	for _, parse := range parsed {
		if _, ok := operatorCombinationsLookup[len(parse.numbers)]; !ok {
			operatorCombinationsLookup[len(parse.numbers)] = itertoolsCombinations(len(parse.numbers), base)
		}
	}
	totalSum := 0
	for _, parse := range parsed {
		operatorCombinations := operatorCombinationsLookup[len(parse.numbers)]
		for _, operatorCombination := range operatorCombinations {
			resCurrentOperatorCombination := parse.numbers[0]
			for i := 1; i < len(operatorCombination); i++ {
				switch operatorCombination[i] {
				case 0:
					resCurrentOperatorCombination += parse.numbers[i]
				case 1:
					resCurrentOperatorCombination *= parse.numbers[i]
				case 2:
					concatenatedNumbers := fmt.Sprintf("%d%d", resCurrentOperatorCombination, parse.numbers[i])
					concatenatedNumbersNum, _ := strconv.Atoi(concatenatedNumbers)
					resCurrentOperatorCombination = concatenatedNumbersNum
				default:
					panic("Unexpected operator")
				}
			}
			// add to result if this combination worked
			if resCurrentOperatorCombination == parse.number {
				totalSum += parse.number
				break
			}
		}

	}
	return totalSum
}
