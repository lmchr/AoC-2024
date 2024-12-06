package Days

import (
	"log"
)

func Day4() {
	lines := readFile("data/day4.txt")
	log.Println("Day 4 Part 1: ", day4Part1(lines))
	log.Println("Day 4 Part 2: ", day4Part2(lines))
}

func day4Part1(lines []string) int {
	needle := "XMAS"
	needleReversed := "SAMX"
	// horizontal
	result := 0
	for _, line := range lines {
		for i := 0; i <= len(line)-len(needle); i++ {
			s := line[i : i+len(needle)]
			result += Bool2int(needle == s)
			result += Bool2int(needleReversed == s)
		}
	}
	// vertical
	for j := 0; j < len(lines[0]); j++ {
		for i := 0; i <= len(lines)-len(needle); i++ {
			s := string(lines[i][j]) + string(lines[i+1][j]) + string(lines[i+2][j]) + string(lines[i+3][j])
			result += Bool2int(needle == s)
			result += Bool2int(needleReversed == s)
		}
	}
	// diagonal upper left to lower right
	for i := 0; i <= len(lines)-len(needle); i++ {
		for j := 0; j <= len(lines[0])-len(needle); j++ {
			s := string(lines[i][j]) + string(lines[i+1][j+1]) + string(lines[i+2][j+2]) + string(lines[i+3][j+3])
			result += Bool2int(needle == s)
			result += Bool2int(needleReversed == s)
		}
	}
	//swap whole field to check other diagonale
	for idx := range lines {
		lines[idx] = Reverse(lines[idx])
	}
	// diagonal upper right to lower left
	for i := 0; i <= len(lines)-len(needle); i++ {
		for j := 0; j <= len(lines[0])-len(needle); j++ {
			s := string(lines[i][j]) + string(lines[i+1][j+1]) + string(lines[i+2][j+2]) + string(lines[i+3][j+3])
			result += Bool2int(needle == s)
			result += Bool2int(needleReversed == s)
		}
	}
	return result
}

func day4Part2(lines []string) int {
	result := 0
	for i := 0; i <= len(lines)-3; i++ {
		for j := 0; j <= len(lines[0])-3; j++ {
			middle := string(lines[i+1][j+1])
			upperLeft := string(lines[i][j])
			upperRight := string(lines[i][j+2])
			lowerLeft := string(lines[i+2][j])
			lowerRight := string(lines[i+2][j+2])
			if middle == "A" && ((upperLeft == "M" && lowerRight == "S" && upperRight == "M" && lowerLeft == "S") || // MAS / MAS
				(upperLeft == "M" && lowerRight == "S" && upperRight == "S" && lowerLeft == "M") || // MAS / SAM
				(upperLeft == "S" && lowerRight == "M" && upperRight == "M" && lowerLeft == "S") || // SAM / MAS
				(upperLeft == "S" && lowerRight == "M" && upperRight == "S" && lowerLeft == "M")) { // SAM / SAM
				result += 1
			}
		}
	}
	return result
}
