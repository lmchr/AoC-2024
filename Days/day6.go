package Days

import (
	"log"
	"sync"
)

func Day6() {
	lines := readFile("data/day6.txt")
	lookup := map[string]int{"^": 0, ">": 1, "v": 2, "<": 3}
	guardX, guardY := getPositionOfGuard(lines, lookup)
	currentGuardSymbol := lookup[string(lines[guardX][guardY])]
	linesBools := linesCharsToLinesBool(lines)
	fieldsVisited := day6Part1(restoreField(linesBools), guardX, guardY, currentGuardSymbol)
	log.Println("Day 6 Part 1: ", len(fieldsVisited))
	log.Println("Day 6 Part 2: ", day6Part2(restoreField(linesBools), guardX, guardY, currentGuardSymbol, fieldsVisited))
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
func simulateOneStep(currentGuardSymbol *int, guardX *int, guardY *int, lines *[][]bool) bool {
	/*
		Returns whether the guard exited the field

	*/
	// guard may turn multiple times at once, but at most 3 times. otherwise the guard is stuck
	for idx := range 3 {
		turned := false
		switch *currentGuardSymbol {
		case 0:
			if *guardX != 0 && (*lines)[*guardX-1][*guardY] {
				*currentGuardSymbol = 1
				turned = true
			}
		case 1:
			if *guardY+1 < len((*lines)[0]) && (*lines)[*guardX][*guardY+1] {
				*currentGuardSymbol = 2
				turned = true
			}
		case 2:
			if *guardX+1 < len((*lines)) && (*lines)[*guardX+1][*guardY] {
				*currentGuardSymbol = 3
				turned = true
			}
		case 3:
			if *guardY != 0 && (*lines)[*guardX][*guardY-1] {
				*currentGuardSymbol = 0
				turned = true
			}
		}
		if !turned {
			break
		}
		if idx == 4 {
			panic("Guard is stuck.")
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
	return *guardX < 0 || *guardY < 0 || *guardX >= len(*lines) || *guardY >= len((*lines)[0])
}

type Position struct {
	X int
	Y int
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

func day6Part1(lines [][]bool, guardX int, guardY int, currentGuardSymbol int) map[Position]bool {
	fieldsVisited := make(map[Position]bool)
	fieldsVisited[Position{X: guardX, Y: guardY}] = true
	for {
		if simulateOneStep(&currentGuardSymbol, &guardX, &guardY, &lines) {
			break
		} else {
			fieldsVisited[Position{X: guardX, Y: guardY}] = true
		}
	}
	return fieldsVisited
}

func day6Part2Parallel(lines [][]bool, x int, y int, guardX int, guardY int, guardDirection int, ch chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	// block cell
	lines[x][y] = true
	loopDetected := false
	pathsTaken := []Position{{X: guardX, Y: guardY}}
	for {
		if simulateOneStep(&guardDirection, &guardX, &guardY, &lines) {
			break
		}
		pathsTaken = append(pathsTaken, Position{X: guardX, Y: guardY})
		// check once 4 positions were visited. before this, no loop can be checked
		if len(pathsTaken) >= 4 {
			// check if the last two steps were done somewhere before this. This means that we have a loop
			for i := 0; i < len(pathsTaken)-3; i++ {
				if (pathsTaken[i] == pathsTaken[len(pathsTaken)-2]) && (pathsTaken[i+1] == pathsTaken[len(pathsTaken)-1]) {
					loopDetected = true
				}
			}
		}
		if loopDetected {
			break
		}
	}
	ch <- loopDetected
}

func day6Part2(lines [][]bool, guardX int, guardY int, guardDirection int, fieldsVisitedPart1 map[Position]bool) int {
	guardXStart := guardX
	guardYStart := guardY
	blockedCounter := 0
	linesBackup := restoreField(lines)
	wg := new(sync.WaitGroup)
	ch := make(chan bool)
	for field := range fieldsVisitedPart1 {
		// The new obstruction can't be placed at the guard's starting position - the guard is there right now and would notice.
		if field.X == guardXStart && field.Y == guardYStart {
			continue
		}
		// already blocked
		if lines[field.X][field.Y] {
			continue
		}
		wg.Add(1)
		go day6Part2Parallel(restoreField(linesBackup), field.X, field.Y, guardX, guardY, guardDirection, ch, wg)
	}
	go func() {
		wg.Wait()
		close(ch)
	}()
	for blocked := range ch {
		if blocked {
			blockedCounter++
		}
	}
	return blockedCounter
}
