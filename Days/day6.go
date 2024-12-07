package Days

import (
	"fmt"
	"log"
)

func Day6() {
	lines := readFile("data/day6.txt")
	lookup := map[string]int{"^": 0, ">": 1, "v": 2, "<": 3}
	guardX, guardY := getPositionOfGuard(lines, lookup)
	currentGuardSymbol := lookup[string(lines[guardX][guardY])]
	linesBools := linesCharsToLinesBool(lines)

	log.Println("Day 6 Part 1: ", day6Part1(restoreField(linesBools), guardX, guardY, currentGuardSymbol))
	log.Println("Day 6 Part 2: ", day6Part2(restoreField(linesBools), guardX, guardY, currentGuardSymbol))
}

func getPositionOfGuard(lines []string, lookup map[string]int) (int, int) {
	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines[i]); j++ {
			if _, ok := lookup[string(lines[i][j])]; ok {
				return i, j
			}
		}
	}
	panic("Guard position was not found.")
}

func simulateOneStepBool(currentGuardSymbol *int, guardX *int, guardY *int, linesBool *[][]bool) bool {
	/*
		Returns whether the guard exited the field
	*/
	switch *currentGuardSymbol {
	case 0:
		if *guardX == 0 {
			break
		}
		if (*linesBool)[*guardX-1][*guardY] {
			*currentGuardSymbol = 1
		}
	case 1:
		if *guardY+1 >= len((*linesBool)[0]) {
			break
		}
		if (*linesBool)[*guardX][*guardY+1] {
			*currentGuardSymbol = 2
		}
	case 2:
		if *guardX+1 >= len((*linesBool)) {
			break
		}
		if (*linesBool)[*guardX+1][*guardY] {
			*currentGuardSymbol = 3
		}
	case 3:
		if *guardY == 0 {
			break
		}
		if (*linesBool)[*guardX][*guardY-1] {
			*currentGuardSymbol = 0
		}
	}
	switch *currentGuardSymbol {
	case 0:
		(*guardX)--
	case 1:
		(*guardY)++
	case 2:
		(*guardX)++
	case 3:
		(*guardY)--
	}
	if *guardX < 0 || *guardY < 0 || *guardX >= len(*linesBool) || *guardY >= len((*linesBool)[0]) {
		return true
	}
	return false
}

type Position struct {
	X         int
	Y         int
	direction int
}

func linesCharsToLinesBool(linesChars []string) [][]bool {
	var linesBool [][]bool
	for x := 0; x < len(linesChars); x++ {
		var linesBoolSubArray []bool
		for y := 0; y < len(linesChars[x]); y++ {
			var isBlocked = string(linesChars[x][y]) == "#"
			linesBoolSubArray = append(linesBoolSubArray, isBlocked)
		}
		linesBool = append(linesBool, linesBoolSubArray)
	}
	return linesBool
}

func day6Part1(linesBools [][]bool, guardX int, guardY int, currentGuardSymbol int) int {
	fieldsVisited := make(map[Position]bool)
	fieldsVisited[Position{X: guardX, Y: guardY, direction: currentGuardSymbol}] = true
	for {
		if simulateOneStepBool(&currentGuardSymbol, &guardX, &guardY, &linesBools) {
			break
		} else {
			fieldsVisited[Position{X: guardX, Y: guardY, direction: currentGuardSymbol}] = true
		}
	}
	return len(fieldsVisited)
}

func day6Part2(linesBools [][]bool, guardX int, guardY int, currentGuardSymbol int) int {
	guardXStart := guardX
	guardYStart := guardY
	currentGuardSymbolStart := currentGuardSymbol
	blockedCounter := 0
	linesBoolsBackup := restoreField(linesBools)
	// block every index on the field ony by one
	for x := 0; x < len(linesBools); x++ {
		for y := 0; y < len(linesBools[x]); y++ {
			done := x*len(linesBools[x]) + y
			from := len(linesBools)*len(linesBools[0]) - 1
			fmt.Println(done, " / ", from, float32(done)/float32(from))
			// The new obstruction can't be placed at the guard's starting position - the guard is there right now and would notice.
			if x == guardXStart && y == guardYStart {
				continue
			}
			// already blocked
			if linesBools[x][y] {
				continue
			}
			linesBools[x][y] = true
			loopDetected := false
			var pathsTaken []Position
			for {
				pathsTaken = append(pathsTaken, Position{X: guardX, Y: guardY, direction: currentGuardSymbol})
				if simulateOneStepBool(&currentGuardSymbol, &guardX, &guardY, &linesBools) {
					break
				}
				// check once 4 positions were visited. before this, no loop can be checked
				if len(pathsTaken) >= 4 {
					// check if the last two steps were done somewhere before this. This means that we have a loop
					for i := 0; i < len(pathsTaken)-2; i++ {
						if (pathsTaken[i] == pathsTaken[len(pathsTaken)-2]) && (pathsTaken[i+1] == pathsTaken[len(pathsTaken)-1]) {
							loopDetected = true
							break
						}
					}
				}
				if loopDetected {
					break
				}
			}
			if loopDetected {
				blockedCounter++
			}
			// undo the blocking and modifications of the field
			linesBools = restoreField(linesBoolsBackup)
			guardX = guardXStart
			guardY = guardYStart
			currentGuardSymbol = currentGuardSymbolStart
		}
	}

	return blockedCounter
}
