package Days

import (
	"log"
	"strconv"
)

func linesToNumbers(lines []string) [][]int {
	var out [][]int
	for _, line := range lines {
		var fieldsI []int
		for i := 0; i < len(line); i++ {
			val, _ := strconv.Atoi(string(line[i]))
			fieldsI = append(fieldsI, val)
		}
		out = append(out, fieldsI)
	}
	return out
}

func Day10() {
	lines := readFile("data/day10.txt")
	linesInt := linesToNumbers(lines)
	log.Println("Day 10 Part 1: ", day10(linesInt, true))
	log.Println("Day 10 Part 2: ", day10(linesInt, false))
}

func getTrailHeads(lines [][]int) []Point {
	var trailHeads []Point
	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines[i]); j++ {
			num := lines[i][j]
			if num == 0 {
				trailHeads = append(trailHeads, Point{X: i, Y: j})
			}
		}
	}
	return trailHeads
}

func hike(lines [][]int, position Point, visitedPositions *map[Point]bool, peakCount *int, checkVisitedPositions bool) {
	if checkVisitedPositions {
		_, ok := (*visitedPositions)[position]
		if ok {
			// already visited
			return
		}
		(*visitedPositions)[position] = true
	}

	currValue := lines[position.X][position.Y]
	if currValue == 9 {
		(*peakCount) += 1
	} else {
		upX := position.X - 1
		rightY := position.Y + 1
		downX := position.X + 1
		leftY := position.Y - 1
		if upX != -1 && currValue+1 == lines[upX][position.Y] {
			hike(lines, Point{upX, position.Y}, visitedPositions, peakCount, checkVisitedPositions)
		}
		if rightY < len(lines) && currValue+1 == lines[position.X][rightY] {
			hike(lines, Point{position.X, rightY}, visitedPositions, peakCount, checkVisitedPositions)
		}
		if downX < len(lines[0]) && currValue+1 == lines[downX][position.Y] {
			hike(lines, Point{downX, position.Y}, visitedPositions, peakCount, checkVisitedPositions)
		}
		if leftY != -1 && currValue+1 == lines[position.X][leftY] {
			hike(lines, Point{position.X, leftY}, visitedPositions, peakCount, checkVisitedPositions)
		}
	}
}

func day10(lines [][]int, checkVisitedPositions bool) int {
	trailHeads := getTrailHeads(lines)
	sumScores := 0
	for _, trailHead := range trailHeads {
		peakCount := 0
		visitedPoints := make(map[Point]bool)
		hike(lines, trailHead, &visitedPoints, &peakCount, checkVisitedPositions)
		sumScores += peakCount
	}
	return sumScores
}
