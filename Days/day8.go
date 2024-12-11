package Days

import (
	"log"
)

type Point struct {
	X int
	Y int
}

func Day8() {
	lines := readFile("data/day8.txt")
	antennas := getAntennas(lines)
	numRows := len(lines)
	numCols := len(lines[0])
	antinodes := day8Part1(antennas, false, 0, numRows, numCols)
	log.Println("Day 8 Part 1: ", len(antinodes))
	log.Println("Day 8 Part 2: ", day8Part2(antennas, numRows, numCols))
}

func calculateAndAddAntinodes(
	antennaCompare1 Point,
	antennaCompare2 Point,
	antinodes *map[Point]bool,
	numRows int,
	numCols int,
	part2 bool,
	recursionPath int) {
	xDiff := AbsInt(antennaCompare1.X - antennaCompare2.X)
	yDiff := AbsInt(antennaCompare1.Y - antennaCompare2.Y)
	var antinodeX1, antinodeY1, antinodeX2, antinodeY2 int
	// antennaCompare1 is "above" antennaCompare2
	if antennaCompare1.X < antennaCompare2.X {
		antinodeX1 = antennaCompare1.X - xDiff
		antinodeX2 = antennaCompare2.X + xDiff
	} else {
		antinodeX1 = antennaCompare1.X + xDiff
		antinodeX2 = antennaCompare2.X - xDiff
	}
	// antennaCompare1 is "left to" antennaCompare2
	if antennaCompare1.Y < antennaCompare2.Y {
		antinodeY1 = antennaCompare1.Y - yDiff
		antinodeY2 = antennaCompare2.Y + yDiff
	} else {
		antinodeY1 = antennaCompare1.Y + yDiff
		antinodeY2 = antennaCompare2.Y - yDiff
	}
	if recursionPath != 2 && antinodeX1 >= 0 && antinodeY1 >= 0 && antinodeX1 < numRows && antinodeY1 < numCols {
		newAntinode := Point{X: antinodeX1, Y: antinodeY1}
		(*antinodes)[newAntinode] = true
		if part2 {
			calculateAndAddAntinodes(newAntinode, antennaCompare1, antinodes, numRows, numCols, part2, 1)
		}
	}
	if recursionPath != 1 && antinodeX2 >= 0 && antinodeY2 >= 0 && antinodeX2 < numRows && antinodeY2 < numCols {
		newAntinode := Point{X: antinodeX2, Y: antinodeY2}
		(*antinodes)[newAntinode] = true
		if part2 {
			calculateAndAddAntinodes(antennaCompare2, newAntinode, antinodes, numRows, numCols, part2, 2)
		}
	}
}
func getAntennas(lines []string) map[string][]Point {
	antennas := make(map[string][]Point)
	for x := 0; x < len(lines); x++ {
		for y := 0; y < len(lines[x]); y++ {
			symbol := string(lines[x][y])
			if symbol != "." {
				antennas[symbol] = append(antennas[symbol], Point{X: x, Y: y})
			}
		}
	}
	return antennas
}

func day8Part1(antennas map[string][]Point, part2 bool, recursionPath int, numRows int, numCols int) map[Point]bool {
	antinodes := make(map[Point]bool)
	for _, antenna := range antennas {
		// compare every antenna to every other antenna of the same symbol
		for idx, antennaCompare1 := range antenna {
			for idy, antennaCompare2 := range antenna {
				if idx == idy {
					continue
				}
				calculateAndAddAntinodes(antennaCompare1, antennaCompare2, &antinodes, numRows, numCols, part2, recursionPath)
			}
		}
	}
	return antinodes
}

func day8Part2(antennas map[string][]Point, numRows int, numCols int) int {
	antinodes := day8Part1(antennas, true, 0, numRows, numCols)
	// just put all antennas in the set of antinodes to ignore duplicates
	for _, antenna := range antennas {
		for _, ant := range antenna {
			antinodes[ant] = true
		}
	}
	return len(antinodes)
}
