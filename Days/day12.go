package Days

import (
	"cmp"
	"log"
	"slices"
)

func Day12() {
	lines := readFile("data/day12.txt")
	log.Println("Day 12 Part 1: ", day12Part1(lines))
	log.Println("Day 12 Part 2: ", day12Part2(lines))
}

func collectRegions(x int, y int, lines []string, region *map[Plant]struct{}) {
	plant := Plant{x: x, y: y, type_: lines[x][y]}
	// skip plants already visited
	if _, ok := (*region)[plant]; ok {
		return
	}
	(*region)[plant] = struct{}{}
	//up
	if x > 0 && lines[x-1][y] == plant.type_ {
		collectRegions(x-1, y, lines, region)
	}
	//right
	if y < len(lines[x])-1 && lines[x][y+1] == plant.type_ {
		collectRegions(x, y+1, lines, region)
	}

	//down
	if x < len(lines)-1 && lines[x+1][y] == plant.type_ {
		collectRegions(x+1, y, lines, region)
	}
	//left
	if y > 0 && lines[x][y-1] == plant.type_ {
		collectRegions(x, y-1, lines, region)
	}
}

type Plant struct {
	x     int
	y     int
	type_ byte
}

func CmpPlantByX(a, b Plant) int {
	return cmp.Compare(a.x, b.x)
}

func CmpPlantByY(a, b Plant) int {
	return cmp.Compare(a.y, b.y)
}
func getRegions(lines []string) []map[Plant]struct{} {
	var regions []map[Plant]struct{}
	for x := 0; x < len(lines); x++ {
		for y := 0; y < len(lines[x]); y++ {
			// skip if the current plant is in any region already
			skip := false
			for _, region := range regions {
				if _, ok := region[Plant{x: x, y: y, type_: lines[x][y]}]; ok {
					skip = true
					break
				}
			}
			if !skip {
				region := make(map[Plant]struct{})
				collectRegions(x, y, lines, &region)
				regions = append(regions, region)
			}
		}
	}
	return regions
}

func countFenceNeeded(region map[Plant]struct{}) int {
	fenceNeeded := 0
	for plant := range region {
		// check how many edges are not surrounded by other plants.
		// check if the nearby plant is not in the current region
		//up
		if _, ok := region[Plant{x: plant.x - 1, y: plant.y, type_: plant.type_}]; !ok {
			fenceNeeded += 1
		}
		//right
		if _, ok := region[Plant{x: plant.x, y: plant.y + 1, type_: plant.type_}]; !ok {
			fenceNeeded += 1
		}
		//down
		if _, ok := region[Plant{x: plant.x + 1, y: plant.y, type_: plant.type_}]; !ok {
			fenceNeeded += 1
		}
		//left
		if _, ok := region[Plant{x: plant.x, y: plant.y - 1, type_: plant.type_}]; !ok {
			fenceNeeded += 1
		}
	}
	return fenceNeeded
}

func day12Part1(lines []string) int {
	// find all regions
	regions := getRegions(lines)
	// count the fence needed
	price := 0
	for _, region := range regions {
		fenceNeeded := countFenceNeeded(region)
		price += fenceNeeded * len(region)
	}
	return price
}

func findPlant(region map[Plant]struct{}, x, y int) (Plant, bool) {
	for plant := range region {
		if plant.x == x && plant.y == y {
			return plant, true
		}
	}
	return Plant{}, false
}

func checkTopAndBottom(lines []string, region map[Plant]struct{}, directionUp bool) int {
	edges := 0
	for x := 0; x < len(lines); x++ {
		var plantsInThisRow []Plant
		for plant := range region {
			if plant.x == x {
				plantsInThisRow = append(plantsInThisRow, plant)
			}
		}
		slices.SortFunc(plantsInThisRow, CmpPlantByY)
		for idx, plant := range plantsInThisRow {
			// current one blocked on the side where the fence should be?
			var xAboveOrBelow int
			if directionUp {
				xAboveOrBelow = plant.x - 1
			} else {
				xAboveOrBelow = plant.x + 1
			}
			_, ok := findPlant(region, xAboveOrBelow, plant.y)
			if !ok {
				// check if the next one is blocked off by one on top. this means the current fencing ends here
				_, ok2 := findPlant(region, xAboveOrBelow, plant.y+1)
				if ok2 {
					edges++
					continue
				}
				// check if there is a plant next to this one or if this plant is at the very right of the field.
				if idx == len(lines)-1 || idx == len(plantsInThisRow)-1 || plantsInThisRow[idx+1].y-1 != plant.y {
					edges++
				}
			}

		}
	}
	return edges
}

func checkLeftAndRight(lines []string, region map[Plant]struct{}, directionLeft bool) int {
	edges := 0
	for y := 0; y < len(lines); y++ {
		var plantsInThisCol []Plant
		for plant := range region {
			if plant.y == y {
				plantsInThisCol = append(plantsInThisCol, plant)
			}
		}
		slices.SortFunc(plantsInThisCol, CmpPlantByX)
		for idx, plant := range plantsInThisCol {
			// current one blocked on the side where the fence should be?
			var yLeftOrRight int
			if directionLeft {
				yLeftOrRight = plant.y - 1
			} else {
				yLeftOrRight = plant.y + 1
			}
			// current one blocked on the side where the fence should be?
			_, ok := findPlant(region, plant.x, yLeftOrRight)
			if !ok {
				// check if the next one is blocked off by one on top. this means the current fencing ends here
				_, ok2 := findPlant(region, plant.x+1, yLeftOrRight)
				if ok2 {
					edges++
					continue
				}
				// check if there is a plant next to this one or if this plant is at the very right of the field.
				if idx == len(lines[0])-1 || idx == len(plantsInThisCol)-1 || plantsInThisCol[idx+1].x-1 != plant.x {
					edges++
				}
			}
		}
	}
	return edges
}

func countEdges(lines []string, region map[Plant]struct{}) int {
	edges := 0
	// 1# scan from top to bottom: search fencings "on top"
	edges += checkTopAndBottom(lines, region, true)
	// 2# scan from top to bottom: search fencings "on bottom"
	edges += checkTopAndBottom(lines, region, false)
	// 3# scan from left to right: search fencings "on left"
	edges += checkLeftAndRight(lines, region, true)
	// 4# scan from left to right: search fencings "on right"
	edges += checkLeftAndRight(lines, region, false)
	return edges
}

func day12Part2(lines []string) int {
	// find all regions
	regions := getRegions(lines)
	// count the fence needed
	price := 0
	for _, region := range regions {
		numEdges := countEdges(lines, region)
		price += numEdges * len(region)
	}
	return price
}
